package entity

import (
	"context"
)

type OrderInfo struct {
	OrderInfoID string  `json:"order_info_id" column:"order_info_id"`
	UserID      string  `json:"user_id" column:"user_id"`
	ShopId      string  `json:"shop_id" column:"shop_id"`
	ReceiverId  string  `json:"receiver_id" column:"receiver_id"`
	TotalAmount float64 `json:"total_amount" column:"total_amount"`
	PayStatus   int     `json:"pay_status" column:"pay_status"`
	PayAmount   float64 `json:"pay_amount" column:"pay_amount"`
	OrderStatus int     `json:"order_status" column:"order_status"`
	Description string  `json:"description" column:"description"`
}

func (u *OrderInfo) TableName() string {
	return "order_info"
}

type OrderInfoRepository interface {
	Create(c context.Context, orderInfo *OrderInfo) error
	Fetch(c context.Context) ([]OrderInfo, error)
	Update(c context.Context, orderInfo OrderInfo) error
}
