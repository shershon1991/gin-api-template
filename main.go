package main

import (
	"52lu/go-import-template/cmd/app"
	"52lu/go-import-template/initialize"
)

func init() {
	// 初始化全局配置文件
	initialize.ViperInit()
	// 初始化zap日志
	initialize.InitZap()
}
func main() {
	app.RunServer()
}
