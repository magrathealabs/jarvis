package gspec

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TestingSuite mock suite testify package
type TestingSuite suite.TestingSuite

// Suite mocking expect object
type Suite struct {
	suite.Suite
	Expect *assert.Assertions
}

// SetupTest add Expect in suite structure
func (suite *Suite) SetupTest() {
	suite.Expect = assert.New(suite.T())
}

// Run a testing suite
func Run(t *testing.T, testingSuite TestingSuite) {
	suite.Run(t, testingSuite)
}
