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
	router.POST("/getCode", handler.GetCode)
	router.POST("/login", handler.Login)
	router.Use(mw.JWTAuth())
	shop := router.Group("/shop")
	{
		shop.GET("/list", handler.ShopList)
		shop.GET("/detail", handler.ShopDetail)
		shop.POST("/add", handler.ShopAdd)
		shop.POST("/update", handler.ShopUpdate)
	}
	err := router.Run("localhost:8888")
	if err != nil {
		return
	}
}
