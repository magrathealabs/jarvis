package models

import (
	"testing"
	"time"

	"github.com/magrathealabs/jarvis/domain/enums"

	"github.com/krakenlab/gspec"
)

type TemperatureSuite struct {
	gspec.Suite
}

func (suite *TemperatureSuite) TestNewTemperature() {
	suite.NotNil(NewTemperature())
}

func (suite *TemperatureSuite) TestToJSON() {
	temperature := NewTemperature()
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

	temperature.TemperatureScale = enums.KelvinTemperaureScale
	suite.Contains(temperature.ToJSON(), enums.KelvinTemperaureScale)

	temperature.TemperatureScale = enums.FahrenheitTemperaureScale
	suite.Contains(temperature.ToJSON(), enums.FahrenheitTemperaureScale)
}

func (suite *TemperatureSuite) TestTemperatureMetricTag() {
	temperature := NewTemperature()

	temperature.RecordedAt = time.Now().Add(time.Hour)
	temperature.RecordedBy = "beacon"
	temperature.Temperature = 15

	suite.Equal("temperature.C.beacon", temperature.TemperatureMetricTag())
}

func (suite *TemperatureSuite) TestRelativeHumidityMetricTag() {
	temperature := NewTemperature()

	temperature.RecordedAt = time.Now().Add(time.Hour)
	temperature.RecordedBy = "beacon"
	temperature.Temperature = 15

	suite.Equal("relative_humidity.beacon", temperature.RelativeHumidityMetricTag())
}

func (suite *TemperatureSuite) TestTemperatureMetricValue() {
	temperature := NewTemperature()

	temperature.RecordedAt = time.Now().Add(time.Hour)
	temperature.RecordedBy = "beacon"
	temperature.Temperature = 15.05

	suite.Equal("15.050000", temperature.TemperatureMetricValue())
}

func (suite *TemperatureSuite) TestRelativeHumidityMetricValue() {
	temperature := NewTemperature()

	temperature.RecordedAt = time.Now().Add(time.Hour)
	temperature.RecordedBy = "beacon"
	temperature.RelativeHumidity = 0.05

	suite.Equal("0.050000", temperature.RelativeHumidityMetricValue())
}

func (suite *TemperatureSuite) TestValid() {
	temperature := NewTemperature()

	temperature.RecordedAt = time.Now().Add(time.Hour)
	temperature.RecordedBy = ""
	temperature.Temperature = -2000
	temperature.TemperatureScale = enums.TemperatureScale("Miojo scale")
	temperature.RelativeHumidity = 1.2

	suite.False(temperature.Valid())
	suite.Equal(5, len(temperature.Errors))
}

func TestTemperatureSuite(t *testing.T) {
	gspec.Run(t, new(TemperatureSuite))
}
