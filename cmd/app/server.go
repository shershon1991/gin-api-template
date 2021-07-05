package app

import (
	"52lu/go-import-template/core"
	"github.com/gin-gonic/gin"
)

// RunServer 启动服务
func RunServer() {
	// 创建默认容器
	engine := gin.Default()
	// 注册路由
	core.RegisterRouters(engine)
	// 启动服务
	_ = engine.Run()
}
