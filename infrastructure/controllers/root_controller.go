package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RootController server get at RootPath
type RootController struct {
	AppController *AppController
}

// NewRootController costructor
func NewRootController(other *AppController) *RootController {
	controller := &RootController{
		AppController: other,
	}

	controller.Routes()

	return controller
}

// Routes setup all handlers
func (controller *RootController) Routes() {
}

// Index GET routes.Root
func (controller *RootController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, "hello jarvis")
}
