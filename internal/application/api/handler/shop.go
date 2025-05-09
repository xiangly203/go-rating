package handler

import (
	"github.com/goccy/go-json"
	model2 "go_gin/internal/application/api/model"
	base2 "go_gin/internal/domain/entity/base"
	"go_gin/internal/domain/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShopList(c *gin.Context) {
	var req model2.ShopListReq
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusOK, base2.RespErr(base2.RespErrWithServer, "服务器错误，请重试"))
		return
	}
	//resKey := "string(req)"
	//if resJson ,err:=cache.Cache.CacheGet(resKey); err == nil {
	//	c.IndentedJSON(http.StatusOK, base.RespSuc(resJson))
	//}
	shopList, total, err := service.ShopList(req)
	if err != nil || len(shopList) == 0 {
		c.IndentedJSON(http.StatusOK, base2.RespErr(base2.RespErrWithServer, "服务器错误，请重试"))
		return
	}
	hasMore := false
	count := int64(len(shopList))
	if total > req.Limit*req.Offset+count {
		hasMore = true
	}
	resJson, _ := json.Marshal(model2.ShopListResp{ShopList: shopList, HasMore: hasMore, Count: count, Total: total})
	//err = cache.Cache.CacheSet(resKey, resJson, config.DefaultCacheTTL)
	//if err != nil {
	//	return
	//}
	c.IndentedJSON(http.StatusOK, base2.RespSuc(resJson))
}

func ShopUpdate(c *gin.Context) {
	var req model2.ShopUpdateReq
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusOK, base2.RespErr(base2.RespErrWithServer, "服务器错误，请重试"))
		return
	}

	c.IndentedJSON(http.StatusOK, base2.RespSuc(model2.GetCodeResp{Code: "200"}))
}

func ShopAdd(c *gin.Context) {
	var req model2.ShopDTO
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusOK, base2.RespErr(base2.RespErrWithServer, "服务器错误，请重试"))
		return
	}
}

func ShopDetail(c *gin.Context) {
	var req model2.ShopDetailReq
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusOK, base2.RespErr(base2.RespErrWithServer, "服务器错误，请重试"))
		return
	}
}
