package shop

import (
	"gorm.io/gorm"
	"time"
)

type SeckillVoucher struct {
	gorm.Model
	VoucherID string    `json:"voucher_id" gorm:"column:voucher_id"`
	Stock     int       `json:"stock" gorm:"column:stock"`
	BeginTime time.Time `json:"begin_time" gorm:"column:begin_time"`
	EndTime   time.Time `json:"end_time" gorm:"column:end_time"`
}

func (u *SeckillVoucher) TableName() string {
	return "seckill_voucher"
}
