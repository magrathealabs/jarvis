package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/magrathealabs/jarvis/infrastructure/handlers/api/v1"
	"github.com/magrathealabs/jarvis/libs/handler"
)

// Handler struct
type Handler struct{}

// NewHandler constructor
func NewHandler() handler.Handler {
	return &Handler{}
}

// SetupRoutes on gin
func (handler *Handler) SetupRoutes(engine *gin.Engine) {
	v1.NewTemperatureHandler().SetupRoutes(engine)
}
