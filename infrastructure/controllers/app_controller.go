package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/magrathealabs/jarvis/infrastructure/routes"
)

// AppController is a generic controller to setup routes
type AppController struct {
	Server *gin.Engine
}

// NewAppController costructor
func NewAppController() *AppController {
	app := &AppController{Server: gin.Default()}
	app.Routes()

	return app
}

// Routes setup all handlers
func (controller *AppController) Routes() {
	controller.Server.GET(routes.Root, controller.Index)
}

// Index GET /
func (controller *AppController) Index(c *gin.Context) {
	c.JSON(
		200,
		"hello jarvis",
	)
}
