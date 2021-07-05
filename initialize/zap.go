/**
 * @Author Mr.LiuQH
 * @Description 初始化zap日志
 * @Date 2021/7/5 5:54 下午
 **/
package initialize

import (
	"52lu/go-import-template/global"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const (
	// 写入位置
	writeFile       = "file"
	writeStd        = "std"
	writeStdAndFile = "all"
	// 日志输出格式
	outJson    = "json"
	outConsole = "console"
)

// 获取最低记录日志级别
func getLevel() zapcore.Level {
	levelMap := map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dpanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}
	if level, ok := levelMap[global.GvaConfig.Log.Level]; ok {
		return level
	}
	return zapcore.DebugLevel
}

// 初始化Logger
func InitZap() {
	logConfig := global.GvaConfig.Log
	// 打开日志文件
	fileHandle, err := os.Create(logConfig.File)
	if err != nil {
		panic(fmt.Sprintf("创建日志文件失败: %s", err))
	}
	defer fileHandle.Close()
	// 设置输出格式
	var encoder zapcore.Encoder
	if logConfig.OutFormat == outJson {
		encoder = zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	} else {
		encoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	}
	// 设置日志文件切割
	writeSyncer := zapcore.AddSync(getLumberjackLogger())
	// 创建NewCore
	zapCore := zapcore.NewCore(encoder,writeSyncer, getLevel())
	// 创建logger
	logger := zap.New(zapCore)
	defer logger.Sync()
	global.GvaLogger = logger
}

// 获取文件切割和归档配置信息
func getLumberjackLogger() *lumberjack.Logger {
	lumberjackConfig := global.GvaConfig.Log.LumberJack
	lumberjackLogger := &lumberjack.Logger{
		Filename:   global.GvaConfig.Log.File,   //日志文件
		MaxSize:    lumberjackConfig.MaxSize,    //单文件最大容量(单位MB)
		MaxBackups: lumberjackConfig.MaxBackups, //保留旧文件的最大数量
		MaxAge:     lumberjackConfig.MaxAge,     // 旧文件最多保存几天
		Compress:   lumberjackConfig.Compress,   // 是否压缩/归档旧文件
	}
	return lumberjackLogger
}
