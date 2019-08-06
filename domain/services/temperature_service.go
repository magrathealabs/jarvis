package services

import (
	"github.com/magrathealabs/jarvis/domain/models"
	"github.com/magrathealabs/jarvis/domain/repositories"
	"github.com/magrathealabs/jarvis/domain/services/forms"
	"github.com/magrathealabs/jarvis/domain/services/responses"
)

// TemperatureService handles temperature metrics
type TemperatureService struct {
	MetricRepository repositories.MetricRepository
}

// NewTemperatureService constructor
func NewTemperatureService(metricRepository repositories.MetricRepository) *TemperatureService {
	return &TemperatureService{MetricRepository: metricRepository}
}

// StoresTemperature in metric repository
func (service *TemperatureService) StoresTemperature(form *forms.StoresTemperatureForm) *responses.StoresTemperatureResponse {
	response := responses.NewStoresTemperature()

	response.Temperature = form.BuildTemperature()
	if response.Temperature.Valid() {
		response.Stored = service.insertTemperature(response.Temperature)
	}

	return response
}

func (service *TemperatureService) insertTemperature(temperature *models.Temperature) bool {
	return service.MetricRepository.InsertMetricAt(
		temperature.MetricTag(),
		temperature.MetricValue(),
		temperature.RecordedAt,
	) == nil
}
