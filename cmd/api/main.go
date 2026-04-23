package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/allenliao0119/go-auth-service/internal/application"
	"github.com/allenliao0119/go-auth-service/internal/config"
	"github.com/allenliao0119/go-auth-service/internal/http"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	app, err := application.New(config)
	if err != nil {
		log.Fatalf("failed to create application: %v", err)
	}

	server, err := http.NewServer(app)
	if err != nil {
		app.Logger.Error(context.Background(), "failed to create server", "error", err)
		os.Exit(1)
	}

	server.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	server.Shutdown()
}
