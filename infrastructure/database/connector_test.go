package database

import (
	"os"
	"testing"

	"github.com/krakenlab/gspec"
	"github.com/magrathealabs/jarvis/infrastructure/env"
)

type ConnectorSuite struct {
	gspec.Suite
}

func (suite *ConnectorSuite) TestNewConnector() {
	defer os.Setenv("DB_ENGINE", env.Database().Engine())

	os.Setenv("DB_ENGINE", "")
	suite.Equal("sqlite3", NewConnector().Service())

	os.Setenv("DB_ENGINE", "mssql")
	suite.Equal("mssql", NewConnector().Service())

	os.Setenv("DB_ENGINE", "mysql")
	suite.Equal("mysql", NewConnector().Service())

	os.Setenv("DB_ENGINE", "postgres")
	suite.Equal("postgres", NewConnector().Service())

	os.Setenv("DB_ENGINE", "sqlite3")
	suite.Equal("sqlite3", NewConnector().Service())
}

func TestConnectorSuite(t *testing.T) {
	gspec.Run(t, new(ConnectorSuite))
}
