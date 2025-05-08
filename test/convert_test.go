package test

import (
	"fmt"
	"go_gin/internal/application/api/model"
	entity "go_gin/internal/domain/entity/shop"
	"go_gin/utils"
	"testing"
)

func TestConvert(t *testing.T) {
	shop := &entity.Shop{
		Name:      "test",
		TypeID:    1,
		Images:    "test",
		Area:      "test",
		Address:   "test",
		X:         1,
		Y:         1,
		AvgPrice:  1,
		Sold:      1,
		Comments:  1,
		Score:     1,
		OpenHours: "test",
	}
	shopDTO, err := utils.ConvertToDTO(shop, &model.ShopDTO{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(shopDTO)
}
