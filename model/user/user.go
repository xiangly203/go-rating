package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName    string `json:"user_name" column:"user_name"`
	PhoneNumber string `json:"phone" column:"phone"`
	UUID        string `json:"uuid" column:"uuid"`
	Email       string `json:"email" column:"email"`
	IsAdmin     bool   `json:"is_admin" column:"is_admin"`
	IsDelete    bool   `json:"is_delete" column:"is_delete"`
}

func (u *User) TableName() string {
	return "users"
}
