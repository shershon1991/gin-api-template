/**
 * @Description 用户相关的路由
 **/
package router

import (
	v1 "52lu/go-import-template/api/v1"
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
}
