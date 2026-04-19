package main

import (
	"log"

	"github.com/allenliao0119/go-auth-service/internal/application"
	"github.com/allenliao0119/go-auth-service/internal/config"
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

	log.Println(app)
}
