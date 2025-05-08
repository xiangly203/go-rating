package entity

type User struct {
	UserName    string `json:"user_name"`
	PhoneNumber string `json:"phone"`
	UUID        string `json:"uuid"`
	Email       string `json:"email"`
	IsAdmin     bool   `json:"is_admin"`
}

