package main

import (
	"52lu/go-import-template/cmd/app"
	"52lu/go-import-template/initialize"
)

func init() {
	// 初始化全局配置文件
	initialize.ViperInit()
}
func main() {
	app.RunServer()
}
