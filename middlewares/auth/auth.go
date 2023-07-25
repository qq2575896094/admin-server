package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/qq2575896094/admin-server/conf"
	"github.com/qq2575896094/admin-server/models"
	"github.com/qq2575896094/admin-server/utils"
	"strings"
)

func CheckTokenAuth() models.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		if token == "" {
			utils.FailWithMessage("Token不能为空", c)
			c.Abort()
			return
		}

		nToken, err := utils.New().RefreshToken(token)

		if err != nil {
			if err == utils.TokenExpired {
				utils.FailAuthorization("授权过期, 请重新登录", c)
				c.Abort()
				return
			}
			utils.FailAuthorization("非法Token, 请重新登录", c)
			c.Abort()
			return
		}

		setToken(c, nToken)
		c.Next()
	}
}

func setToken(c *gin.Context, token string) {
	secure := IsHttps(c)
	c.Header("Token", token)
	c.SetCookie("Token", token, conf.Config.Token.TokenExpiresDuration, "/", "", secure, true)
}

// IsHttps 判断是否是https请求
func IsHttps(c *gin.Context) bool {
	if strings.HasPrefix(c.GetHeader("URL"), "https") || c.Request.TLS != nil {
		return true
	}
	return false
}
