package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/qq2575896094/admin-server/conf"
	"time"
)

type JwtCustomClaims struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JwtSign struct {
	SingingKey []byte
}

var TokenExpired = errors.New("token is expired")

func FormatJwtExpiresTime(dur int) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(dur)))
}

// CreateClaims 创建claims
func (*JwtSign) CreateClaims(userId string, username string) *JwtCustomClaims {
	expiresTime := FormatJwtExpiresTime(conf.Config.Token.TokenExpiresDuration)

	return &JwtCustomClaims{
		UserId:   userId,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "jwt",                          // 签发人
			ExpiresAt: expiresTime,                    // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
			Subject:   "Token",                        // 主题
		},
	}
}

// CreateToken 创建Token
func (j *JwtSign) CreateToken(userId string, username string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, j.CreateClaims(userId, username)).SignedString(j.SingingKey)
}

// ParseToken 解析(验证)Token
func (j *JwtSign) ParseToken(token string) (*JwtCustomClaims, error) {
	claims := JwtCustomClaims{}
	tokenClaims, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return j.SingingKey, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*JwtCustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// RefreshToken 更新token
func (j *JwtSign) RefreshToken(token string) (nToken string, err error) {
	tokenClaim, err := j.ParseToken(token)
	if err != nil {
		return
	}

	if tokenClaim.ExpiresAt.Unix() > time.Now().Unix() {
		return j.CreateToken(tokenClaim.UserId, tokenClaim.Username)
	}

	return "", TokenExpired
}

// New 创建JwtSign实例
func New() *JwtSign {
	return &JwtSign{
		SingingKey: []byte(conf.Config.Token.TokenSigningKey),
	}
}
