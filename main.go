package main

import (
	"github.com/gin-gonic/gin"
	"go_gin/api/mw"
	"go_gin/api/mw/logger"
	"go_gin/bootstrap"
	"go_gin/handlers"
)

func main() {
	//gin.DisableConsoleColor()
	//////记录日志
	//f, _ := os.Create("./logs/gin.log")
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	//env =
	//mysql.Init()
	bootstrap.MysqlInit()
	//router := gin.Default()
	router := gin.New()
	router.Use(gin.Recovery(), logger.TraceLogger())
	router.POST("/getCode", handlers.GetCode)
	router.POST("/login", handlers.Login)
	router.Use(mw.JWTAuth())
	router.Run("localhost:8888")
}
