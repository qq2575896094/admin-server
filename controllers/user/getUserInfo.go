package user

import (
	"errors"
	"github.com/qq2575896094/admin-server/models"
	"github.com/qq2575896094/admin-server/servers"
	"github.com/qq2575896094/admin-server/utils"
)

func (*userHandler) GetUserInfoHandler() models.HandlerFunc {
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
