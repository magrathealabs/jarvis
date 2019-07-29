package env

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type ServerEnvsTest struct {
	gspec.Suite
}

func (suite *ServerEnvsTest) TestHost() {
	suite.Equal("0.0.0.0", Server().Host())
}

func (suite *ServerEnvsTest) TestPort() {
	suite.Equal("3000", Server().Port())
}

func (suite *ServerEnvsTest) TestAddr() {
	suite.Equal("0.0.0.0:3000", Server().Addr())
}

func TestServerEnvsTest(t *testing.T) {
	gspec.Run(t, new(ServerEnvsTest))
}
