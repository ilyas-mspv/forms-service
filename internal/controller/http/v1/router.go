package v1

import (
	_ "forms/docs"
	"forms/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log/slog"
)

// NewRouter creates new v1 router
// @title Forms
// @description Forms API
// @version 1.0
// @host localhost:8080
// @BasePath /v1
func NewRouter(handler *gin.Engine, log *slog.Logger, service *service.FormUseCase) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler) // todo make disabling
	handler.GET("/swagger/*any", swaggerHandler)

	h := handler.Group("/v1")
	{
		newFormRoutes(h, log, service)
	}
}
