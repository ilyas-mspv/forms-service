package main

import (
	"forms/config"
	"forms/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	err = app.Run(cfg)
	if err != nil {
		log.Fatalf("run error: %v", err)
	}
}
