package responses

import (
	"github.com/magrathealabs/jarvis/domain/models"
)

// StoresTemperatureResponse output
type StoresTemperatureResponse struct {
	Temperature *models.Temperature `json:"temperature"`
	Stored      bool
}

// NewStoresTemperature Response
func NewStoresTemperature() *StoresTemperatureResponse {
	return &StoresTemperatureResponse{}
}
