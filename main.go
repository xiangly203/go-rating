package main

import (
	"github.com/gin-gonic/gin"
	"go_gin/api/handler"
	"go_gin/api/mw"
	"go_gin/api/mw/logger"
	"go_gin/biz/dal"
)

func main() {
	//gin.DisableConsoleColor()
	//////记录日志
	//f, _ := os.Create("./logs/gin.log")
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	dal.MysqlInit()
	//router := gin.Default()
	router := gin.New()
	router.Use(gin.Recovery(), logger.TraceLogger())
	router.POST("/getCode", handler.GetCode)
	router.POST("/login", handler.Login)
	router.Use(mw.JWTAuth())
	router.Run("localhost:8888")
	//test.TokenTest()
}
