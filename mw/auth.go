package mw

import (
	"github.com/gin-gonic/gin"
	"go_gin/config"
	"go_gin/model/base"
	"go_gin/utils"
	"net/http"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			ctx.Abort()
			ctx.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithToken, "无登录权限"))
			return
		}
		_, err := utils.ParseToken(auth)
		if err != nil {
			ctx.Abort()
			ctx.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithToken, err.Error()))
			return
		}
	}
}
