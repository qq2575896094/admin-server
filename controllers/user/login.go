package user

import (
	"github.com/qq2575896094/admin-server/models"
	"github.com/qq2575896094/admin-server/servers"
	"github.com/qq2575896094/admin-server/utils"
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

		utils.Ok(&userInfo, c)
	}
}
