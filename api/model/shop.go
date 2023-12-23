package model

type ShopListReq struct {
	ShopName string `json:"shop_name"`
	ShopType string `json:"shop_type"`
}
type ShopDTO struct {
	Name      string  `json:"name" column:"name"`
	TypeID    uint64  `json:"type_id" column:"type_id"`
	Images    string  `json:"images" column:"images"`
	Area      string  `json:"area" column:"area"`
	Address   string  `json:"address" column:"address"`
	X         float64 `json:"x" column:"x"`
	Y         float64 `json:"y" column:"y"`
	AvgPrice  uint64  `json:"avg_price" column:"avg_price"`
	Sold      uint32  `json:"sold" column:"sold"`
	Comments  uint32  `json:"comments" column:"comments"`
	Score     uint8   `json:"score" column:"score"`
	OpenHours string  `json:"open_hours" column:"open_hours"`
}

type ShopListResp struct {
	ShopList []ShopDTO `json:"shop_list"`
}
