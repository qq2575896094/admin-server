package user

import (
	"github.com/qq2575896094/admin-server/models"
	"github.com/qq2575896094/admin-server/utils"
	"net/http"
)

type userParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (*userHandler) LoginHandler() models.HandlerFunc {
	return func(c *models.Context) {
		var user userParams
		if err := c.ShouldBind(&user); err != nil {
			utils.FailWithMessage(err.Error(), c)
			return
		}

		c.String(http.StatusOK, "创建用户功能正在加班加点开发中, 请稍后...")
	}
}
