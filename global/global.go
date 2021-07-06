package global

import (
	"52lu/go-import-template/config"
	"go.uber.org/zap"
)

// 常量
const (
	RootDir    = "./"                // 根目录
	ConfigFile = "./config/app.yaml" // 配置文件
)

// 变量
var (
	GvaConfig config.ServerConfig // 全局配置
	GvaLogger *zap.Logger         // 日志
)
