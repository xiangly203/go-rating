package entity

import (
	"time"
)

type SeckillVoucher struct {
	VoucherID string    `json:"voucher_id" gorm:"column:voucher_id"`
	Stock     int       `json:"stock" gorm:"column:stock"`
	BeginTime time.Time `json:"begin_time" gorm:"column:begin_time"`
	EndTime   time.Time `json:"end_time" gorm:"column:end_time"`
}
