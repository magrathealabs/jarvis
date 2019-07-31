package datastore

import (
	"testing"

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

func (suite *MetricRepositorySuite) TestInsertMetric() {
	err := NewMetricRepositoryFromEnv().InsertMetric("testin", "0")

	suite.NoError(err)
}

func TestMetricRepositorySuite(t *testing.T) {
	gspec.Run(t, new(MetricRepositorySuite))
}
