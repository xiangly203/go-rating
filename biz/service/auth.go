package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"go_gin/biz/dal/db"
	model "go_gin/biz/entity/auth"
	entity "go_gin/biz/entity/user"
	"go_gin/config"
	"os"
	"time"
)

func GenerateToken(userInfo entity.UserInfo, tokeType string) (string, error) {
	var tokenExpireDuration time.Duration
	if config.AccessToken == tokeType {
		tokenExpireDuration = config.AccessTokenExpireDuration
	} else if config.RefreshToken == tokeType {
		tokenExpireDuration = config.RefreshTokenExpireDuration
	} else {
		return "", errors.New("token 类型错误")
	}
	err := godotenv.Load(".env.local")
	if err != nil {
		return "", errors.New("载入 .env 文件失败")
	}
	key := os.Getenv("JWT_KEY")
	byteKey := []byte(key)
	expirationTime := time.Now().Add(tokenExpireDuration)
	claims := model.MyCustomClaims{
		User: userInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Subject:   tokeType,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(byteKey)
	return tokenString, err
}

func ParseToken(tokenString string) (*model.MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	}, jwt.WithLeeway(time.Minute*1))
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.MyCustomClaims); ok {
		return claims, nil
	}
	return nil, errors.New("无法解析token")
}

func RefreshToken(refreshTokenString string) (string, error) {
	myCustomClaims, err := ParseToken(refreshTokenString)
	if err != nil {
		return "", err
	}
	if ok, err := CheckUserExistByPhoneOrName(myCustomClaims.User.UserName, ""); err != nil && ok {
		return "", errors.New("用户不存在")
	}
	//if time.Now().After(myCustomClaims.MapClaims.) {
	//	return "", errors.New("Token 过期")
	//}
	token, err := GenerateToken(myCustomClaims.User, config.AccessToken)
	if err != nil {
		return "", err
	}
	return token, nil
}

func LoginCheck(phone string, code string) bool {
	val, err := db.GetRedisVal(phone)
	if err != nil || val != code {
		return false
	}
	return true
}
