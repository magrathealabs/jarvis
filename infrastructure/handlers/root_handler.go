package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/magrathealabs/jarvis/domain/repositories"
	"github.com/magrathealabs/jarvis/infrastructure/handlers/api"
	"github.com/magrathealabs/jarvis/infrastructure/routes"
	"github.com/magrathealabs/jarvis/libs/handler"
)

// RootHandler struct
type RootHandler struct {
	MetricRepository repositories.MetricRepository
}

// NewRootHandler constructor
func NewRootHandler(metricRepository repositories.MetricRepository) handler.Handler {
	return &RootHandler{
		MetricRepository: metricRepository,
	}
}

// SetupRoutes on gin
func (handler *RootHandler) SetupRoutes(engine *gin.Engine) {
	engine.GET(routes.Root, handler.Index)

	api.NewHandler(handler.MetricRepository).SetupRoutes(engine)
}

// Index over / path
func (handler *RootHandler) Index(c *gin.Context) {
	c.String(http.StatusOK, "OK!")
}
