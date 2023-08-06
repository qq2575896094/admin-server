package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qq2575896094/admin-server/dao"
	"github.com/qq2575896094/admin-server/initialize"
	"github.com/qq2575896094/admin-server/middlewares"
	"github.com/qq2575896094/admin-server/router"
)

func init() {
	// 初始化log文件
	//utils.InitLogConf()

	// 加载配置文件
	initialize.LoadConf("conf", "yaml")

	// 初始化 db servers
	initialize.InitServer()
	// 初始化user collection
	dao.InitUserCollection()
}

func main() {
	r := gin.New()

	r.Use(middlewares.Logger())

	// 设置路由
	router.SetRouter(r)

	r.Run(":8088")
}
