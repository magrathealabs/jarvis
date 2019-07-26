package controllers

import (
	"net/http"
	"testing"

	"github.com/krakenlab/gspec"
	"github.com/krakenlab/gspec/websuite"
	"github.com/magrathealabs/jarvis/infrastructure/routes"
)

type AppControllerSuite struct {
	websuite.WebSuite
}

func (suite *AppControllerSuite) SetupTest() {
	suite.Server = NewAppController().Server

	suite.WebSuite.SetupTest()
}

func (suite *AppControllerSuite) TestNewAppController() {
	appController := NewAppController()

	suite.NotNil(appController)
}

func (suite *AppControllerSuite) TestIndex() {
	rr := suite.GET(routes.Root, "")
	suite.Equal(http.StatusOK, rr.Code)
}

func TestAppControllerSuite(t *testing.T) {
	gspec.Run(t, new(AppControllerSuite))
}
