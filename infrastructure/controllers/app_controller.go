package controllers

import "github.com/magrathealabs/jarvis/infrastructure/web"

// AppController is a generic controller to setup routes
type AppController struct {
	Engine *web.Engine

	RootController *RootController
}

// NewAppController costructor
func NewAppController() *AppController {
	engine := web.New()
	engine.Setup()

	return &AppController{Engine: engine}
}

// Setup instance all sub controllers
func (controller *AppController) Setup() {
	controller.RootController = NewRootController(controller)

	defer controller.RootController.Setup()
}
