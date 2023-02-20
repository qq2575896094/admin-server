package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qq2575896094/admin-server/db"
	"github.com/qq2575896094/admin-server/middleware/log"
	"github.com/qq2575896094/admin-server/router"
	"github.com/qq2575896094/admin-server/utils"
)

func init() {
	// 加载配置文件
	if err := utils.LoadConf("conf", "yaml"); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 初始化log文件
	utils.InitLogConf()

	// 初始化mongodb server
	db.InitServer()
}

func main() {
	r := gin.New()
	r.Use(log.Logger())

	// 设置路由
	router.SetRouter(r)

	r.Run(":8088")
}
