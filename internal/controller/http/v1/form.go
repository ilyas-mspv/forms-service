package v1

import (
	"forms/internal/models"
	"forms/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type formRoutes struct {
	s service.FormUseCase
}

func newFormRoutes(handler *gin.RouterGroup, service *service.FormUseCase) {

	f := formRoutes{s: *service}

	h := handler.Group("/form")
	{
		h.GET("/all", f.getAll)
		h.POST("/create", f.createForm)
	}
}

type formResponse struct {
	Forms []*models.Form `json:"forms"`
}

func (f *formRoutes) getAll(context *gin.Context) {
	forms, err := f.s.GetAllForms(context.Request.Context())
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	context.JSON(http.StatusOK, formResponse{forms})
}

type formRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Identifier  string `json:"identifier"`
}

func (f *formRoutes) createForm(context *gin.Context) {
	var form formRequest
	if err := context.ShouldBindJSON(&form); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	//todo validate user input

	id, err := f.s.CreateForm(context.Request.Context(), &models.Form{
		Name:        form.Name,
		Description: form.Description,
		Identifier:  form.Identifier,
	})
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	context.JSON(http.StatusOK, id)
}
