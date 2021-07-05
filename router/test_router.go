/**
 * @Author Mr.LiuQH
 * @Description TODO
 * @Date 2021/7/5 3:44 下午
 **/
package router

import (
	v1 "52lu/go-import-template/api/v1"
	"github.com/gin-gonic/gin"
)

// 测试接口
func InitTestRouter(group *gin.RouterGroup)  {
	group.GET("/hello",v1.Hello)
	group.GET("/test",v1.Test)
}