package http

import (
	"github.com/allenliao0119/go-auth-service/internal/application"
	"github.com/gin-gonic/gin"
)

func registerRoutes(engine *gin.Engine, app *application.Application) {
	registerRootRoutes(engine, app)
	registerAPIRoutes(engine, app)
}
