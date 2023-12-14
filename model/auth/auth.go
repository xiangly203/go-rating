package model

import (
	"github.com/golang-jwt/jwt/v5"
	"go_gin/model/user"
)

type GetCodeReq struct {
	Phone string `json:"phone"`
}

type GetCodeResp struct {
	Code string `json:"code"`
}

type LoginReq struct {
	GetCodeReq
	Code string `json:"code"`
}

type LoginResp struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type MyCustomClaims struct {
	User user.UserInfo `json:"user"`
	jwt.RegisteredClaims
}
