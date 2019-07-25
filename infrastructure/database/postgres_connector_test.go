package database

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type PostgresConnectorSuite struct {
	gspec.Suite
}

func (suite *PostgresConnectorSuite) TestNewPostgresConnector() {
	suite.NotNil(NewPostgresConnector())
}

func (suite *PostgresConnectorSuite) TestVariables() {
	expected := map[string]string{
		"Host":   "localhost",
		"Port":   "5432",
		"User":   "postgres",
		"Dbname": "test",
		"Pass":   "",
	}

	suite.Equal(expected, NewPostgresConnector().Variables())
}

func (suite *PostgresConnectorSuite) TestConnect() {
	conn, err := NewPostgresConnector().Connect()

	suite.NotNil(conn)
	suite.Error(err)
}

func (suite *PostgresConnectorSuite) TestService() {
	connector := NewPostgresConnector()

	suite.NotEmpty(connector.Service())
	suite.Equal("postgres", connector.Service())
}

func (suite *PostgresConnectorSuite) TestArgs() {
	expected := "host=localhost port=5432 user=postgres dbname=test password="

	suite.Equal(
		expected,
		NewPostgresConnector().Args(),
	)
}

func TestPostgresConnectorSuite(t *testing.T) {
	gspec.Run(t, new(PostgresConnectorSuite))
}
