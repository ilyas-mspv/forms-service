package main

import (
	"fmt"
	"forms/config"
	"forms/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Print(err)
	}

	err = app.Run(cfg)
	if err != nil {
		fmt.Print(err)
	}
}
