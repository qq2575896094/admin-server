package auth

import (
	"github.com/qq2575896094/admin-server/types"
	"net/http"
)

func (*authHandler) Create() types.HandlerFunc {
	return func(c *types.Context) {
		c.String(http.StatusOK, "创建用户功能正在加班加点开发中, 请稍后...")
	}
}
