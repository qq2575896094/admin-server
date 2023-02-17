package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qq2575896094/admin-server/controller/user"
	"github.com/qq2575896094/admin-server/middleware/authorization"
)

func setUserRouter(router *gin.Engine) {
	r := router.Group("user")
	userHandler := user.New()
	{
		r.POST("/register", userHandler.Create()) // 注册账号
		r.POST("/login", userHandler.Create())    // 登录账号
		r.POST("/logout", userHandler.Create())   // 退出账号
	}
	r.Use(authorization.CheckTokenAuth())
	{
		r.GET("/getUserInfo", userHandler.GetUserInfo()) // 获取自身信息
		//r.POST("changePassword", baseApi.ChangePassword) // 用户修改密码
		//r.PUT("setUserInfo", baseApi.SetUserInfo)        // 设置用户信息
		//r.PUT("setSelfInfo", baseApi.SetSelfInfo)        // 设置自身信息
		r.POST("resetPassword", func(context *gin.Context) {

		})
	}
}
