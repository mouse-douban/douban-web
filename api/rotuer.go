package api

import (
	"context"
	"douban-webend/api/middleware"
	"douban-webend/api/users"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

var (
	Addr                  = ":8080"
	server   *http.Server = nil
	listener net.Listener = nil
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

// 这种路由注册的设计是从 Apifox 自动生成的 mock 代码里学到的，并且改了一点来适配分级路由
// TODO 设计 Group 的注册方式
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

	engine.Use(middleware.Cors()) // 跨域，放在服务路由加载前

	if useTLS {
		engine.Use(middleware.TLSHandle(Addr)) // TLS
	}

	for k, v := range routes {
		for _, router := range v {
			relativePath := k + router.Pattern
			engine.Handle(router.Method, relativePath, router.HandlerFunctions...)
		}
	}

	return engine
}

func InitRouter(useTLS bool) {
	log.Println("Server started!")

	router := newRouter(useTLS)

	server = &http.Server{
		Addr:    Addr,
		Handler: router,
	}

	var err error

	// os.Args[0] 是命令名称，从 1 开始才是参数
	if len(os.Args) == 2 && os.Args[1] == "reload" { // 子进程接管
		// 这里是 3 的原因是 一个进程的 linux 文件描述符，0，1，2 分别代表标准输入，标准输出，标准错误输出，已经被占用
		// 所以父进程传递过来的文件，描述符是从 3 开始的，这里使用 3 来获得父进程传入的 tcp socket 的文件描述符 ( 果然 linux 万物皆文件
		listenerFd := os.NewFile(3, "")
		// 拿到一个新的 listener，父进程任务已经完成，(龟野先生，天皇陛下...（大雾
		listener, err = net.FileListener(listenerFd)
	} else { // 常规启动
		listener, err = net.Listen("tcp", Addr)
	}

	if err != nil {
		log.Fatalf("listen tcp %v failed, cause %v\n", Addr, err)
	}

	go func() { //不要阻塞主 goroutine
		if useTLS {
			err = server.ServeTLS(listener, "config/api.pem", "config/api.key")
			if err != nil {
				log.Fatalf("serve err %v\n", err)
			}
		} else {
			err = server.Serve(listener)
			if err != nil {
				log.Fatalf("serve err %v\n", err)
			}
		}
	}()

	listenSignal()
}

func listenSignal() {
	// 监听部分
	sig := make(chan os.Signal, 1)
	// 监听关闭信号和重启信号
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	for {
		sign := <-sig
		log.Printf("receive signal: %v\n", sign)
		ctx, _ := context.WithTimeout(context.Background(), time.Second*10) // 10s延迟，到点强制回收资源(不执行cancelFunc)
		switch sign {
		case syscall.SIGINT, syscall.SIGTERM: // 中止信号  kill INT [pid] 或者 Ctrl+C 发送的就是这两种信号，kill -9 的信号截停不了（保证用户能够终止程序
			log.Println("Server shutdown...")
			signal.Stop(sig) // 停止信号
			err := server.Shutdown(ctx)
			if err != nil {
				log.Fatalf("shutdown err!! %v\n", err)
				return
			}
			log.Println("shutdown gracefully...")
		case syscall.SIGUSR2: // 重启信号  kill -31/-USR2 [pid]
			log.Println("Server reloading...")
			err := reload()
			if err != nil {
				log.Fatalf("reload err!! %v\n", err)
			}
			err = server.Shutdown(ctx)
			if err != nil {
				log.Fatalf("shutdown err!! %v\n", err)
				return
			}
		}
	}
}

// 热重启
func reload() error {

	tcpListener, _ := listener.(*net.TCPListener)

	file, err := tcpListener.File() // 拿到 tcp socket 文件
	if err != nil {
		return err
	}

	// 使用命令来重新运行一遍
	log.Printf("all file prepared, use cmd ' %v reload ' to reload\n", os.Args[0])
	cmd := exec.Command(os.Args[0], "reload")
	cmd.Stdin = os.Stdin   // 绑定 fd 0
	cmd.Stdout = os.Stdout // 绑定 fd 1
	cmd.Stderr = os.Stderr // 绑定 fd 2

	// 剩下绑定的是从 3 开始
	cmd.ExtraFiles = []*os.File{file}

	return cmd.Start()
}
