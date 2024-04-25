package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go_gin/api/handler"
	"go_gin/api/mw"
	"go_gin/api/mw/logger"
	"go_gin/biz/dal"
	"io"
	"time"
)

func main() {
	dal.PostgresInit()
	dal.CacheInit()
	router := gin.New()
	zapLog, _ := zap.NewProduction()
	router.Use(logger.GinzapWithConfig(zapLog, &logger.Config{
		UTC:        true,
		TimeFormat: time.RFC3339,
		Context: logger.Fn(func(c *gin.Context) []zapcore.Field {
			var fields []zapcore.Field
			if requestID := c.Writer.Header().Get("X-Request-Id"); requestID != "" {
				fields = append(fields, zap.String("request_id", requestID))
			}

			if trace.SpanFromContext(c.Request.Context()).SpanContext().IsValid() {
				fields = append(fields, zap.String("trace_id", trace.SpanFromContext(c.Request.Context()).SpanContext().TraceID().String()))
				fields = append(fields, zap.String("span_id", trace.SpanFromContext(c.Request.Context()).SpanContext().SpanID().String()))
			}

			// log request body
			var body []byte
			var buf bytes.Buffer
			tee := io.TeeReader(c.Request.Body, &buf)
			body, _ = io.ReadAll(tee)
			c.Request.Body = io.NopCloser(&buf)
			fields = append(fields, zap.String("body", string(body)))
			return fields
		}),
	}))

	router.Use(logger.RecoveryWithZap(zapLog, true))
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
