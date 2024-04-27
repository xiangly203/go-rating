package logger

import (
	"errors"
	"github.com/google/uuid"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"regexp"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Fn func(c *gin.Context) []zapcore.Field

// Skipper is a function to skip logs based on provided Context
type Skipper func(c *gin.Context) bool

// ZapLogger is the minimal logger interface compatible with zap.Logger
type ZapLogger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
}

type Config struct {
	TimeFormat      string
	UTC             bool
	SkipPaths       []string
	SkipPathRegexps []*regexp.Regexp
	Context         Fn
	DefaultLevel    zapcore.Level
	Skipper         Skipper
}

func Ginzap(logger ZapLogger, timeFormat string, utc bool) gin.HandlerFunc {
	return GinzapWithConfig(logger, &Config{TimeFormat: timeFormat, UTC: utc, DefaultLevel: zapcore.InfoLevel})
}

func GinzapWithConfig(logger ZapLogger, conf *Config) gin.HandlerFunc {
	skipPaths := make(map[string]bool, len(conf.SkipPaths))
	for _, path := range conf.SkipPaths {
		skipPaths[path] = true
	}

	return func(c *gin.Context) {
		start := time.Now()
		logid := uuid.NewString()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Writer.Header().Set("X-Request-Id", logid)
		c.Next()
		track := true

		if _, ok := skipPaths[path]; ok || (conf.Skipper != nil && conf.Skipper(c)) {
			track = false
		}

		if track && len(conf.SkipPathRegexps) > 0 {
			for _, reg := range conf.SkipPathRegexps {
				if !reg.MatchString(path) {
					continue
				}

				track = false
				break
			}
		}

		if track {
			end := time.Now()
			latency := end.Sub(start)
			if conf.UTC {
				end = end.UTC()
			}

			fields := []zapcore.Field{
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.Duration("latency", latency),
				zap.String("X-Request-Id", logid),
			}
			if conf.TimeFormat != "" {
				fields = append(fields, zap.String("time", end.Format(conf.TimeFormat)))
			}

			if conf.Context != nil {
				fields = append(fields, conf.Context(c)...)
			}

			if len(c.Errors) > 0 {
				// Append error field if this is an erroneous request.
				for _, e := range c.Errors.Errors() {
					logger.Error(e, fields...)
				}
			} else {
				zl, ok := logger.(*zap.Logger)
				if ok {
					zl.Log(conf.DefaultLevel, "", fields...)
				} else {
					logger.Error(path, fields...)
				}
			}
		}
	}
}

func defaultHandleRecovery(c *gin.Context, err interface{}) {
	c.AbortWithStatus(http.StatusInternalServerError)
}

func RecoveryWithZap(logger ZapLogger, stack bool) gin.HandlerFunc {
	return CustomRecoveryWithZap(logger, stack, defaultHandleRecovery)
}

func CustomRecoveryWithZap(logger ZapLogger, stack bool, recovery gin.RecoveryFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne.Err, &se) {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) //nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				recovery(c, err)
			}
		}()
		c.Next()
	}
}
