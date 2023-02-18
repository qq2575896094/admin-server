package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRouter(router *gin.Engine) {
	// 设置静态资源
	router.Static("/favicon.ico", "./public/favicon.ico")
	router.StaticFS("/static", http.Dir("public"))
	router.StaticFile("/", "./public/index.html")

	// 设置user组路由
	setUserRouter(router)
}
