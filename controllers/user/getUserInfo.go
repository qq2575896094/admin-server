package user

import (
	"github.com/qq2575896094/admin-server/models"
	"github.com/qq2575896094/admin-server/utils"
)

type userInfo struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	password string `-`
}

func (*userHandler) GetUserInfoHandler() models.HandlerFunc {
	return func(c *models.Context) {
		utils.Ok(userInfo{"雨田", 18, "123"}, c)
	}
}
