package utils

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

	logrus.SetOutput(&logger)
	// TODO: 规范日志格式
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02-03 15:04:05",
	})
	logrus.WithField("app", "admin-server")
}
