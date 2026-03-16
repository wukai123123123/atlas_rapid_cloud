package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var globalConfig *Config

// LoadConfig 加载配置文件
// configPath: 配置文件路径，如果为空则使用默认的 config.toml
func LoadConfig(configPath string) (*Config, error) {
	if configPath == "" {
		return nil, fmt.Errorf("请指定配置文件路径")
	}

	config := &Config{}
	if _, e := toml.DecodeFile(configPath, config); e != nil {
		return nil, fmt.Errorf("读取配置文件失败 [%s]: %w", configPath, e)
	}

	globalConfig = config
	return config, nil
}

// GetConfig 获取全局配置实例
func GetConfig() *Config {
	if globalConfig == nil {
		panic("配置尚未加载")
	}
	return globalConfig
}
