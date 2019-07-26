package controllers

import (
	"net/http"
	"testing"

	"github.com/magrathealabs/jarvis/infrastructure/routes"

	"github.com/krakenlab/gspec"
	"github.com/krakenlab/gspec/websuite"
)

type RootControllerSuite struct {
	websuite.WebSuite

	AppController  *AppController
	RootController *RootController
}

func (suite *RootControllerSuite) SetupTest() {
	suite.AppController = NewAppController()
	suite.AppController.Setup()

	suite.RootController = suite.AppController.RootController
	suite.WebSuite.Server = suite.AppController.Engine.Server
	suite.Suite.SetupTest()
}

func (suite *RootControllerSuite) TestIndex() {
	response := suite.GET(routes.Root, "")
	suite.Equal(http.StatusOK, response.Code)
}

func TestRootControllerSuite(t *testing.T) {
	gspec.Run(t, new(RootControllerSuite))
}
