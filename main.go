package main

import (
	"52lu/go-import-template/core"
	"52lu/go-import-template/global"
	"52lu/go-import-template/initialize"
	"fmt"
)

func init() {
	// 初始化全局配置文件
	initialize.InitConfig()
	// 初始化zap日志
	initialize.InitLogger()
	// 初始化mysql
	initialize.InitMysql()
}
func main() {
	// 启动服务
	core.RunServer()
	// 程序退出前释放资源
	defer closeResource()
}

// 程序退出前释放资源
func closeResource()  {
	// 关闭数据库连接
	if global.GvaMysqlClient != nil {
		fmt.Println("程序退出！")
		db, _ := global.GvaMysqlClient.DB()
		_ = db.Close()
	}
}