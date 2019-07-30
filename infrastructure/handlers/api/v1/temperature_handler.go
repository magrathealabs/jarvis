package v1

import (
	"net/http"

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
	engine.GET(routes.APIV1Temperature, handler.Index)
}

// Index over /api/v1/temperature path
func (handler *TemperatureHandler) Index(c *gin.Context) {
	c.String(http.StatusOK, "OK!")
}
