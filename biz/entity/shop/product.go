package shop

import (
	"context"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName    string  `json:"product_name" column:"product_name"`
	ProductID      string  `json:"product_id" column:"product_id"`
	ShopId         string  `json:"shop_id" column:"shop_id"`
	StockCount     int64   `json:"stock_count" column:"stock_count"`
	SaleStatus     bool    `json:"sale_status" column:"sale_status"`
	OrginPrice     float64 `json:"orgin_price" column:"orgin_price"`
	DiscountPrice  float64 `json:"discount_price" column:"discount_price"`
	SaleStockCount int64   `json:"sale_stock_count" column:"sale_stock_count"`
	ImageUrl       string  `json:"image_url" column:"image_url"`
	Description    string  `json:"description" column:"description"`
}

func (u *Product) TableName() string {
	return "product"
}

type ProductRepository interface {
	Create(c context.Context, product *Product) error
	Fetch(c context.Context) ([]Product, error)
	Update(c context.Context, products *[]Product) error
}
