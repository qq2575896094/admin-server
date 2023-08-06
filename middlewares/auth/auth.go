package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/qq2575896094/admin-server/models"
	"github.com/qq2575896094/admin-server/servers"
	"github.com/qq2575896094/admin-server/utils"
)

func CheckTokenAuth() models.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		if token == "" {
			utils.FailWithMessage("Token不能为空", c)
			c.Abort()
			return
		}

		// TODO: 频繁刷新, 影响性能
		userId, nToken, err := utils.New().RefreshToken(token)

		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				utils.FailAuthorization("授权过期, 请重新登录", c)
				c.Abort()
				return
			}
			utils.FailAuthorization("非法Token, 请重新登录", c)
			c.Abort()
			return
		}

		if err := servers.NewJwtService(c).SetToken(userId, nToken); err != nil {
			utils.FailWithMessage("设置登录状态失败", c)
			c.Abort()
			return
		}

		c.Next()
	}
}
