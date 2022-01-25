package api

import (
	"douban-webend/api/middleware"
	"douban-webend/api/users"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var Addr = ":8080"

type HandleFunctions []gin.HandlerFunc
type Routes map[string][]Route // Routes	group(key) => []Route

// Route 表示每一个路由
type Route struct {
	Name             string
	Method           string // it is the string for the HTTP method. ex) GET, POST etc..
	Pattern          string
	HandlerFunctions HandleFunctions `json:"-"`
}

// 所有路由
var routes = Routes{
	"/": []Route{
		{
			Name:             "外链跳转",
			Method:           http.MethodGet,
			Pattern:          "/wild",
			HandlerFunctions: HandleFunctions{middleware.WildChecker(), handleWild},
		},
		{
			Name:             "Swagger文档",
			Method:           http.MethodGet,
			Pattern:          "/swagger",
			HandlerFunctions: HandleFunctions{},
		},
		{
			Name:             "我的主页",
			Method:           http.MethodGet,
			Pattern:          "/mine",
			HandlerFunctions: HandleFunctions{middleware.Auth(), handleMine},
		},
	},
	"/users": []Route{
		{
			Name:             "用户登录",
			Method:           http.MethodPost,
			Pattern:          "/login",
			HandlerFunctions: HandleFunctions{users.HandleLogin},
		},
		{
			Name:             "OAuth登录",
			Method:           http.MethodGet,
			Pattern:          "/login",
			HandlerFunctions: HandleFunctions{users.HandleOAuthRedirect},
		},
		{
			Name:             "用户注册",
			Method:           http.MethodPost,
			Pattern:          "/register",
			HandlerFunctions: HandleFunctions{users.HandleRegister},
		},
		{
			Name:             "发送验证码",
			Method:           http.MethodGet,
			Pattern:          "/verify",
			HandlerFunctions: HandleFunctions{users.HandleVerify},
		},
		{
			Name:             "获取用户的主页信息",
			Method:           http.MethodGet,
			Pattern:          "/:id",
			HandlerFunctions: HandleFunctions{users.HandleAccountIndexInfo},
		},
		{
			Name:             "获取用户的想看",
			Method:           http.MethodGet,
			Pattern:          "/:id/before",
			HandlerFunctions: HandleFunctions{users.HandleUserBefore},
		},
		{
			Name:             "获取用户的看过",
			Method:           http.MethodGet,
			Pattern:          "/:id/after",
			HandlerFunctions: HandleFunctions{users.HandleUserAfter},
		},
	},
	"/subjects": []Route{
		{
			Name:             "获取电影列表", // 该路由压力较大，考虑使用集群
			Method:           http.MethodGet,
			Pattern:          "/",
			HandlerFunctions: HandleFunctions{},
		},
	},
	"/oauth": []Route{
		{
			Name:             "OAuth Redirect uri",
			Method:           http.MethodGet,
			Pattern:          "/:platform",
			HandlerFunctions: HandleFunctions{users.HandleOAuthLogin},
		},
	},
}

func newRouter(useTLS bool) *gin.Engine {
	engine := gin.Default()

	for k, v := range routes {
		for _, router := range v {
			relativePath := k + router.Pattern
			engine.Handle(router.Method, relativePath, router.HandlerFunctions...)
		}
	}

	engine.Use(middleware.Cors()) // 跨域

	if useTLS {
		engine.Use(middleware.TLSHandle(Addr))
	}

	return engine
}

func InitRouter(useTLS bool) {
	log.Println("Server started!")

	router := newRouter(useTLS)

	if useTLS {
		log.Fatalln(router.RunTLS(Addr, "config/api.pem", "config/api.key"))
	} else {
		log.Fatalln(router.Run(Addr))
	}

}
