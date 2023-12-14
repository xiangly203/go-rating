package config

import "time"

const (
	CodeTTL                    = 2 * time.Minute
	RefreshTokenExpireDuration = time.Second * 10
	AccessTokenExpireDuration  = time.Hour * 12
	AccessToken                = "accessToken"
	RefreshToken               = "refreshToken"
	RespOK                     = 200
	RespErrWithToken           = 5001
	RespErrWithPhoneOrCode     = 5002
	RespErrWithServer          = 5003
	RespErrTokenExpire         = 5004
)
