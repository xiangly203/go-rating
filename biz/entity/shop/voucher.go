package entity

import "gorm.io/gorm"

type Voucher struct {
	gorm.Model
	VoucherID   string  `json:"voucher_id" gorm:"column:voucher_id"`
	ShopId      string  `json:"shop_id" gorm:"column:shop_id"`
	Title       string  `json:"title" gorm:"column:title"`
	SubTitle    string  `json:"sub_title" gorm:"column:sub_title"`
	Rules       string  `json:"rules" gorm:"column:rules"`
	PayValue    float64 `json:"pay_value" gorm:"column:pay_value"`
	ActualValue float64 `json:"actual_value" gorm:"column:actual_value"`
	Type        string  `json:"type" gorm:"column:type"`
	Status      string  `json:"status" gorm:"column:status"`
}

func (u *Voucher) TableName() string {
	return "voucher"
}
