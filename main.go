package main

import (
	"github.com/gin-gonic/gin"
	"go_gin/handlers"
	"io"
	"os"
)

func main() {
	gin.DisableConsoleColor()
	//记录日志
	f, _ := os.Create("./logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router := gin.Default()
	router.POST("/getCode", handlers.GetCode)
	router.POST("/login", handlers.Login)
	router.Run("localhost:8888")
}
