package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"go_gin/config"
	"go_gin/model/user"
	utils "go_gin/utils/common"
	"os"
	"time"
)

type MyCustomClaims struct {
	User user.UserInfo `json:"user"`
	jwt.RegisteredClaims
}

func GenerateToken(userInfo user.UserInfo, tokeType string) (string, error) {
	var tokenExpireDuration time.Duration
	if tokeType == "accessToken" {
		tokenExpireDuration = config.AccessTokenExpireDuration
	} else if tokeType == "refreshToken" {
		tokenExpireDuration = config.RefreshTokenExpireDuration
	} else {
		return "", errors.New("token 类型错误")
	}
	err := godotenv.Load(".env")
	if err != nil {
		return "", errors.New("载入 .env 文件失败")
	}
	key := os.Getenv("JWT_KEY")
	byteKey := []byte(key)
	expirationTime := time.Now().Add(tokenExpireDuration)
	claims := &MyCustomClaims{
		userInfo,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Subject:   tokeType,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(byteKey)
	return tokenString, err
}

func ParseToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok {
		return claims, nil
	}
	return nil, errors.New("无法解析token")
}

func RefreshToken(refreshTokenString string) (string, error) {
	myCustomClaims, err := ParseToken(refreshTokenString)
	if err != nil {
		return "", err
	}
	if ok, err := utils.CheckUserExistByPhoneOrName(myCustomClaims.User.UserName, ""); err != nil && ok {
		return "", errors.New("用户不存在")
	}
	token, err := GenerateToken(myCustomClaims.User, "accessToken")
	if err != nil {
		return "", err
	}
	return token, nil
}
