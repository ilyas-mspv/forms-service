package app

import (
	"context"
	"errors"
	"forms/config"
	v1 "forms/internal/controller/http/v1"
	"forms/internal/lib/logger"
	"forms/internal/service"
	"forms/internal/storage/postgres"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func Run(cfg *config.Config) error {

	log := logger.SetupLogger(cfg.App.Env)
	log.Info("Starting app", "app", cfg.App.Name, "version", cfg.App.Version, "env", cfg.App.Env)

	pg, err := postgres.New(cfg.DB.ConnectionURL)
	defer pg.Close()
	if err != nil {
		log.Error("Failed to connect to database", logger.Err(err))
		return err
	}

	// todo auth
	// todo metrics
	// todo tracing
	// todo healthcheck
	// todo redis cache
	// todo rate limiter

	useCase := service.New(pg)

	handler := gin.New()
	v1.NewRouter(handler, log, useCase)
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(cfg.HTTP.Port),
		Handler: handler,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Info("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Stopping app", "app", cfg.App.Name, "version", cfg.App.Version, "env", cfg.App.Env)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = srv.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown:", logger.Err(err))
		return err
	}

	return nil
}
