package forms

import (
	"testing"
	"time"

	"github.com/magrathealabs/jarvis/domain/enums"

	"github.com/krakenlab/gspec"
)

type StoresTemperatureFormSuite struct {
	gspec.Suite
}

func (suite *StoresTemperatureFormSuite) TestNewStoresTemperatureForm() {
	suite.NotNil(NewStoresTemperature())
}

func (suite *StoresTemperatureFormSuite) TestBuildTemperature() {
	form := NewStoresTemperature()
	form.RecordedBy = "device"
	form.RecordedAt = time.Now()
	form.Temperature = -5
	form.TemperatureScale = enums.CelsiusTemperaureScale

	temperature := form.BuildTemperature()

	suite.Equal(form.RecordedBy, temperature.RecordedBy)
	suite.Equal(form.RecordedAt, temperature.RecordedAt)
	suite.Equal(form.Temperature, temperature.Temperature)
	suite.Equal(form.TemperatureScale, temperature.TemperatureScale)
}

func TestStoresTemperatureFormSuite(t *testing.T) {
	gspec.Run(t, new(StoresTemperatureFormSuite))
}
