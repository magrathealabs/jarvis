package controllers

import (
	"testing"

	"github.com/krakenlab/gspec"
	"github.com/krakenlab/gspec/websuite"
)

type AppControllerSuite struct {
	websuite.WebSuite

	AppController *AppController
}

func (suite *AppControllerSuite) SetupTest() {
	suite.AppController = NewAppController()
	suite.AppController.Setup()

	suite.WebSuite.Server = suite.AppController.Engine.Server
	suite.Suite.SetupTest()
}

func (suite *AppControllerSuite) TestNewAppController() {
	suite.NotEmpty(suite.AppController)
	suite.NotEmpty(suite.AppController.RootController)
	suite.NotEmpty(suite.AppController.RootController.AppController)
}

func TestAppControllerSuite(t *testing.T) {
	gspec.Run(t, new(AppControllerSuite))
}
