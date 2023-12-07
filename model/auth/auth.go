package auth

type GetCodeReq struct {
	Phone string `json:"phone"`
}

type GetCodeResp struct {
	Code string `json:"code"`
}

type LoginReq struct {
	GetCodeReq
	Code string `json:"code"`
}

type Token struct {
	Token  string `json:"token"`
	Expire string `json:"expire"`
}
type LoginResp struct {
	Token
}
