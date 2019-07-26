package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/magrathealabs/jarvis/infrastructure/routes"
	"github.com/magrathealabs/jarvis/infrastructure/web"
)

// RootController server get at RootPath
type RootController struct {
	Engine        *web.Engine
	AppController *AppController
}

// NewRootController costructor
func NewRootController(other *AppController) *RootController {
	return &RootController{
		AppController: other,
		Engine:        other.Engine,
	}
}

// Setup this controller
func (controller *RootController) Setup() {
	controller.Engine.Server.GET(routes.Root, controller.Index)
}

// Index GET routes.Root
func (controller *RootController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, "hello jarvis")
}
