package initialize

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitLogConf() {
	logger := lumberjack.Logger{
		Filename:   viper.GetString("log.filename"),
		MaxSize:    viper.GetInt("log.maxSize"), // megabytes
		MaxBackups: viper.GetInt("log.maxBackups"),
		MaxAge:     viper.GetInt("log.maxAge"),    //days
		Compress:   viper.GetBool("log.compress"), // disabled by default
	}
	logrus.WithFields(logrus.Fields{
		"servers": "admin-servers",
	})
	logrus.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: false,
	})
	logrus.SetOutput(&logger)
}
