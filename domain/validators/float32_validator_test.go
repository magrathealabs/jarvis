package validators

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type Float32ValidatorSuite struct {
	gspec.Suite
}

func (suite *Float32ValidatorSuite) TestNewFloat32Validator() {
	suite.NotNil(NewFloat32Validator(0))
}

func (suite *Float32ValidatorSuite) TestInBetween() {
	suite.True(NewFloat32Validator(0).InBetween(-100, 100))
	suite.False(NewFloat32Validator(-200).InBetween(-100, 100))
	suite.False(NewFloat32Validator(200).InBetween(-100, 100))

}

func TestFloat32ValidatorSuite(t *testing.T) {
	gspec.Run(t, new(Float32ValidatorSuite))
}
