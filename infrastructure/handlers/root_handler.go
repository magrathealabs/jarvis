package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RootHandler struct
type RootHandler struct{}

// NewRootHandler constructor
func NewRootHandler() Handler {
	return &RootHandler{}
}

// SetupRoutes on gin
func (handler *RootHandler) SetupRoutes(engine *gin.Engine) {
	engine.GET("/", handler.Index)
}

// Index over / path
func (handler *RootHandler) Index(c *gin.Context) {
	c.String(http.StatusOK, "OK!")
}
