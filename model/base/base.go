package base

type BaseResp struct {
	Message string `json:"message"`
	Status  int64  `json:"status"`
	//Data    map[string]map[string]interface{} `json:"data"`
}
