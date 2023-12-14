package model

import "gorm.io/gorm"

type OrderInfo struct {
	gorm.Model
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

type OrderItem struct {
	gorm.Model
	OrderItemID string  `json:"order_item_id" column:"order_item_id"`
	OrderInfoID string  `json:"order_info_id" column:"order_info_id"`
	OrderAmount int64   `json:"order_amount" column:"order_amount"`
	OrginPrice  float64 `json:"orgin_price" column:"orgin_price"`
	RealPrice   float64 `json:"real_price" column:"real_price"`
}

func (u *OrderItem) TableName() string {
	return "order_item"
}
