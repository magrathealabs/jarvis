package models

import (
	"testing"

	"github.com/magrathealabs/jarvis/domain/enums"

	"github.com/krakenlab/gspec"
)

type TemperatureSuite struct {
	gspec.Suite
}

func (suite *TemperatureSuite) TestNewTemperature() {
	suite.NotNil(NewTemperature(0, enums.CelsiusTemperaureScale))
}

func (suite *TemperatureSuite) TestToJSON() {
	temperature := NewTemperature(0, enums.CelsiusTemperaureScale)
	json := temperature.ToJSON()

	suite.Contains(json, "id")
	suite.Contains(json, "created_at")
	suite.Contains(json, "deleted_at")
	suite.Contains(json, "updated_at")

	suite.Contains(json, "temperature")
	suite.Contains(json, "temperature_scale")
	suite.Contains(json, "relative_humidity")
	suite.Contains(json, "recorded_by")
	suite.Contains(json, "recorded_at")

	suite.Contains(json, "0")
	suite.Contains(json, enums.CelsiusTemperaureScale)

	temperature = NewTemperature(0, enums.KelvinTemperaureScale)
	suite.Contains(temperature.ToJSON(), enums.KelvinTemperaureScale)

	temperature = NewTemperature(0, enums.FahrenheitTemperaureScale)
	suite.Contains(temperature.ToJSON(), enums.FahrenheitTemperaureScale)
}

func TestTemperatureSuite(t *testing.T) {
	gspec.Run(t, new(TemperatureSuite))
}
