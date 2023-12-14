package mw

import (
	"github.com/gin-gonic/gin"
	"go_gin/config"
	model "go_gin/model/auth"
	"go_gin/model/base"
	"go_gin/usecase"
	"net/http"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if len(token) == 0 {
			ctx.Abort()
			ctx.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithToken, "无登录权限"))
			return
		}
		cliam, err := usecase.ParseToken(token)
		if err != nil {
			ctx.Abort()
			ctx.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithToken, err.Error()))
			return
		}
		if cliam.Subject == config.AccessToken {
			newAccessToken, _ := usecase.RefreshToken(token)
			ctx.Abort()
			ctx.IndentedJSON(http.StatusOK, base.RespSuc(model.LoginResp{AccessToken: newAccessToken, RefreshToken: ""}))
		}
		ctx.Next()
	}
}
