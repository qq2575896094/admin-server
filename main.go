package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qq2575896094/admin-server/router"
)

func main() {
	r := gin.Default()

	// 设置路由
	router.SetRouter(r)

	r.Run(":8088")
}
