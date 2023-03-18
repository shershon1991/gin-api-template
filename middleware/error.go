package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shershon1991/gin-api-template/global"
	"shershon1991/gin-api-template/model/response"
)

// 捕获请求全局错误
func CatchErrorMiddleWare() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 捕获错误
		defer func() {
			if err := recover(); err != nil {
				errMsg := fmt.Sprintf("运行异常: %s", err)
				// 捕获错误
				if global.GvaLogger != nil {
					global.GvaLogger.Error(errMsg)
				}
				// todo邮件通知

				// 错误响应
				response.Error(ctx, errMsg)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
