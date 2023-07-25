package user

import (
	"github.com/qq2575896094/admin-server/models"
	"net/http"
)

func (*userHandler) LogoutHandler() models.HandlerFunc {
	return func(c *models.Context) {
		c.String(http.StatusOK, "创建用户功能正在加班加点开发中, 请稍后...")
	}
}
