/**
 * @Description elasticSearch配置
 **/
package config

import "time"

/**
 * elastic
 * @Description: ES配置
 **/
type elastic struct {
	Enable              bool          `yaml:"enable"`              // 是否开启
	Url                 string        `yaml:"url"`                 // 服务地址，多个地址用逗号隔开
	Sniff               bool          `yaml:"sniff"`               // 是否转换请求地址，默认为true,当等于true时 请求http://ip:port/_nodes/http，将其返回的url作为请求路径
	HealthCheckInterval time.Duration `yaml:"healthCheckInterval"` // 心跳检测间隔
	LogPre              string        `yaml:"logPre"`              // 日志前缀
}
