package database

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type MySQLConnectorSuite struct {
	gspec.Suite
}

func (suite *MySQLConnectorSuite) TestNewMySQLConnector() {
	suite.NotNil(NewMySQLConnector())
}

func (suite *MySQLConnectorSuite) TestService() {
	connector := NewMySQLConnector()

	suite.NotEmpty(connector.Service())
	suite.Equal("mysql", connector.Service())
}

func (suite *MySQLConnectorSuite) TestArgs() {
	connector := NewMySQLConnector()

	suite.NotEmpty(connector.Args())
}

func TestMySQLConnectorSuite(t *testing.T) {
	gspec.Run(t, new(MySQLConnectorSuite))
}
