package service

import (
	"errors"
	"go_gin/biz/dal/db"
)

func CheckUserExistByPhoneOrName(userName string, userPhone string) (bool, error) {
	if _, err := db.FindUserByNameOrPhoneNumber(userName, userPhone); err != nil {
		return false, errors.New("用户不存在")
	}
	return true, nil
}
