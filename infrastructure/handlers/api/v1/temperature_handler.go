package v1

import (
	"net/http"

	"github.com/magrathealabs/jarvis/domain/services"

	"github.com/magrathealabs/jarvis/domain/services/forms"

	"github.com/gin-gonic/gin"
	"github.com/magrathealabs/jarvis/domain/repositories"
	"github.com/magrathealabs/jarvis/infrastructure/routes"
	"github.com/magrathealabs/jarvis/libs/handler"
)

// TemperatureHandler struct
type TemperatureHandler struct {
	MetricRepository repositories.MetricRepository
}

// NewTemperatureHandler constructor
func NewTemperatureHandler(metricRepository repositories.MetricRepository) handler.Handler {
	return &TemperatureHandler{
		MetricRepository: metricRepository,
	}
}

// SetupRoutes on gin
func (handler *TemperatureHandler) SetupRoutes(engine *gin.Engine) {
	engine.POST(routes.APIV1Temperature, handler.Create)
}

// Create over /api/v1/temperature path (Post)
func (handler *TemperatureHandler) Create(c *gin.Context) {
	form := forms.NewStoresTemperature()

	if c.BindJSON(form) == nil {
		service := services.NewTemperatureService(handler.MetricRepository)
		c.JSON(http.StatusOK, service.StoresTemperature(form))
	}
}
