package models

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type EndPointResponseTimeSuite struct {
	gspec.Suite
}

func (suite *EndPointResponseTimeSuite) TestNewEndPointResponseTime() {
	suite.NotNil(NewEndPointResponseTime(0, "testing"))
}

func (suite *EndPointResponseTimeSuite) TestToJSON() {
	endPointresponseTime := NewEndPointResponseTime(0, "testing")
	json := endPointresponseTime.ToJSON()

	suite.Contains(json, "id")
	suite.Contains(json, "created_at")
	suite.Contains(json, "deleted_at")
	suite.Contains(json, "updated_at")

	suite.Contains(json, "value")
	suite.Contains(json, "route")

	suite.Contains(json, "0")
	suite.Contains(json, "testing")
}

func (suite *EndPointResponseTimeSuite) TestMetricValue() {
	endPointresponseTime := NewEndPointResponseTime(0, "/route/api/testing")
	suite.Equal("0.000000", endPointresponseTime.MetricValue())
}

func (suite *EndPointResponseTimeSuite) TestMetricTag() {
	endPointresponseTime := NewEndPointResponseTime(0, "/route/api/testing/")
	suite.Equal("end_point_response_time.route_api_testing", endPointresponseTime.MetricTag())

}

func TestEndPointResponseTimeSuite(t *testing.T) {
	gspec.Run(t, new(EndPointResponseTimeSuite))
}
