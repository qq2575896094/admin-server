package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qq2575896094/admin-server/router"
	"github.com/qq2575896094/admin-server/utils"
)

func init() {
	// 加载配置文件
	if err := utils.LoadConf("conf", "yaml"); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func main() {
	r := gin.Default()

	// 设置路由
	router.SetRouter(r)

	r.Run(":8088")
}
