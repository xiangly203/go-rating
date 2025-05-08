package service

import (
	"errors"
	"go_gin/biz/dal/db"
	"go_gin/internal/application/api/model"
)

func ShopList(shop model.ShopListReq) ([]*model.ShopDTO, int64, error) {

	shopList, total, err := db.ShopList(shop, shop.Offset*shop.Limit, shop.Limit)
	if err != nil {
		return nil, 0, errors.New("服务器错误，请重试")
	}
	return shopList, total, nil
}
