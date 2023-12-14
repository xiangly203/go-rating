package model

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	ShopID      string `json:"shop_id" column:"shop_id"`
	ShopName    string `json:"shop_name" column:"shop_name"`
	Description string `json:"description" column:"description"`
	Status      int    `json:"status" column:"status"`
}

func (u *Product) Shop() string {
	return "shops"
}
