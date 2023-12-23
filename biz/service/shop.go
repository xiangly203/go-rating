package service

import (
	"errors"
	"go_gin/api/model"
	"go_gin/biz/dal/db"
	"go_gin/utils"
)

func ShopList(shop model.ShopListReq) ([]*model.ShopDTO, error) {
	shopList, err := db.ShopList(shop.ShopName, shop.ShopType)
	if err != nil {
		return nil, errors.New("服务器错误，请重试")
	}
	shops,err := utils.ConvertToDTO(shopList, &model.ShopDTO{})
	if err != nil {
		return nil, err
	}
	return shops, nil
}
