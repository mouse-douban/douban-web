package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

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
			HandlerFunctions: HandleFunctions{},
		},
		{
			Name:             "Swagger文档",
			Method:           http.MethodGet,
			Pattern:          "/swagger",
			HandlerFunctions: HandleFunctions{},
		},
	},
	"/users": []Route{
		{
			Name: "用户登录",
			Method: http.MethodPost,
			Pattern: "/login",
			HandlerFunctions: HandleFunctions{},
		},
		{
			Name: "OAuth登录",
			Method: http.MethodGet,
			Pattern: "/login",
			HandlerFunctions: HandleFunctions{},
		},
		{
			Name: "用户注册",
			Method: http.MethodPost,
			Pattern: "/register",
			HandlerFunctions: HandleFunctions{},
		},
		{
			Name: "发送验证码",
			Method: http.MethodGet,
			Pattern: "/verify",
			HandlerFunctions: HandleFunctions{},
		},
		{
			Name: "获取用户的主页信息",
			Method: http.MethodGet,
			Pattern: "/:id",
			HandlerFunctions: HandleFunctions{},
		},
		{
			Name: "获取用户的想看",
			Method: http.MethodGet,
			Pattern: "/:id/before",
			HandlerFunctions: HandleFunctions{},
		},
		{
			Name: "获取用户的看过",
			Method: http.MethodGet,
			Pattern: "/:id/after",
			HandlerFunctions: HandleFunctions{},
		},
	},
	"/subjects": []Route{
		{
			Name: "获取电影列表", // 该路由压力较大，考虑使用集群
			Method: http.MethodGet,
			Pattern: "/",
			HandlerFunctions: HandleFunctions{},
		},
	},
}

func newRouter() *gin.Engine {
	engine := gin.Default()

	for k, v := range routes {
		for _, router := range v {
			relativePath := k + router.Pattern
			engine.Handle(router.Method, relativePath, router.HandlerFunctions...)
		}
	}

	return engine
}

func InitRouter() {
	log.Println("Server started!")

	router := newRouter()

	log.Fatalln(router.Run(":8080"))
}
