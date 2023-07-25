package conf

type Token struct {
	TokenSigningKey      string `yaml:"tokenSigningKey"`
	TokenExpiresDuration int    `yaml:"tokenExpiresDuration"`
	TokenRefreshDuration int    `yaml:"tokenRefreshDuration"`
}
