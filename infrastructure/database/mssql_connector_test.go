package database

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type MsSQLConnectorSuite struct {
	gspec.Suite
}

func (suite *MsSQLConnectorSuite) TestNewMsSQLConnector() {
	suite.NotNil(NewMsSQLConnector())
}

func (suite *MsSQLConnectorSuite) TestService() {
	connector := NewMsSQLConnector()

	suite.NotEmpty(connector.Service())
	suite.Equal("mssql", connector.Service())
}

func (suite *MsSQLConnectorSuite) TestArgs() {
	connector := NewMsSQLConnector()

	suite.NotEmpty(connector.Args())
}

func TestMsSQLConnectorSuite(t *testing.T) {
	gspec.Run(t, new(MsSQLConnectorSuite))
}
