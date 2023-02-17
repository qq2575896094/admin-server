package router

import "github.com/gin-gonic/gin"

func setAuthRouter(router *gin.Engine) {
	r := router.Group("auth")
	{
		r.POST("login", func(context *gin.Context) {

		})
		//r.POST("create", authorityApi.CreateAuthority) // 创建角色
		//r.POST("delete", authorityApi.DeleteAuthority) // 删除角色
		//r.PUT("update", authorityApi.UpdateAuthority)  // 更新角色
		//r.POST("copy", authorityApi.CopyAuthority)     // 拷贝角色
	}
}
