package datastore

import (
	"testing"

	"github.com/magrathealabs/jarvis/domain/models"

	"github.com/krakenlab/gspec"
)

type MetricRepositorySuite struct {
	gspec.Suite
}

func (suite *MetricRepositorySuite) TestNewMetricRepository() {
	suite.NotNil(NewMetricRepository("localhost", 2003))
}

func (suite *MetricRepositorySuite) TestNewMetricRepositoryFromEnv() {
	suite.NotNil(NewMetricRepositoryFromEnv())
}

func (suite *MetricRepositorySuite) TestInsertEndPointResponseTime() {
	metric := models.NewEndPointResponseTime(0.05, "/testing/graphite/endpoints")
	err := NewMetricRepositoryFromEnv().InsertEndPointResponseTime(metric)

	suite.NoError(err)
}

func TestMetricRepositorySuite(t *testing.T) {
	gspec.Run(t, new(MetricRepositorySuite))
}
