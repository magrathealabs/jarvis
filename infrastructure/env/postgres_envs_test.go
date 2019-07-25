package env

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type PostgresEnvsTest struct {
	gspec.Suite
}

func (suite *PostgresEnvsTest) TestEngine() {
	suite.Equal("localhost", Postgres().Host())
}

func TestPostgresEnvsTest(t *testing.T) {
	gspec.Run(t, new(PostgresEnvsTest))
}
