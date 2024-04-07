package mw

import (
	"github.com/gin-gonic/gin"
	"go_gin/api/model"
	entity "go_gin/biz/entity/base"
	"go_gin/biz/service"
	"go_gin/config"
	"net/http"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if len(token) == 0 {
			ctx.Abort()
			ctx.IndentedJSON(http.StatusOK, entity.RespErr(config.RespErrWithToken, "无登录权限"))
			return
		}
		cliam, err := service.ParseToken(token)
		if err != nil {
			ctx.Abort()
			ctx.IndentedJSON(http.StatusOK, entity.RespErr(config.RespErrWithToken, err.Error()))
			return
		}
		if cliam.Subject == config.AccessToken {
			newAccessToken, _ := service.RefreshToken(token)
			ctx.Abort()
			ctx.IndentedJSON(http.StatusOK, entity.RespSuc(model.LoginResp{AccessToken: newAccessToken, RefreshToken: ""}))
		}
		ctx.Next()
	}
}
