package utils

import (
	"github.com/spf13/viper"
)

// LoadConf 加载配置文件
func LoadConf(configName string, configType string) error {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath("configs/")

	return viper.ReadInConfig()
}
