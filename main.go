package main

import (
	"52lu/go-import-template/cmd/app"
	"52lu/go-import-template/init"
)

func init() {
	// 初始化全局配置文件
	init.ViperInit()
}
func main() {
	app.RunServer()
}
