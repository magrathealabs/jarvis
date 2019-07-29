package database

import (
	"os"
	"testing"

	"github.com/krakenlab/gspec"
)

type ConnectorSuite struct {
	gspec.Suite
}

func (suite *ConnectorSuite) TestNewConnector() {
	os.Setenv("DB_ENGINE", "postgres")
	suite.Equal("postgres", NewConnector().Service())
}

func TestConnectorSuite(t *testing.T) {
	gspec.Run(t, new(ConnectorSuite))
}
