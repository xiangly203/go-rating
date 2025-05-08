package db

import (
	"go_gin/internal/domain/entity/user"
)

func CreateUsers(users []*user.User) error {
	return Postgres.Create(users).Error
}

func FindUserByNameOrPhoneNumber(userName string, phoneNumber string) ([]*user.User, error) {
	res := make([]*user.User, 0)
	if err := Postgres.Where(Postgres.Or("user_name = ?", userName).
		Or("phone = ?", phoneNumber)).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
