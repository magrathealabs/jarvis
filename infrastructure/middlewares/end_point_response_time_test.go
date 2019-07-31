package middlewares

import (
	"testing"

	"github.com/magrathealabs/jarvis/infrastructure/datastore"

	"github.com/krakenlab/gspec"
)

type EndPointResponseTimeSuite struct {
	gspec.Suite
}

func (suite *EndPointResponseTimeSuite) TestEndPointResponseTime() {
	repository := datastore.NewMetricRepositoryFromEnv()
	suite.NotNil(EndPointResponseTime(repository))
}

func TestEndPointResponseTimeSuite(t *testing.T) {
	gspec.Run(t, new(EndPointResponseTimeSuite))
}
