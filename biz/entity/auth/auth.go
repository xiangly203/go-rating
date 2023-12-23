package entity

import (
	"github.com/golang-jwt/jwt/v5"
	entity "go_gin/biz/entity/user"
)

type MyCustomClaims struct {
	User entity.UserInfo `json:"user"`
	jwt.RegisteredClaims
}

//type LoginService interface {
//	GetCode(c context.Context, phone string) (string, error)
//	Login(c context.Context, req model.LoginReq) (LoginResp error)
//	GenerateToken(c context.Context, userInfo entity.UserInfo, tokeType string) (string, error)
//	ParseToken(c context.Context, tokenString string) (string, error)
//	RefreshToken(c context.Context, refreshTokenString string) (string, error)
//}
