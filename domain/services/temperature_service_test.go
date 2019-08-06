package services

import (
	"testing"
	"time"

	"github.com/magrathealabs/jarvis/domain/enums"
	"github.com/magrathealabs/jarvis/domain/services/forms"

	"github.com/krakenlab/gspec"
)

type TemperatureServiceSuite struct {
	gspec.Suite
}

func (suite *TemperatureServiceSuite) TestNewTemperatureService() {
	suite.NotNil(NewTemperatureService(NewMetricRepositoryMock()))
}

func (suite *TemperatureServiceSuite) TestStoresTemperature() {
	metricRepository := NewMetricRepositoryMock()
	metricRepository.InsertMetricAtReturn = nil

	service := NewTemperatureService(metricRepository)

	form := forms.NewStoresTemperature()
	form.RecordedAt = time.Now()
	form.RecordedBy = "jarvis_beacon"
	form.Temperature = 45
	form.TemperatureScale = enums.CelsiusTemperaureScale

	response := service.StoresTemperature(form)

	suite.Equal(response.Temperature.RecordedAt, form.RecordedAt)
	suite.Equal(response.Temperature.RecordedBy, form.RecordedBy)
	suite.Equal(response.Temperature.Temperature, form.Temperature)
	suite.Equal(response.Temperature.TemperatureScale, form.TemperatureScale)

	suite.True(response.Stored)
}

func TestTemperatureServiceSuite(t *testing.T) {
	gspec.Run(t, new(TemperatureServiceSuite))
}
