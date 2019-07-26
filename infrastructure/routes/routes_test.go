package routes

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type RoutesSuite struct {
	gspec.Suite
}

func (suite *RoutesSuite) TestRootConst() {
	suite.NotEmpty(Root)
	suite.Contains(Routes, Root)
}

func (suite *RoutesSuite) TestAPIV1TemperatureConst() {
	suite.NotEmpty(APIV1Temperature)
	suite.Contains(Routes, APIV1Temperature)
}

func (suite *RoutesSuite) TestRoutesConst() {
	for firstIndex, firstRoute := range Routes {
		for secondIndex, secondRoute := range Routes {
			if firstIndex == secondIndex {
				continue
			}

			suite.NotEqual(firstRoute, secondRoute)
		}
	}
}

func TestRoutesSuite(t *testing.T) {
	gspec.Run(t, new(RoutesSuite))
}
