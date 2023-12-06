package auth

import "go_gin/model/base"

type GetCodeReq struct {
	Phone string `json:"phone"`
}

type GetCodeResp struct {
	base.BaseResp
	Data map[string]map[string]string `json:"data"`
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
	base.BaseResp
	Data map[string]Token `json:"data"`
}
