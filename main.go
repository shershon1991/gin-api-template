package main

import (
	"52lu/go-import-template/core"
	"52lu/go-import-template/global"
	"52lu/go-import-template/initialize"
)

func init() {
	// 捕获启动时错误
	defer global.CatchError()
	// 程序退出前释放资源
	defer closeResource()

	// 初始化全局配置文件
	initialize.InitConfig()
	// 初始化zap日志
	initialize.InitLogger()
	// 初始化gorm
	initialize.InitGorm()
	// 初始化redis
	initialize.InitRedis()
}

func main() {
	// 启动服务
	core.RunServer()
}
// 程序退出前释放资源
func closeResource()  {
	// 关闭数据库连接
	if global.GvaMysqlClient != nil {
		db, _ := global.GvaMysqlClient.DB()
		_ = db.Close()
	}
}
