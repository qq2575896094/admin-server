package user

import (
	"github.com/qq2575896094/admin-server/models"
	"github.com/qq2575896094/admin-server/servers"
	"github.com/qq2575896094/admin-server/utils"
	"github.com/redis/go-redis/v9"
)

func (*userHandler) LoginHandler() models.HandlerFunc {
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
