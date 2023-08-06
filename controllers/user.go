package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/qq2575896094/admin-server/models"
	"github.com/qq2575896094/admin-server/servers"
	"github.com/qq2575896094/admin-server/utils"
	"github.com/redis/go-redis/v9"
	"net/http"
)

type User interface {
	Register() models.HandlerFunc
	Login() models.HandlerFunc
	Logout() models.HandlerFunc
	GetUserInfo() models.HandlerFunc
	ChangePassword() models.HandlerFunc
}

type userHandler struct{}

func New() User {
	return &userHandler{}
}

// Register 注册
func (*userHandler) Register() models.HandlerFunc {
	return func(c *models.Context) {
		var user models.UserRegisterParams

		if err := c.ShouldBindJSON(&user); err != nil {
			utils.FailWithMessage(err.Error(), c)
			return
		}

		userInfo, err := servers.NewUserServer().SignUp(&user)
		if err != nil {
			utils.FailWithMessage(err.Error(), c)
			return
		}

		utils.Ok(userInfo, c)
	}
}

func (*userHandler) Login() models.HandlerFunc {
	return func(c *models.Context) {
		var user models.UserLoginParams
		if err := c.ShouldBindJSON(&user); err != nil {
			utils.FailWithMessage(err.Error(), c)
			return
		}

		userInfo, err := servers.NewUserServer().Login(&user)
		if err != nil {
			utils.FailWithMessage(err.Error(), c)
			return
		}

		// 签发jwt
		signToken(c, userInfo)
	}
}

// Logout 登出
func (*userHandler) Logout() models.HandlerFunc {
	return func(c *models.Context) {
		c.String(http.StatusOK, "创建用户功能正在加班加点开发中, 请稍后...")
	}
}

// GetUserInfo 获取用户信息
func (*userHandler) GetUserInfo() models.HandlerFunc {
	return func(c *models.Context) {
		userId, _ := c.Get("userId")

		id, ok := userId.(string)
		if !ok {
			utils.FailWithMessage(errors.New("获取用户信息失败").Error(), c)
			return
		}

		userInfo, err := servers.NewUserServer().GetUserInfo(id)
		if err != nil {
			utils.FailWithMessage(err.Error(), c)
			return
		}

		utils.Ok(&userInfo, c)
	}
}

func (*userHandler) ChangePassword() models.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// signToken 签发Token
func signToken(c *models.Context, userInfo *models.UserInfo) {
	token, err := utils.New().CreateToken(userInfo.Id, userInfo.Username)
	if err != nil {
		utils.FailWithMessage("获取Token失败", c)
		return
	}

	jwtService := servers.NewJwtService(c)
	if _, err := jwtService.GetRedisToken(userInfo.Id); err == redis.Nil { // 获取redis token
		if err := jwtService.SetToken(userInfo.Id, token); err != nil {
			utils.FailWithMessage("设置登录状态失败", c)
			return
		}

		utils.Ok(nil, c)
	} else {
		if err := jwtService.SetToken(userInfo.Id, token); err != nil {
			utils.FailWithMessage("设置登录状态失败", c)
			return
		}
		utils.Ok(nil, c)
	}
}
