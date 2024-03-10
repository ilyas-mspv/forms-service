package v1

import (
	"forms/internal/service"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, service *service.FormUseCase) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	h := handler.Group("/v1")
	{
		newFormRoutes(h, service)
	}
}
