/**
 * @Author Mr.LiuQH
 * @Description 不需要验证登录
 * @Date 2021/7/5 3:44 下午
 **/
package router

import (
	"52lu/go-import-template/api/demo"
	"52lu/go-import-template/middleware"
	"github.com/gin-gonic/gin"
)

// 系统路由
func InitDemoRouter(engine *gin.Engine) {
	// 系统路由
	systemRouter := engine.Group("system")
	{
		// 获取全局变量
		systemRouter.GET("config", demo.GetConfig)

	}
	// 不需要登录的路由
	noLoginGroup := engine.Group("demo/user")
	{
		// 登录
		noLoginGroup.POST("login", demo.Login)
		// 注册
		noLoginGroup.POST("register", demo.Register)
	}
	// 需要登录
	tokenGroup := engine.Group("demo/user").Use(middleware.JWTAuthMiddleware())
	{
		tokenGroup.POST("/detail", demo.GetUser)
	}
	// 测试路由
	testRouter := engine.Group("test")
	{
		// redis测试使用
		testRouter.GET("redis", demo.RdTest)
	}
	// es相关路由
	esGroup := engine.Group("es")
	{
		esGroup.GET("create", demo.CreateIndex)
		esGroup.GET("searchById", demo.SearchById)
	}
}
