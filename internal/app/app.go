package app

import (
	"fmt"
	"forms/config"
	v1 "forms/internal/controller/http/v1"
	service "forms/internal/service"
	"forms/internal/storage/postgres"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) error {

	pg, err := postgres.New(cfg.DB.ConnectionURL)
	if err != nil {
		return err
	}

	useCase := service.New(pg)

	err = pg.Ping()
	if err != nil {
		return err
	}
	fmt.Print(cfg)
	handler := gin.New()
	v1.NewRouter(handler, useCase)
	err = handler.Run()
	if err != nil {
		return err
	}

	return nil
}
