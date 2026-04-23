package middleware

import (
	"context"

	"github.com/allenliao0119/go-auth-service/internal/logger"
	"github.com/gin-gonic/gin"
)

type InjectLogger struct {
	logger logger.Logger
}

func NewInjectLogger(logger logger.Logger) *InjectLogger {
	return &InjectLogger{
		logger: logger,
	}
}

func (i *InjectLogger) Execute() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), logger.ContextKeyLogger, i.logger.With(c))
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
