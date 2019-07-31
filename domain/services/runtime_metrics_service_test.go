package services

import (
	"errors"
	"testing"
	"time"

	"github.com/krakenlab/gspec"
)

type RuntimeMetricsServiceSuite struct {
	gspec.Suite
}

func (suite *RuntimeMetricsServiceSuite) TestNewRuntimeMetricsService() {
	suite.NotNil(NewRuntimeMetricsService(NewMetricRepositoryMock()))
}

func (suite *RuntimeMetricsServiceSuite) TestStoresEndpointRuntime() {
	metricRepository := NewMetricRepositoryMock()
	services := NewRuntimeMetricsService(metricRepository)

	metricRepository.InsertMetricReturn = nil
	err := services.StoresEndpointRuntime("/", time.Now())
	suite.NoError(err)

	metricRepository.InsertMetricReturn = errors.New("error on insert")
	err = services.StoresEndpointRuntime("/", time.Now())
	suite.Error(err)
}

func TestRuntimeMetricsServiceSuite(t *testing.T) {
	gspec.Run(t, new(RuntimeMetricsServiceSuite))
}
