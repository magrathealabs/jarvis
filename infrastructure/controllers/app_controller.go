package controllers

// AppController is a generic controller to setup routes
type AppController struct {
	RootController *RootController
}

// NewAppController costructor
func NewAppController() *AppController {
	app := &AppController{}
	app.Routes()

	return app
}

// Routes setup all handlers
func (controller *AppController) Routes() {
	controller.RootController = NewRootController(controller)
}
