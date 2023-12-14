package mysql

import (
	"go_gin/model/user"
)

func CreateUsers(users []*user.User) error {
	return DB.Create(users).Error
}

func FindUserByNameOrPhoneNumber(userName string, phoneNumber string) ([]*user.User, error) {
	res := make([]*user.User, 0)
	if err := DB.Where(DB.Or("user_name = ?", userName).
		Or("phone_number = ?", phoneNumber)).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
