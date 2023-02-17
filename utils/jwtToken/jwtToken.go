package jwtToken

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

type JwtCustomClaims struct {
	UserId string
	jwt.RegisteredClaims
}

type JwtSign struct {
	SingingKey []byte
}

var (
	TokenInvalid = errors.New("token is invalid")
	TokenExpired = errors.New("token is expired")
)

func FormatJwtExpiresTime(dur time.Duration) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(dur))
}

// CreateClaims 创建claims
func (*JwtSign) CreateClaims(userId string) *JwtCustomClaims {
	expiresTime := FormatJwtExpiresTime(viper.GetDuration("Token.TokenExpiresDuration") * time.Second)

	return &JwtCustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "jwt",                          // 签发人
			ExpiresAt: expiresTime,                    // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
			Subject:   "Token",                        // 主题
		},
	}
}

// CreateToken 创建Token
func (j *JwtSign) CreateToken(claims *JwtCustomClaims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.SingingKey)
}

// UpdateToken 更新token
func (j *JwtSign) UpdateToken(claims *JwtCustomClaims) (string, error) {
	// TODO: 需要优化, 当大量请求过来时, 频繁更新token, 会影响性能
	return j.CreateToken(claims)

}

// ParseToken 解析Token
func (j *JwtSign) ParseToken(tokenString string) (*JwtCustomClaims, error) {
	claims := JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return j.SingingKey, nil
	})

	if err != nil {
		if e, ok := err.(*jwt.ValidationError); ok {
			if e.Errors&jwt.ValidationErrorExpired != 0 { // token 过期
				return nil, TokenExpired
			}
			return nil, TokenInvalid
		}
	}

	if claims, ok := token.Claims.(*JwtCustomClaims); token != nil && ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// New 创建JwtSign实例
func New() *JwtSign {
	return &JwtSign{
		SingingKey: []byte(viper.GetString("Token.TokenSigningKey")),
	}
}
