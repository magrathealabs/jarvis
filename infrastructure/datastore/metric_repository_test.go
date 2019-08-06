package datastore

import (
	"testing"
	"time"

	"github.com/krakenlab/gspec"
)

type MetricRepositorySuite struct {
	gspec.Suite
}

func (suite *MetricRepositorySuite) TestNewMetricRepository() {
	suite.NotNil(NewMetricRepository("localhost", 2003))

	suite.Panics(func() { NewMetricRepository("localhost", -1) })
}

func (suite *MetricRepositorySuite) TestNewMetricRepositoryFromEnv() {
	suite.NotNil(NewMetricRepositoryFromEnv())
}

func (suite *MetricRepositorySuite) TestInsertMetric() {
	err := NewMetricRepositoryFromEnv().InsertMetric("testin", "0")

	suite.NoError(err)
}

func (suite *MetricRepositorySuite) TestInsertMetricAt() {
	err := NewMetricRepositoryFromEnv().InsertMetricAt("testin", "0", time.Now())

	suite.NoError(err)
}

func TestMetricRepositorySuite(t *testing.T) {
	gspec.Run(t, new(MetricRepositorySuite))
}
