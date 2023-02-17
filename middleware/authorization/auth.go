package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/qq2575896094/admin-server/model/response"
	"github.com/qq2575896094/admin-server/types"
	"github.com/qq2575896094/admin-server/utils/jwtToken"
	"github.com/spf13/viper"
	"time"
)

func CheckTokenAuth() types.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		if token == "" {
			response.FailWithMessage("未登录或非法访问", c)
			c.Abort()
			return
		}

		jwtSign := jwtToken.New()

		// 校验token合法性
		clamis, err := jwtSign.ParseToken(token)
		if err != nil {
			if err == jwtToken.TokenExpired {
				response.FailAuthorization("授权过期, 请重新登录", c)
			}
			response.FailAuthorization("非法token, 请重新登录", c)
			c.Abort()
			return
		}

		// 在最后的一小时过期时间内, 更新token
		if clamis.ExpiresAt.Unix()-time.Now().Unix() < int64(viper.GetDuration("token.tokenRefreshDuration")*time.Second) {
			nToken, _ := jwtSign.UpdateToken(clamis)
			clamis.ExpiresAt = jwtToken.FormatJwtExpiresTime(viper.GetDuration("token.tokenExpiresDuration") * time.Second)
			// TODO: 后面存入redis的时候, 也需要更新过期时间
			c.Header("Token", nToken)
		}

		c.Next()
	}
}
