package mw

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"strconv"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		slog.Debug("Debug message")
		//slog.Info("Info message")
		//slog.Warn("Warning message")
		//slog.Error("Error message")

		c.Next()
		status := c.Writer.Status()
		slog.Debug(strconv.Itoa(status))
	}
}
