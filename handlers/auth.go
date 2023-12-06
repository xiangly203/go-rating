package handlers

import (
	"github.com/gin-gonic/gin"
	"go_gin/lib"
	"go_gin/model/auth"
	_ "go_gin/model/auth"
	"go_gin/model/base"
	"net/http"
)

func GetCode(ctx *gin.Context) {
	var req auth.GetCodeReq
	if err := ctx.BindJSON(&req); err != nil {
		resp := base.RespErr(0, "服务器错误，晴重试")
		ctx.IndentedJSON(http.StatusOK, resp)
		return
	}
	code, err := lib.GenCode(req.Phone)
	if err != nil {
		resp := base.RespErr(0, "手机号错误，晴重试")
		ctx.IndentedJSON(http.StatusOK, resp)
		return
	}
	resp := base.RespSuc(auth.GetCodeResp{Code: code})
	ctx.IndentedJSON(http.StatusOK, resp)
}
