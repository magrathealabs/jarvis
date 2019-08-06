package models

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type EndpointRuntimeSuite struct {
	gspec.Suite
}

func (suite *EndpointRuntimeSuite) TestNewEndpointRuntime() {
	suite.NotNil(NewEndpointRuntime(0, "testing"))
}

func (suite *EndpointRuntimeSuite) TestToJSON() {
	endpointRuntime := NewEndpointRuntime(0, "testing")
	json := endpointRuntime.ToJSON()

	suite.Contains(json, "id")
	suite.Contains(json, "created_at")
	suite.Contains(json, "deleted_at")
	suite.Contains(json, "updated_at")

	suite.Contains(json, "value")
	suite.Contains(json, "route")

	suite.Contains(json, "0")
	suite.Contains(json, "testing")
}

func (suite *EndpointRuntimeSuite) TestMetricValue() {
	endpointRuntime := NewEndpointRuntime(0, "/route/api/testing")
	suite.Equal("0", endpointRuntime.MetricValue())
}

func (suite *EndpointRuntimeSuite) TestMetricTag() {
	endpointRuntime := NewEndpointRuntime(0, "/route/api/testing/")
	suite.Equal("end_point_runtime.route_api_testing", endpointRuntime.MetricTag())

	endpointRuntime = NewEndpointRuntime(0, "/")
	suite.Equal("end_point_runtime.root", endpointRuntime.MetricTag())

}

func TestEndpointRuntimeSuite(t *testing.T) {
	gspec.Run(t, new(EndpointRuntimeSuite))
}
