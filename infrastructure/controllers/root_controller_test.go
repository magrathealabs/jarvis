package controllers

import (
	"testing"

	"github.com/krakenlab/gspec"
	"github.com/krakenlab/gspec/websuite"
)

type RootControllerSuite struct {
	websuite.WebSuite

	AppController  *AppController
	RootController *RootController
}

func (suite *RootControllerSuite) SetupTest() {
}

func TestRootControllerSuite(t *testing.T) {
	gspec.Run(t, new(RootControllerSuite))
}
