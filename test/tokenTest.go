package test

import (
	"fmt"
	"go_gin/model/user"
	"go_gin/utils"
)

func TokenTest() {
	user := user.UserInfo{
		UserName:  "test",
		UserPhone: "12345678900",
	}
	accessToken, err := utils.GenerateToken(user, "accessToken")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(accessToken)
	refreshToken, err := utils.GenerateToken(user, "refreshToken")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(refreshToken)
	myClaims, err := utils.ParseToken(accessToken)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myClaims)

	refreshedToken, err := utils.RefreshToken(refreshToken)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(refreshedToken)
}
