package utils

import (
	"github.com/qq2575896094/admin-server/conf"
	"github.com/spf13/viper"
	"os"
)

// LoadConf 加载配置文件
func LoadConf(configName string, configType string) {
	workDir, _ := os.Getwd()
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(workDir)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&conf.Config)
	if err != nil {
		panic(err)
	}
}
