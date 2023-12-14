package usecase

import (
	"errors"
	"go_gin/repository/mysql"
)

func CheckUserExistByPhoneOrName(userName string, userPhone string) (bool, error) {
	if _, err := mysql.FindUserByNameOrPhoneNumber(userName, userPhone); err != nil {
		return false, errors.New("用户不存在")
	}
	return true, nil
}
