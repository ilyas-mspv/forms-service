package v1

import (
	"forms/internal/lib/logger"
	"forms/internal/models"
	"forms/internal/service"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type formRoutes struct {
	s *service.FormUseCase
	l *slog.Logger
}

type errorResponse struct {
	Error string `json:"error"`
}

func newFormRoutes(handler *gin.RouterGroup, log *slog.Logger, service *service.FormUseCase) {

	f := formRoutes{s: service, l: log}

	h := handler.Group("/form")
	{
		h.GET("/all", f.getAll)
		h.POST("/", f.createForm)
	}
}

type formResponse struct {
	Forms []*models.Form `json:"forms"`
}

// getAll returns all forms
// @Summary     Show all forms
// @Description Show all forms without details
// @ID          form
// @Tags  	    form
// @Accept      json
// @Produce     json
// @Success     200 {object} formResponse
// @Failure     500 {object} errorResponse
// @Router      /form/all [get]
func (f *formRoutes) getAll(context *gin.Context) {
	const op = "controller.http.v1.getAll"
	forms, err := f.s.GetAllForms(context.Request.Context())
	if err != nil {
		f.l.Error("failed to get all forms", logger.Err(err), slog.String("operation", op))
		context.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}
	context.JSON(http.StatusOK, formResponse{forms})
}

type formRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Identifier  string `json:"identifier"`
}

// createForm creates new form
// @Summary     Create new form
// @Description Create new form
// @ID          create-form
// @Tags  	    form
// @Accept      json
// @Produce     json
// @Param       request body formRequest true "request body"
// @Success     200 {integer} integer
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /form [post]
func (f *formRoutes) createForm(context *gin.Context) {
	const op = "controller.http.v1.createForm"
	var form formRequest
	if err := context.ShouldBindJSON(&form); err != nil {
		f.l.Error("failed to bind json", logger.Err(err), slog.String("operation", op))
		context.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}
	//todo validate user input
	id, err := f.s.CreateForm(context.Request.Context(), &models.Form{
		Name:        form.Name,
		Description: form.Description,
		Identifier:  form.Identifier,
	})
	if err != nil {
		f.l.Error("failed to create form", logger.Err(err), slog.String("operation", op))
		context.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}
	context.JSON(http.StatusOK, id)
}
