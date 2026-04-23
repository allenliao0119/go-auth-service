package http

import (
	"net/http"

	"github.com/allenliao0119/go-auth-service/internal/application"
	"github.com/gin-gonic/gin"
)

func registerRootRoutes(r gin.IRouter, _ *application.Application) {
	r.GET("healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}
