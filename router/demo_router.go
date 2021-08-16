/**
 * @Author Mr.LiuQH
 * @Description 不需要验证登录
 * @Date 2021/7/5 3:44 下午
 **/
package router

import (
	v1 "52lu/go-import-template/api/demo"
	"52lu/go-import-template/middleware"
	"github.com/gin-gonic/gin"
)

// 系统路由
func InitDemoRouter(engine *gin.Engine) {
	// 系统路由
	systemRouter := engine.Group("system")
	{
		// 获取全局变量
		systemRouter.GET("config", v1.GetConfig)

	}
	// 不需要登录的路由
	noLoginGroup := engine.Group("v1/user")
	{
		// 登录
		noLoginGroup.POST("login", v1.Login)
		// 注册
		noLoginGroup.POST("register", v1.Register)
	}
	// 需要登录
	tokenGroup := engine.Group("v1/user").Use(middleware.JWTAuthMiddleware())
	{
		tokenGroup.POST("/detail", v1.GetUser)
	}
	// 测试路由
	testRouter := engine.Group("test")
	{
		// redis测试使用
		testRouter.GET("redis", v1.RdTest)
	}
	// es相关路由
	esGroup := engine.Group("es")
	{
		esGroup.GET("create", v1.CreateIndex)
		esGroup.GET("searchById", v1.SearchById)
	}
}
