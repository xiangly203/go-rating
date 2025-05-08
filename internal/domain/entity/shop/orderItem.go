package entity

type OrderItem struct {
	OrderItemID   string  `json:"order_item_id" column:"order_item_id"`
	OrderInfoID   string  `json:"order_info_id" column:"order_info_id"`
	OrderAmount   int64   `json:"order_amount" column:"order_amount"`
	OriginalPrice float64 `json:"original_price" column:"original_price"`
	RealPrice     float64 `json:"real_price" column:"real_price"`
}