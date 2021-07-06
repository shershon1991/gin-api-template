/**
 * @Author Mr.LiuQH
 * @Description 不需要验证登录
 * @Date 2021/7/5 3:44 下午
 **/
package router

import (
	v1 "52lu/go-import-template/api/v1"
	"github.com/gin-gonic/gin"
)

// 不需要验证登录的路由
func InitNoLoginRouter(group *gin.RouterGroup)  {
	// 获取全局变量
	group.GET("/system/config",v1.GetConfig)

}
