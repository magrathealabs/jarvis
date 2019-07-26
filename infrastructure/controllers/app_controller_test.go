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
}

func TestAppControllerSuite(t *testing.T) {
	gspec.Run(t, new(AppControllerSuite))
}
