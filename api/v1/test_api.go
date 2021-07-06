/**
 * @Author Mr.LiuQH
 * @Description 测试专用api
 * @Date 2021/7/5 3:42 下午
 **/
package v1

import (
	"52lu/go-import-template/global"
	"52lu/go-import-template/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

// 测试日志记录
func Hello(ctx *gin.Context) {
	global.GvaLogger.Sugar().Infof("日志写入测试: %v",strings.Repeat("hello",6))
	global.GvaLogger.Info("Info记录",zap.String("name","张三"))
	response.OkWithData(ctx,global.GvaConfig)
}

func Test(ctx *gin.Context) {
	response.Ok(ctx)
}
