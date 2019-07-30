package env

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type GraphiteEnvsSuite struct {
	gspec.Suite
}

func (suite *GraphiteEnvsSuite) TestHost() {
	suite.Equal("localhost", Graphite().Host())
}

func (suite *GraphiteEnvsSuite) TestPort() {
	suite.Equal("2003", Graphite().Port())
}

func (suite *GraphiteEnvsSuite) TestPortAsInt() {
	suite.Equal(2003, Graphite().PortAsInt())

}

func TestGraphiteEnvsSuite(t *testing.T) {
	gspec.Run(t, new(GraphiteEnvsSuite))
}
