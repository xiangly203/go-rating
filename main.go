package main

import (
	"go_gin/biz/dal/cache"
	"go_gin/biz/dal/db"
	handler2 "go_gin/internal/application/api/handler"
	"go_gin/internal/application/api/mw"
	logger2 "go_gin/internal/application/api/mw/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	db.PostgresInit()
	cache.CacheInit()
	router := gin.New()
	logger2.LogInit()

	router.Use(logger2.Ginzap(logger2.Logger, time.RFC3339, true))

	router.Use(logger2.RecoveryWithZap(logger2.Logger, true))
	userApi := router.Group("/user")
	{
		userApi.GET("/code", handler2.GetCode)
		userApi.POST("/login", handler2.Login)
	}
	router.Use(mw.JWTAuth())
	shopApi := router.Group("/shop")
	{
		shopApi.GET("/list", handler2.ShopList)
		shopApi.GET("/detail", handler2.ShopDetail)
		shopApi.POST("/add", handler2.ShopAdd)
		shopApi.POST("/update", handler2.ShopUpdate)

	}
	err := router.Run(":8888")
	if err != nil {
		return
	}
}
