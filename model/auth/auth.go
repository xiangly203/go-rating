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

type LoginResp struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}
