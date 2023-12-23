package db

import (
	"go_gin/biz/dal"
	entity "go_gin/biz/entity/shop"
)

func ShopList(shopName string, shopType string) ([]*entity.Shop, error) {
	res := make([]*entity.Shop, 0)
	if err := dal.Mysql.Where(dal.Mysql.Or("shop_name = ?", shopName).
		Or("shop_type = ?", shopType)).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
