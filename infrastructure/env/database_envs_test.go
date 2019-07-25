package env

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type DatabaseEnvsTest struct {
	gspec.Suite
}

func (suite *DatabaseEnvsTest) TestEngine() {
	suite.Equal("", Database().Engine())
}

func TestDatabaseEnvsTest(t *testing.T) {
	gspec.Run(t, new(DatabaseEnvsTest))
}
