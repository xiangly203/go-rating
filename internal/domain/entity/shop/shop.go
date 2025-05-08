package entity

type Shop struct {
	Name      string  `json:"name" column:"name"`
	TypeID    int     `json:"type_id" column:"type_id"`
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