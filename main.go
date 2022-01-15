package main

import (
	"52lu/go-import-template/global"
	"52lu/go-import-template/initialize"
	"52lu/go-import-template/middleware"
	"52lu/go-import-template/router/demo"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	// 程序退出前释放资源
	defer closeResource()
	// 加载启动前配置
	initialize.SetLoadInit()
	// 启动服务
	RunServer()
}

// 程序退出前释放资源
func closeResource() {
	// 关闭数据库连接
	if global.GvaMysqlClient != nil {
		db, _ := global.GvaMysqlClient.DB()
		_ = db.Close()
	}
}

// RunServer 启动服务
func RunServer() {
	engine := gin.New()
	// 注册公共中间件
	engine.Use(gin.Recovery(), middleware.CatchErrorMiddleWare())
	// 获取自定义http配置
	httpServer := getCustomHttpServer(engine)
	// 注册路由
	registerRouters(engine)
	// 打印服务信息
	printServerInfo()
	// 启动服务
	err := httpServer.ListenAndServe()
	if err != nil {
		panic("启动失败: " + err.Error())
	}
}

// 获取自定义HTTP SERVER
func getCustomHttpServer(engine *gin.Engine) *http.Server {
	// 创建自定义配置服务
	httpServer := &http.Server{
		//ip和端口号
		Addr: global.GvaConfig.App.Addr,
		//调用的处理器，如为nil会调用http.DefaultServeMux
		Handler: engine,
		//计算从成功建立连接到request body(或header)完全被读取的时间
		ReadTimeout: time.Second * 10,
		//计算从request body(或header)读取结束到 response write结束的时间
		WriteTimeout: time.Second * 10,
		//请求头的最大长度，如为0则用DefaultMaxHeaderBytes
		MaxHeaderBytes: 1 << 20,
	}
	return httpServer
}

// 打印服务信息
func printServerInfo() {
	appConfig := global.GvaConfig.App
	fmt.Printf("\n【 当前环境: %s 当前版本: %s 接口地址: http://%s 】\n", appConfig.Env, appConfig.Version, appConfig.Addr)
}

// 注册路由入口
func registerRouters(engine *gin.Engine) {
	// 注册演示路由
	demo.InitDemoRouter(engine)
}
