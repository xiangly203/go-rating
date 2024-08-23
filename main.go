package main

import (
	"github.com/gin-gonic/gin"
	"go_gin/api/handler"
	"go_gin/api/mw"
	"go_gin/api/mw/logger"
	"go_gin/biz/dal/cache"
	"go_gin/biz/dal/db"
	"time"
)

func main() {
	db.PostgresInit()
	cache.CacheInit()
	router := gin.New()
	logger.LogInit()

	router.Use(logger.Ginzap(logger.Logger, time.RFC3339, true))

	router.Use(logger.RecoveryWithZap(logger.Logger, true))
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
	err := router.Run(":8888")
	if err != nil {
		return
	}
}
