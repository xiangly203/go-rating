package handler

import (
	"go_gin/api/model"
	"go_gin/biz/dal/db"
	"go_gin/biz/entity/base"
	"go_gin/biz/service"
	"go_gin/config"
	"go_gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShopList(c *gin.Context) {
	var req model.ShopListReq
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithServer, "服务器错误，请重试"))
		return
	}

	shopList, err := service.ShopList(req)
	if err != nil || len(shopList) == 0 {
		c.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithServer, "服务器错误，请重试"))
		return
	}
	shopDTO, err := make([]model.ShopDTO, len(shopList))

	c.IndentedJSON(http.StatusOK, base.RespSuc(model.ShopListResp{}))
}

func ShopUpdate(c *gin.Context) {
	var req model.GetCodeReq
	if err := c.BindJSON(&req); err != nil || !utils.IsMobile(req.Phone) {
		c.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithServer, "服务器错误，请重试"))
		return
	}
	code, err := db.GetRedisVal(req.Phone)
	if err != nil || len(code) == 0 {
		code, err = utils.GenCode(req.Phone)
		ok, _ := db.SetRedisVal(req.Phone, code, config.CodeTTL)
		if err != nil || ok != "OK" {
			c.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithServer, "服务器错误，请重试"))
			return
		}
	}
	c.IndentedJSON(http.StatusOK, base.RespSuc(model.GetCodeResp{Code: code}))
}

func ShopAdd(c *gin.Context) {
	var req model.GetCodeReq
	if err := c.BindJSON(&req); err != nil || !utils.IsMobile(req.Phone) {
		c.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithServer, "服务器错误，请重试"))
		return
	}
	code, err := db.GetRedisVal(req.Phone)
	if err != nil || len(code) == 0 {
		code, err = utils.GenCode(req.Phone)
		ok, _ := db.SetRedisVal(req.Phone, code, config.CodeTTL)
		if err != nil || ok != "OK" {
			c.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithServer, "服务器错误，请重试"))
			return
		}
	}
	c.IndentedJSON(http.StatusOK, base.RespSuc(model.GetCodeResp{Code: code}))
}
