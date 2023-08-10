package servers

import (
	"context"
	"github.com/qq2575896094/admin-server/conf"
	"github.com/qq2575896094/admin-server/constants"
	"github.com/qq2575896094/admin-server/models"
	"strings"
	"time"
)

type JwtService struct {
	context context.Context
	ctx     *models.Context
}

// GetRedisToken 获取redis token
func (j *JwtService) GetRedisToken(userId string) (token string, err error) {
	return constants.RedisClient.Get(j.context, userId).Result()
}

// SetRedisToken 设置token
func (j *JwtService) SetRedisToken(userId string, token string) error {
	dur := time.Duration(conf.Config.Token.TokenExpiresDuration) * time.Second
	return constants.RedisClient.Set(context.Background(), userId, token, dur).Err()
}

// SetToken 设置token
func (j *JwtService) SetToken(userId string, token string) error {
	if err := j.SetRedisToken(userId, token); err != nil {
		return err
	}
	secure := j.IsHttps()
	j.ctx.SetCookie("Token", token, conf.Config.Token.TokenExpiresDuration, "/", "", secure, true)
	return nil
}

// IsHttps 判断是否是https请求
func (j *JwtService) IsHttps() bool {
	if strings.HasPrefix(j.ctx.GetHeader("URL"), "https") || j.ctx.Request.TLS != nil {
		return true
	}
	return false
}

func NewJwtService(ctx *models.Context) *JwtService {
	return &JwtService{context: constants.RedisContext, ctx: ctx}
}
