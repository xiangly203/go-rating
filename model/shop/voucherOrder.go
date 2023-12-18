package model

import (
	"gorm.io/gorm"
	"time"
)

type VoucherOrder struct {
	gorm.Model
	VoucherOrderID string    `json:"order_info_id" column:"order_info_id"`
	UserID         string    `json:"user_id" column:"user_id"`
	TotalAmount    float64   `json:"total_amount" column:"total_amount"`
	ReceiverId     string    `json:"receiver_id" column:"receiver_id"`
	PayStatus      int       `json:"pay_status" column:"pay_status"`
	PayAmount      float64   `json:"pay_amount" column:"pay_amount"`
	Status         int       `json:"order_status" column:"order_status"`
	UseTime        time.Time `json:"use_time" column:"use_time"`
	RefundTime     time.Time `json:"refund_time" column:"refund_time"`
}

func (u *VoucherOrder) TableName() string {
	return "voucher_order"
}
