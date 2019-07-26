package web

import (
	"os"
	"testing"

	"github.com/krakenlab/gspec"
)

type WebSuite struct {
	gspec.Suite
}

func (suite *WebSuite) TestNew() {
	suite.NotNil(New())
	suite.NotNil(New().Server)
}

func (suite *WebSuite) TestRun() {
	os.Setenv("WEB_PORT", "-1")
	os.Setenv("WEB_INTERFACE", "localhost")

	suite.Error(New().Run())
}

func (suite *WebSuite) TestPrivateGinInterfaceFormat() {
	os.Setenv("WEB_PORT", "3030")
	os.Setenv("WEB_INTERFACE", "localhost")

	suite.Equal("localhost:3030", New().ginInterfaceFormat())
}

func TestWebSuite(t *testing.T) {
	gspec.Run(t, new(WebSuite))
}
