/**
 * @Description redis验证使用
 **/
package demo

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"shershon1991/gin-api-template/global"
	"strings"
)

func LogTest(ctx *gin.Context) {
	// Sugar模式
	global.GvaLogger.Sugar().Infof("日志写入测试: %v", strings.Repeat("hello", 6))
	// 默认模式
	global.GvaLogger.Info("Info记录", zap.String("name", "张三"))
}
