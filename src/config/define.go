package config

import (
	"fmt"
)

// ServerConfig 服务器配置
type ServerConfig struct {
	Host string `toml:"host"` // 服务器地址
	Port int    `toml:"port"` // 服务器端口
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host            string `toml:"host"`              // 数据库地址
	Port            int    `toml:"port"`              // 数据库端口
	User            string `toml:"user"`              // 数据库用户名
	Password        string `toml:"password"`          // 数据库密码
	DBName          string `toml:"dbname"`            // 数据库名称
	SSLMode         string `toml:"sslmode"`           // SSL 模式
	MaxOpenConns    int    `toml:"max_open_conns"`    // 最大连接数
	MaxIdleConns    int    `toml:"max_idle_conns"`    // 最大空闲连接数
	ConnMaxLifetime int    `toml:"conn_max_lifetime"` // 连接最大生命周期 单位秒
}

// GetDSN 获取数据库连接字符串
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}

// RedisConfig Redis 配置
type RedisConfig struct {
	Host     string `toml:"host"`     // Redis 地址
	Port     int    `toml:"port"`     // Redis 端口
	Username string `toml:"username"` // Redis 用户名
	Password string `toml:"password"` // Redis 密码
	DB       int    `toml:"db"`       // Redis 数据库索引
	Scheme   string `toml:"scheme"`   // Redis 模式名称
}

// Config 总配置结构
type Config struct {
	Server   ServerConfig   `toml:"server"`   // 服务器配置
	Database DatabaseConfig `toml:"database"` // 数据库配置
	Redis    RedisConfig    `toml:"redis"`    // Redis 配置
}
