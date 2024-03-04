package model

type ShopListReq struct {
	Name   string `json:"name"`
	TypeID int64  `json:"type_id"`
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
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
	HasMore  bool       `json:"has_more"`
	Total    int64      `json:"total"`
	Count    int64      `json:"count"`
	ShopList []*ShopDTO `json:"shop_list"`
}
