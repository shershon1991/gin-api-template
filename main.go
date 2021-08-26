package main

import (
	"52lu/go-import-template/global"
	"52lu/go-import-template/initialize"
)

func main() {
	// 程序退出前释放资源
	defer closeResource()
	// 加载启动前配置
	initialize.SetLoadInit()
	// 启动服务
	RunServer()
}
// 程序退出前释放资源
func closeResource()  {
	// 关闭数据库连接
	if global.GvaMysqlClient != nil {
		db, _ := global.GvaMysqlClient.DB()
		_ = db.Close()
	}
}
