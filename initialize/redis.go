// Package initialize /**
package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"shershon1991/gin-api-template/global"
)

// InitRedis 初始化redis客户端
func initRedis() {
	if !global.GvaConfig.Redis.Enable {
		return
	}
	// 创建
	redisClient := redis.NewClient(&redis.Options{
		Addr:     global.GvaConfig.Redis.Addr,
		Password: global.GvaConfig.Redis.Password,
		DB:       global.GvaConfig.Redis.DefaultDB,
	})
	// 使用超时上下文，验证redis
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), global.GvaConfig.Redis.DialTimeout)
	defer cancelFunc()
	_, err := redisClient.Ping(timeoutCtx).Result()
	if err != nil {
		panic("redis初始化失败! " + err.Error())
	}
	global.GvaRedis = redisClient
}
