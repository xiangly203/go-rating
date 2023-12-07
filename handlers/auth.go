package handlers

import (
	"github.com/gin-gonic/gin"
	"go_gin/model/auth"
	_ "go_gin/model/auth"
	"go_gin/model/base"
	"go_gin/utils"
	config "go_gin/utils/common"
	"go_gin/utils/dal/redis"
	"net/http"
)

func GetCode(ctx *gin.Context) {
	var req auth.GetCodeReq
	if err := ctx.BindJSON(&req); err != nil || !utils.IsMobile(req.Phone) {
		ctx.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithServer, "服务器错误，请重试"))
		return
	}
	code, err := redis.GetRedisVal(req.Phone)
	if err != nil || len(code) == 0 {
		code, err = utils.GenCode(req.Phone)
		ok, _ := redis.SetRedisVal(req.Phone, code, config.CodeTTL)
		if err != nil || ok != "OK" {
			ctx.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithServer, "服务器错误，请重试"))
			return
		}
	}
	ctx.IndentedJSON(http.StatusOK, base.RespSuc(auth.GetCodeResp{Code: code}))
}

func Login(ctx *gin.Context) {
	var req auth.LoginReq
	if err := ctx.BindJSON(&req); err != nil || !utils.IsMobile(req.Phone) {
		ctx.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithServer, "服务器错误，请重试"))
		return
	}
	code, err := redis.GetRedisVal(req.Phone)
	if err != nil || len(code) == 0 || req.Code != code {
		ctx.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithPhoneOrCode, "手机号码或者验证码错误，请重试"))
		return
	}
	ctx.IndentedJSON(http.StatusOK, base.RespSuc(auth.GetCodeResp{Code: code}))
}
