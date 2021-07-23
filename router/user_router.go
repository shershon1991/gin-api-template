/**
 * @Description 用户相关的路由
 **/
package router

import (
	v1 "52lu/go-import-template/api/v1"
	"52lu/go-import-template/middleware"
	"github.com/gin-gonic/gin"
)

/**
 * @description: 用户相关的路由
 * @param engine
 */
func InitUserRouter(engine *gin.Engine) {
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
}
