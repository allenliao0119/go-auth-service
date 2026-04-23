package application

import (
	"github.com/allenliao0119/go-auth-service/internal/config"
	"github.com/allenliao0119/go-auth-service/internal/logger"
)

type Application struct {
	Config *config.Config
	Logger logger.Logger
	UseCase *UseCase
}

func New(cfg *config.Config) (*Application, error) {
	logger := logger.NewSlogger(cfg.Logger)

	app := &Application{
		Config: cfg,
		Logger: logger,
	}

	app.UseCase = NewUseCase(app)
	return app, nil
}
