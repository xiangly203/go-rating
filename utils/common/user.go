package utils

import (
	"errors"
	"go_gin/utils/dal/mysql"
)

func CheckUserExistByPhoneOrName(userName string, userPhone string) (bool, error) {
	if _, err := mysql.FindUserByNameOrPhoneNumber(userName, userPhone); err != nil {
		return false, errors.New("用户不存在")
	}
	return true, nil
}
