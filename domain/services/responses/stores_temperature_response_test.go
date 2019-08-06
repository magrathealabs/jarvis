package responses

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type StoresTemperatureResponseSuite struct {
	gspec.Suite
}

func (suite *StoresTemperatureResponseSuite) TestNewStoresTemperatureResponse() {
	suite.NotNil(NewStoresTemperature())
}

func (suite *StoresTemperatureResponseSuite) TestStored() {
	suite.False(NewStoresTemperature().Stored)
}

func (suite *StoresTemperatureResponseSuite) TestTemperature() {
	suite.Nil(NewStoresTemperature().Temperature)
}

func TestStoresTemperatureResponseSuite(t *testing.T) {
	gspec.Run(t, new(StoresTemperatureResponseSuite))
}
