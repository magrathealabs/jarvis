package middlewares

import (
	"testing"

	"github.com/magrathealabs/jarvis/infrastructure/datastore"

	"github.com/krakenlab/gspec"
)

type EndpointRuntimeSuite struct {
	gspec.Suite
}

func (suite *EndpointRuntimeSuite) TestEndpointRuntime() {
	repository := datastore.NewMetricRepositoryFromEnv()
	suite.NotNil(EndpointRuntime(repository))
}

func TestEndpointRuntimeSuite(t *testing.T) {
	gspec.Run(t, new(EndpointRuntimeSuite))
}
