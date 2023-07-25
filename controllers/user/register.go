package user

import (
	"github.com/qq2575896094/admin-server/models"
	"github.com/qq2575896094/admin-server/servers"
	"github.com/qq2575896094/admin-server/utils"
)

func (*userHandler) RegisterHandler() models.HandlerFunc {
	return func(c *models.Context) {
		var user models.UserRegisterParams

		if err := c.ShouldBind(&user); err != nil {
			utils.FailWithMessage(err.Error(), c)
			return
		}

		resp, err := servers.NewUserServer().SignUp(&user)
		if err != nil {
			utils.FailWithMessage(err.Error(), c)
			return
		}

		utils.Ok(resp, c)
	}
}
