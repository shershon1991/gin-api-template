package main

import (
	"52lu/go-import-template/core"
	"52lu/go-import-template/initialize"
)

func init() {
	// 初始化全局配置文件
	initialize.InitConfig()
	// 初始化zap日志
	initialize.InitLogger()
}
func main() {
	// 启动服务
	core.RunServer()
}
