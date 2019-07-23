package env

import (
	"os"
	"testing"

	"github.com/krakenlab/gspec"
)

type DbEnvSuite struct {
	gspec.Suite
}

func (suite *DbEnvSuite) TestNewDbEnv() {
	suite.NotNil(NewDbEnv())
}

func (suite *DbEnvSuite) TestEngine() {
	env := NewDbEnv()

	suite.Equal("sqlite", env.Engine())

	os.Setenv("DB_ENGINE", "postgres")
	suite.Equal("postgres", env.Engine())
}

func TestDbEnvSuite(t *testing.T) {
	gspec.Run(t, new(DbEnvSuite))
}
