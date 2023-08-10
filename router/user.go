package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qq2575896094/admin-server/controllers"
	"github.com/qq2575896094/admin-server/middlewares"
)

func setUserRouter(router *gin.Engine) {
	r := router.Group("user")
	userHandler := controllers.New()
	{
		r.POST("/register", userHandler.Register()) // 注册账号
		r.POST("/login", userHandler.Login())       // 登录账号
	}

	r.Use(middlewares.CheckAuth())
	{
		r.POST("/getUserInfo", userHandler.GetUserInfo())       // 获取自身信息
		r.POST("/changePassword", userHandler.ChangePassword()) // 用户修改密码
		//r.PUT("setUserInfo", userHandler.SetUserInfo)        				// 设置用户信息
		r.POST("/logout", userHandler.Register()) // 退出账号
	}
}
