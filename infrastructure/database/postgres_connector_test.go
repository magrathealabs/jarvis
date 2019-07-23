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

func (suite *PostgresConnectorSuite) TestService() {
	connector := NewPostgresConnector()

	suite.NotEmpty(connector.Service())
	suite.Equal("postgres", connector.Service())
}

func (suite *PostgresConnectorSuite) TestArgs() {
	connector := NewPostgresConnector()

	suite.NotEmpty(connector.Args())
}

func TestPostgresConnectorSuite(t *testing.T) {
	gspec.Run(t, new(PostgresConnectorSuite))
}
