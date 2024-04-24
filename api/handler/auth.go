package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go_gin/api/model"
	"go_gin/biz/dal/db"
	"go_gin/biz/entity/base"
	"go_gin/biz/entity/user"
	"go_gin/biz/service"
	"go_gin/config"
	"go_gin/utils"
	"net/http"
)

func GetCode(c *gin.Context) {
	var phone = c.Query("phone")
	if phone == "" || !utils.IsMobile(phone) {
		c.IndentedJSON(http.StatusOK, base.RespErr(base.RespErrWithServer, "服务器错误，请重试"))
		return
	}
	code, err := db.GetRedisVal(phone)
	if err != nil || len(code) == 0 {
		code, err = utils.GenCode(phone)
		ok, _ := db.SetRedisVal(phone, code, config.CodeTTL)
		if err != nil || ok != "OK" {
			c.IndentedJSON(http.StatusOK, base.RespErr(base.RespErrWithServer, "服务器错误，请重试"))
			return
		}
	}
	c.IndentedJSON(http.StatusOK, base.RespSuc(model.GetCodeResp{Code: code}))
}

func Login(ctx *gin.Context) {
	var req model.LoginReq
	if err := ctx.BindJSON(&req); err != nil || !utils.IsMobile(req.Phone) {
		ctx.IndentedJSON(http.StatusOK, base.RespErr(base.RespErrWithServer, "服务器错误，请重试"))
		return
	}
	code, err := db.GetRedisVal(req.Phone)
	if err != nil || len(code) == 0 || req.Code != code {
		ctx.IndentedJSON(http.StatusOK, base.RespErr(base.RespErrWithPhoneOrCode, "手机号码或者验证码错误，请重试"))
		return
	}
	users, err := db.FindUserByNameOrPhoneNumber("", req.Phone)
	if err != nil {
		ctx.IndentedJSON(http.StatusOK, base.RespErr(base.RespErrWithServer, "服务器错误，请重试"))
	}
	if len(users) == 0 {
		uuidStr := uuid.NewString()
		newUser := &user.User{
			UserName:    utils.RandStr(10),
			PhoneNumber: req.Phone,
			Email:       "",
			IsAdmin:     false,
			UUID:        uuidStr,
		}
		usersToCreate := []*user.User{newUser}
		err = db.CreateUsers(usersToCreate)
		if err != nil {
			ctx.IndentedJSON(http.StatusOK, base.RespErr(base.RespErrWithServer, "服务器错误，请重试"))
			return
		}
		users, _ = db.FindUserByNameOrPhoneNumber("", req.Phone)
	}
	userFirst := users[0]
	userInfo := model.UserInfo{UserName: userFirst.UserName, UserPhone: userFirst.PhoneNumber}
	refreshToken, err := service.GenerateToken(userInfo, "refreshToken")
	if err != nil {
		ctx.IndentedJSON(http.StatusOK, base.RespErr(base.RespErrWithServer, "服务器错误，请重试"))
		return
	}
	accessToken, err := service.GenerateToken(userInfo, "accessToken")
	if err != nil {
		ctx.IndentedJSON(http.StatusOK, base.RespErr(base.RespErrWithServer, "服务器错误，请重试"))
		return
	}
	ctx.IndentedJSON(http.StatusOK, base.RespSuc(model.LoginResp{AccessToken: accessToken, RefreshToken: refreshToken}))
}
