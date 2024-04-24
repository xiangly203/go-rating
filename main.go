package main

import (
	"github.com/gin-gonic/gin"
	"go_gin/api/handler"
	"go_gin/api/mw"
	"go_gin/api/mw/logger"
	"go_gin/biz/dal"
)

func main() {
	dal.PostgresInit()
	dal.CacheInit()
	router := gin.New()
	router.Use(gin.Recovery(), logger.TraceLogger())
	userApi := router.Group("/user")
	{
		userApi.GET("/code", handler.GetCode)
		userApi.POST("/login", handler.Login)
	}
	router.Use(mw.JWTAuth())
	shopApi := router.Group("/shop")
	{
		shopApi.GET("/list", handler.ShopList)
		shopApi.GET("/detail", handler.ShopDetail)
		shopApi.POST("/add", handler.ShopAdd)
		shopApi.POST("/update", handler.ShopUpdate)
	}
	err := router.Run("localhost:8888")
	if err != nil {
		return
	}
}
