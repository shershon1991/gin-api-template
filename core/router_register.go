/**
 * @Author Mr.LiuQH
 * @Description 路由注册入口
 * @Date 2021/7/5 3:17 下午
 **/
package core

import (
	"52lu/go-import-template/router"
	"github.com/gin-gonic/gin"
)
// 注册路由入口
func RegisterRouters(engine *gin.Engine)  {
	initV1(engine)
}

func initV1(engine *gin.Engine)  {
	// v1版本接口
	v1 := engine.Group("v1")
	router.InitTestRouter(v1)
}