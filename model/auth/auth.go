package model

import (
	"context"
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
type LoginService interface {
	GetCode(ctx context.Context, phone string) (string, error)
	Login(ctx context.Context, req LoginReq) (LoginResp error)
	GenerateToken(ctx context.Context, userInfo user.UserInfo, tokeType string) (string, error)
	ParseToken(ctx context.Context, tokenString string) (string, error)
	RefreshToken(ctx context.Context, refreshTokenString string) (string, error)
}
