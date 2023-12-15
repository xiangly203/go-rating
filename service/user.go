package service

import (
	"errors"
	"go_gin/repository"
)

func CheckUserExistByPhoneOrName(userName string, userPhone string) (bool, error) {
	if _, err := repository.FindUserByNameOrPhoneNumber(userName, userPhone); err != nil {
		return false, errors.New("用户不存在")
	}
	return true, nil
}
