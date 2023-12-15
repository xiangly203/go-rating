package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go_gin/config"
	model "go_gin/model/auth"
	"go_gin/model/base"
	"go_gin/model/user"
	"go_gin/repository"
	"go_gin/service"
	"go_gin/utils"
	"net/http"
)

func GetCode(ctx *gin.Context) {
	var req model.GetCodeReq
	if err := ctx.BindJSON(&req); err != nil || !utils.IsMobile(req.Phone) {
		ctx.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithServer, "服务器错误，请重试"))
		return
	}
	code, err := repository.GetRedisVal(req.Phone)
	if err != nil || len(code) == 0 {
		code, err = utils.GenCode(req.Phone)
		ok, _ := repository.SetRedisVal(req.Phone, code, config.CodeTTL)
		if err != nil || ok != "OK" {
			ctx.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithServer, "服务器错误，请重试"))
			return
		}
	}
	ctx.IndentedJSON(http.StatusOK, base.RespSuc(model.GetCodeResp{Code: code}))
}

func Login(ctx *gin.Context) {
	var req model.LoginReq
	if err := ctx.BindJSON(&req); err != nil || !utils.IsMobile(req.Phone) {
		ctx.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithServer, "服务器错误，请重试"))
		return
	}
	code, err := repository.GetRedisVal(req.Phone)
	if err != nil || len(code) == 0 || req.Code != code {
		ctx.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithPhoneOrCode, "手机号码或者验证码错误，请重试"))
		return
	}
	users, err := repository.FindUserByNameOrPhoneNumber("", req.Phone)
	if err != nil {
		ctx.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithServer, "服务器错误，请重试"))
	}
	if len(users) == 0 {
		uuidStr := uuid.NewString()
		newUser := &user.User{
			UserName:    utils.RandStr(10),
			PhoneNumber: req.Phone,
			Email:       "",
			IsAdmin:     false,
			IsDelete:    false,
			UUID:        uuidStr,
		}
		usersToCreate := []*user.User{newUser}
		err = repository.CreateUsers(usersToCreate)
		if err != nil {
			ctx.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithServer, "服务器错误，请重试"))
			return
		}
		users, _ = repository.FindUserByNameOrPhoneNumber("", req.Phone)
	}
	userFirst := users[0]
	userInfo := user.UserInfo{UserName: userFirst.UserName, UserPhone: userFirst.PhoneNumber}
	refreshToken, err := service.GenerateToken(userInfo, "refreshToken")
	if err != nil {
		ctx.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithServer, "服务器错误，请重试"))
		return
	}
	accessToken, err := service.GenerateToken(userInfo, "accessToken")
	if err != nil {
		ctx.IndentedJSON(http.StatusOK, base.RespErr(config.RespErrWithServer, "服务器错误，请重试"))
		return
	}
	ctx.IndentedJSON(http.StatusOK, base.RespSuc(model.LoginResp{AccessToken: accessToken, RefreshToken: refreshToken}))
}
