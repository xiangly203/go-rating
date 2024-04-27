package handler

import (
	"github.com/goccy/go-json"
	"go_gin/api/model"
	"go_gin/biz/entity/base"
	"go_gin/biz/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShopList(c *gin.Context) {
	var req model.ShopListReq
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusOK, base.RespErr(base.RespErrWithServer, "服务器错误，请重试"))
		return
	}
	//resKey := "string(req)"
	//if resJson ,err:=cache.Cache.CacheGet(resKey); err == nil {
	//	c.IndentedJSON(http.StatusOK, base.RespSuc(resJson))
	//}
	shopList, total, err := service.ShopList(req)
	if err != nil || len(shopList) == 0 {
		c.IndentedJSON(http.StatusOK, base.RespErr(base.RespErrWithServer, "服务器错误，请重试"))
		return
	}
	hasMore := false
	count := int64(len(shopList))
	if total > req.Limit*req.Offset+count {
		hasMore = true
	}
	resJson, _ := json.Marshal(model.ShopListResp{ShopList: shopList, HasMore: hasMore, Count: count, Total: total})
	//err = cache.Cache.CacheSet(resKey, resJson, config.DefaultCacheTTL)
	//if err != nil {
	//	return
	//}
	c.IndentedJSON(http.StatusOK, base.RespSuc(resJson))
}

func ShopUpdate(c *gin.Context) {
	var req model.ShopUpdateReq
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusOK, base.RespErr(base.RespErrWithServer, "服务器错误，请重试"))
		return
	}

	c.IndentedJSON(http.StatusOK, base.RespSuc(model.GetCodeResp{Code: "200"}))
}

func ShopAdd(c *gin.Context) {
	var req model.ShopDTO
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusOK, base.RespErr(base.RespErrWithServer, "服务器错误，请重试"))
		return
	}
}

func ShopDetail(c *gin.Context) {
	var req model.ShopDetailReq
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusOK, base.RespErr(base.RespErrWithServer, "服务器错误，请重试"))
		return
	}
}
