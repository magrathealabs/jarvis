package database

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type Sqlite3ConnectorSuite struct {
	gspec.Suite
}

func (suite *Sqlite3ConnectorSuite) TestNewSqlite3Connector() {
	suite.NotNil(NewSqlite3Connector())
}

func (suite *Sqlite3ConnectorSuite) TestService() {
	connector := NewSqlite3Connector()

	suite.NotEmpty(connector.Service())
	suite.Equal("sqlite3", connector.Service())
}

func (suite *Sqlite3ConnectorSuite) TestArgs() {
	connector := NewSqlite3Connector()

	suite.NotEmpty(connector.Args())
}

func TestSqlite3ConnectorSuite(t *testing.T) {
	gspec.Run(t, new(Sqlite3ConnectorSuite))
}
