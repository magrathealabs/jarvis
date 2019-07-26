package env

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type WebEnvsSuite struct {
	gspec.Suite
}

func (suite *WebEnvsSuite) TestInterface() {
	suite.Equal("0.0.0.0", Web().Interface())
}

func (suite *WebEnvsSuite) TestPort() {
	suite.Equal("3030", Web().Port())
}

func TestWebEnvsSuite(t *testing.T) {
	gspec.Run(t, new(WebEnvsSuite))
}
