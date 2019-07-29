package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/magrathealabs/jarvis/infrastructure/routes"
)

// RootHandler struct
type RootHandler struct{}

// NewRootHandler constructor
func NewRootHandler() Handler {
	return &RootHandler{}
}

// SetupRoutes on gin
func (handler *RootHandler) SetupRoutes(engine *gin.Engine) {
	engine.GET(routes.Root, handler.Index)
}

// Index over / path
func (handler *RootHandler) Index(c *gin.Context) {
	c.String(http.StatusOK, "OK!")
}
