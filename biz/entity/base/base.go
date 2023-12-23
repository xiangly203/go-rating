package base

import (
	jsoniter "github.com/json-iterator/go"
	"go_gin/config"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type BaseResp struct {
	Message string      `json:"message"`
	Status  int64       `json:"status"`
	Data    interface{} `json:"data"`
}

func RespSuc(data interface{}) BaseResp {
	return BaseResp{Status: config.RespOK, Message: "success", Data: data}
}

func RespErr(status int64, message string) BaseResp {
	return BaseResp{Status: status, Message: message, Data: ""}
}
