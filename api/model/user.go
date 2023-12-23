package model

import "github.com/golang-jwt/jwt/v5"

type UserInfo struct {
	UserName  string `json:"user_name"`
	UserPhone string `json:"user_phone"`
}

type MyCustomClaims struct {
	User UserInfo `json:"user"`
	jwt.RegisteredClaims
}
