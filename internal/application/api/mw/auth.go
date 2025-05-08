package mw

import (
	"github.com/gin-gonic/gin"
	"go_gin/config"
	"go_gin/internal/application/api/model"
	base2 "go_gin/internal/domain/entity/base"
	"go_gin/internal/domain/service"
	"net/http"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if len(token) == 0 {
			ctx.Abort()
			ctx.IndentedJSON(http.StatusOK, base2.RespErr(base2.RespErrWithToken, "无登录权限"))
			return
		}
		cliam, err := service.ParseToken(token)
		if err != nil {
			ctx.Abort()
			ctx.IndentedJSON(http.StatusOK, base2.RespErr(base2.RespErrWithToken, err.Error()))
			return
		}
		if cliam.Subject == config.AccessToken {
			newAccessToken, _ := service.RefreshToken(token)
			ctx.Abort()
			ctx.IndentedJSON(http.StatusOK, base2.RespSuc(model.LoginResp{AccessToken: newAccessToken, RefreshToken: ""}))
		}
		ctx.Next()
	}
}
