package demo

import (
	"github.com/gin-gonic/gin"
	"shershon1991/gin-api-template/global"
	"shershon1991/gin-api-template/model/response"
)

// 配置信息
func GetConfig(ctx *gin.Context) {
	response.OkWithData(ctx, global.GvaConfig)
}
