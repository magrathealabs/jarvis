package env

import (
	"os"
	"testing"

	"github.com/krakenlab/gspec"
)

type EnvSuite struct {
	gspec.Suite
}

func (suite *EnvSuite) TestNew() {
	suite.NotNil(New())
}

func (suite *EnvSuite) TestRead() {
	env := New()

	suite.NotEmpty(env.Read("GOPATH", ""))
	suite.Equal(os.Getenv("GOPATH"), env.Read("GOPATH", ""))
	suite.Equal("DEFAULT", env.Read("MOCHILEIRO", "DEFAULT"))
}

func TestEnvSuite(t *testing.T) {
	gspec.Run(t, new(EnvSuite))
}
