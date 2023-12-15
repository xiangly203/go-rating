package model

import (
	"context"
	"gorm.io/gorm"
)

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

type OrderItemRepository interface {
	Create(c context.Context, orderItem *OrderItem) error
	Fetch(c context.Context) ([]OrderItem, error)
	Update(c context.Context, orderItem *[]OrderItem) error
}
