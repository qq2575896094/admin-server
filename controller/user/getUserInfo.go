package user

import (
	"github.com/qq2575896094/admin-server/types"
	"github.com/qq2575896094/admin-server/utils/response"
)

type userInfo struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	password string `-`
}

func (*userHandler) GetUserInfo() types.HandlerFunc {
	return func(c *types.Context) {
		response.Ok(userInfo{"雨田", 18, "123"}, c)
	}
}
