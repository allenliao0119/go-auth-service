package middleware

import (
	"context"

	"github.com/allenliao0119/go-auth-service/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const HeaderKeyRequestID string = "X-Request-ID"

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(HeaderKeyRequestID)
		if requestID == "" {
			requestID = uuid.New().String()
		}
		ctx := context.WithValue(c.Request.Context(), logger.ContextKeyRequestID, requestID)
		c.Request = c.Request.WithContext(ctx)
		c.Header(HeaderKeyRequestID, requestID)
		c.Next()
	}
}