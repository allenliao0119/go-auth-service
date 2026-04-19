package application

import (
	"github.com/allenliao0119/go-auth-service/internal/config"
	"github.com/allenliao0119/go-auth-service/internal/logger"
)

type Application struct {
	Config *config.Config
	Logger logger.Logger
}

func New(cfg *config.Config) (*Application, error) {
	logger := logger.NewSlogger(cfg.Logger)

	return &Application{
		Config: cfg,
		Logger: logger,
	}, nil
}
