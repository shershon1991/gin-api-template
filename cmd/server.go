package cmd

import "github.com/gin-gonic/gin"

// RunServer 启动服务
func RunServer() {
	// 创建容器
	engine := gin.Default()
	// todo 注册路由

	// 启动服务
	_ = engine.Run()
}
