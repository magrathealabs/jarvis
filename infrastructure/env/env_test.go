package env

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type EnvFuncSuite struct {
	gspec.Suite
}

func (suite *EnvFuncSuite) TestEnvFunc() {
	gopath := Env("GOPATH", "err")
	otherEnv := Env("NOT GOPATH", "", "ok")
	failEnv := Env("NOT GOPATH", "", "")

	suite.NotEmpty(gopath)
	suite.NotEmpty(otherEnv)
	suite.NotEqual("err", gopath)
	suite.Equal("ok", otherEnv)
	suite.Equal("", failEnv)
}

func TestEnvFuncSuite(t *testing.T) {
	gspec.Run(t, new(EnvFuncSuite))
}
