package entity

type Resp struct {
	Message string `json:"message"`
	Status  int64  `json:"status"`
	Data    any    `json:"data"`
}

func RespSuc(data any) Resp {
	return Resp{Status: SUCCESS, Message: GetMsg(SUCCESS), Data: data}
}

func RespErr(status int64, message string) Resp {
	return Resp{Status: status, Message: message, Data: ""}
}
