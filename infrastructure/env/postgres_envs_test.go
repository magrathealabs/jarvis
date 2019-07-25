package env

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type PostgresEnvsTest struct {
	gspec.Suite
}

func (suite *PostgresEnvsTest) TestHost() {
	suite.Equal("localhost", Postgres().Host())
}
func (suite *PostgresEnvsTest) TestPort() {
	suite.Equal("5432", Postgres().Port())
}
func (suite *PostgresEnvsTest) TestUser() {
	suite.Equal("postgres", Postgres().User())
}
func (suite *PostgresEnvsTest) TestDbname() {
	suite.Equal("test", Postgres().Dbname())
}
func (suite *PostgresEnvsTest) TestPass() {
	suite.Equal("", Postgres().Pass())
}

func TestPostgresEnvsTest(t *testing.T) {
	gspec.Run(t, new(PostgresEnvsTest))
}
