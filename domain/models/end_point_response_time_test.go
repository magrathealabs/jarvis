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

func TestEndPointResponseTimeSuite(t *testing.T) {
	gspec.Run(t, new(EndPointResponseTimeSuite))
}
