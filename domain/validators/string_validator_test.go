package validators

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type StringValidatorSuite struct {
	gspec.Suite
}

func (suite *StringValidatorSuite) TestStringValidator() {
	suite.NotNil(NewStringValidator(""))
}

func (suite *StringValidatorSuite) TestNil() {
	suite.True(NewStringValidator("").Nil())
	suite.False(NewStringValidator("Miojo").Nil())
}

func (suite *StringValidatorSuite) TestNotNil() {
	suite.False(NewStringValidator("").NotNil())
	suite.True(NewStringValidator("Miojo").NotNil())
}

func TestStringValidatorSuite(t *testing.T) {
	gspec.Run(t, new(StringValidatorSuite))
}
