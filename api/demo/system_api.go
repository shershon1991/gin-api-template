package demo

import (
	"52lu/go-import-template/global"
	"52lu/go-import-template/model/response"
	"github.com/gin-gonic/gin"
)

// 配置信息
func GetConfig(ctx *gin.Context)  {
	response.OkWithData(ctx,global.GvaConfig)
}