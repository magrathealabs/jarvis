package api

import (
	"github.com/gin-gonic/gin"
	"github.com/magrathealabs/jarvis/domain/repositories"
	v1 "github.com/magrathealabs/jarvis/infrastructure/handlers/api/v1"
	"github.com/magrathealabs/jarvis/libs/handler"
)

// Handler struct
type Handler struct {
	MetricRepository repositories.MetricRepository
}

// NewHandler constructor
func NewHandler(metricRepository repositories.MetricRepository) handler.Handler {
	return &Handler{
		MetricRepository: metricRepository,
	}
}

// SetupRoutes on gin
func (handler *Handler) SetupRoutes(engine *gin.Engine) {
	v1.NewTemperatureHandler(handler.MetricRepository).SetupRoutes(engine)
}
