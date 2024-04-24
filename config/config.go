package config

import "time"

const (
	AccessToken     = "accessToken"
	RefreshToken    = "refreshToken"
	CodeTTL         = 2 * time.Minute
	DefaultCacheTTL = 5 * time.Minute // CacheTTL

	RefreshTokenExpireDuration = time.Second * 10
	AccessTokenExpireDuration  = time.Hour * 12

	LenLogid = 32
)
