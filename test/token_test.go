package test

import (
	"fmt"
	"go_gin/model/user"
	"go_gin/usecase"
)

func TokenTest() {
	user := user.UserInfo{
		UserName:  "test",
		UserPhone: "12345678900",
	}
	accessToken, err := usecase.GenerateToken(user, "accessToken")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(accessToken)
	//refreshToken, err := utils.GenerateToken(user, "refreshToken")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(refreshToken)
	//time.Sleep(time.Minute * 5)
	myClaims, err := usecase.ParseToken(accessToken)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myClaims)

	//refreshedToken, err := utils.RefreshToken(refreshToken)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(refreshedToken)
}
