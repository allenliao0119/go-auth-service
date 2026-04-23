package http

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/allenliao0119/go-auth-service/internal/application"
	"github.com/allenliao0119/go-auth-service/internal/http/middleware"
	"github.com/allenliao0119/go-auth-service/internal/logger"
	"github.com/gin-gonic/gin"
)

const (
	gracefulShutdownPeriod             = 5 * time.Second
	defaultHTTPServerReadHeaderTimeout = 2 * time.Second
	defaultHTTPServerIdleTimeout = 60 * time.Second
)

type Server struct {
	server *http.Server
	engine *gin.Engine
	logger logger.Logger
}

func NewServer(app *application.Application) (*Server, error) {
	engine := gin.New()

	engine.Use(gin.Recovery())
	engine.Use(middleware.RequestID())
	engine.Use(middleware.NewInjectLogger(app.Logger).Execute())
	engine.Use(middleware.NewRequestLogger(app.Logger).Execute())

	registerRoutes(engine, app)

	server := &http.Server{
		Addr:              ":" + app.Config.Server.Port,
		Handler:           engine,
		ReadHeaderTimeout: defaultHTTPServerReadHeaderTimeout,
		IdleTimeout:       defaultHTTPServerIdleTimeout,
	}

	return &Server{
		server: server,
		engine: engine,
		logger: app.Logger,
	}, nil
}

func (s *Server) Start() {
	go func() {
		if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error(context.Background(), "server error", "error", err)
		}
	}()
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownPeriod)
	defer cancel()

	s.logger.Info(ctx, "server is shutting down...")

	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error(ctx, "server shutdown error", "error", err)
	}

	s.logger.Info(ctx, "server exiting")
}
