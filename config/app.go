// Package config /**
package config

// 应用信息
type app struct {
	Desc       string `yaml:"desc"`
	Addr       string `yaml:"addr"`
	ConfigFile string `yaml:"configFile"`
	Version    string `yaml:"version"`
	Env        string `yaml:"env"`
}

// MySQL信息
type mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// Redis
type redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// ServerConfig 配置信息
type ServerConfig struct {
	App   app   `yaml:"app"`
	Mysql mysql `yaml:"mysql"`
	Redis redis `yaml:"redis"`
	Log   log   `yaml:"log"`
}
