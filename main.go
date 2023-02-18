package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qq2575896094/admin-server/middleware/log"
	"github.com/qq2575896094/admin-server/router"
	"github.com/qq2575896094/admin-server/utils"
	"github.com/sirupsen/logrus"
)

func init() {
	// 加载配置文件
	if err := utils.LoadConf("conf", "yaml"); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 初始化log文件
	utils.InitLogConf()
}

func main() {
	r := gin.New()
	r.Use(log.Logger())

	// 设置路由
	router.SetRouter(r)

	logrus.Infoln("welcome to admin server~~")

	r.Run(":8088")
}
