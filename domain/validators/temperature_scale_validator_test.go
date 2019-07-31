package validators

import (
	"testing"

	"github.com/magrathealabs/jarvis/domain/enums"

	"github.com/krakenlab/gspec"
)

type TemperatureScaleValidatorSuite struct {
	gspec.Suite
}

func (suite *TemperatureScaleValidatorSuite) TestNewTemperatureScaleValidator() {
	suite.NotNil(NewTemperatureScaleValidator(enums.CelsiusTemperaureScale))
}

func (suite *TemperatureScaleValidatorSuite) TestValidEnum() {
	suite.True(NewTemperatureScaleValidator(enums.CelsiusTemperaureScale).ValidEnum())
	suite.True(NewTemperatureScaleValidator(enums.KelvinTemperaureScale).ValidEnum())
	suite.True(NewTemperatureScaleValidator(enums.FahrenheitTemperaureScale).ValidEnum())

	// Miojo's temperature scale
	suite.False(NewTemperatureScaleValidator(enums.TemperatureScale("M")).ValidEnum())
}

func (suite *TemperatureScaleValidatorSuite) TestInvalidEnum() {
	suite.False(NewTemperatureScaleValidator(enums.CelsiusTemperaureScale).InvalidEnum())
	suite.False(NewTemperatureScaleValidator(enums.KelvinTemperaureScale).InvalidEnum())
	suite.False(NewTemperatureScaleValidator(enums.FahrenheitTemperaureScale).InvalidEnum())

	// Miojo's temperature scale
	suite.True(NewTemperatureScaleValidator(enums.TemperatureScale("M")).InvalidEnum())
}

func TestTemperatureScaleValidatorSuite(t *testing.T) {
	gspec.Run(t, new(TemperatureScaleValidatorSuite))
}
