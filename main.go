package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qq2575896094/admin-server/dao"
	"github.com/qq2575896094/admin-server/middlewares/log"
	"github.com/qq2575896094/admin-server/router"
	"github.com/qq2575896094/admin-server/utils"
)

func init() {
	// 初始化log文件
	//utils.InitLogConf()

	// 加载配置文件
	utils.LoadConf("conf", "yaml")

	// 初始化 db servers
	utils.InitServer()
	// 初始化user collection
	dao.InitUserCollection()
}

func main() {
	r := gin.New()

	r.Use(log.Logger())

	// 设置路由
	router.SetRouter(r)

	r.Run(":8088")
}
