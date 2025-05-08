package db

import (
	"errors"
	"go_gin/internal/application/api/model"
	entity "go_gin/internal/domain/entity/shop"
)

func ShopList(s model.ShopListReq, offset int64, limit int64) ([]*model.ShopDTO, int64, error) {
	res := make([]*model.ShopDTO, 0)
	var total int64
	err := Postgres.Model(&entity.Shop{}).Where(&entity.Shop{Name: s.Name, TypeID: int(s.TypeID)}).Limit(int(limit)).Offset(int(offset)).Count(&total).Find(&res).Error
	if err != nil {
		return nil, 0, errors.New("服务器错误，请重试")
	}
	return res, total, nil
}
