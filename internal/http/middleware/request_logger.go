package middleware

import (
	"time"

	"github.com/allenliao0119/go-auth-service/internal/logger"
	"github.com/gin-gonic/gin"
)

type RequestLogger struct {
	logger logger.Logger
}

func NewRequestLogger(logger logger.Logger) *RequestLogger {
	return &RequestLogger{
		logger: logger,
	}
}

func (l *RequestLogger) Execute() gin.HandlerFunc {
	return func(c *gin.Context) {
		l := l.logger.With(c)

		start := time.Now()
		method := c.Request.Method
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		ip := c.ClientIP()

		l.Info(c, "Incoming request", map[string]any{
			"method":     method,
			"path":       path,
			"query":      query,
			"client_ip":  ip,
			"user_agent": c.Request.UserAgent(),
		})

		c.Next()

		latency := time.Since(start)

		l.Info(c, "Outgoing response", map[string]any{
			"method":    method,
			"path":      path,
			"query":     query,
			"client_ip": ip,
			"latency":   latency,
			"status":    c.Writer.Status(),
			"body_size": c.Writer.Size(),
		})
	}
}
