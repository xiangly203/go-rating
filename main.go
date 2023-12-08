package main

import "go_gin/test"

func main() {
	//gin.DisableConsoleColor()
	////记录日志
	//f, _ := os.Create("./logs/gin.log")
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	//router := gin.Default()
	//router.POST("/getCode", handlers.GetCode)
	//router.POST("/login", handlers.Login)
	//router.Run("localhost:8888")

	test.TokenTest()
}
