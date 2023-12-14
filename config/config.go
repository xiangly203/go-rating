package config

import "time"

//import (
//	"time"
//)

const (
	CodeTTL                    = 2 * time.Minute
	RefreshTokenExpireDuration = time.Hour * 2
	AccessTokenExpireDuration  = time.Hour * 12
	RespOK                     = 200
	RespErrWithToken           = 5001
	RespErrWithPhoneOrCode     = 5002
	RespErrWithServer          = 5003
)
