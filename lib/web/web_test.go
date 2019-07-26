package web

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type WebSuite struct {
	gspec.Suite
}

func (suite *WebSuite) TestNew() {
	suite.NotNil(New("localhost:3000"))
	suite.NotNil(New("localhost:3000").Server)
}

func (suite *WebSuite) TestRun() {
	suite.Error(New("localhost:-1").Run())
}

func TestWebSuite(t *testing.T) {
	gspec.Run(t, new(WebSuite))
}
