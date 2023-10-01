package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qq2575896094/admin-server/controllers"
	"github.com/qq2575896094/admin-server/middlewares"
)

func setUserRouter(router *gin.Engine) {
	userHandler := controllers.New()
	r := router.Group("")
	{
		r.POST("/user/register", userHandler.Register()) // 注册账号
		r.POST("/user/login", userHandler.Login())       // 登录账号
	}

	user := router.Group("user").Use(middlewares.CheckAuth())
	{
		user.POST("/getUserInfo", userHandler.GetUserInfo())       // 获取自身信息
		user.POST("/changePassword", userHandler.ChangePassword()) // 用户修改密码
		//user.PUT("setUserInfo", userHandler.SetUserInfo)        				// 设置用户信息
		user.POST("/logout", userHandler.Register()) // 退出账号
	}
}
