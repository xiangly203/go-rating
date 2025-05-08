package entity

type Voucher struct {
	VoucherID   string  `json:"voucher_id"`
	ShopId      string  `json:"shop_id"`
	Title       string  `json:"title"`
	SubTitle    string  `json:"sub_title"`
	Rules       string  `json:"rules"`
	PayValue    float64 `json:"pay_value"`
	ActualValue float64 `json:"actual_value"`
	Type        string  `json:"type"`
	Status      string  `json:"status"`
}
