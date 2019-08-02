package services

import (
	"github.com/magrathealabs/jarvis/domain/repositories"
	"github.com/magrathealabs/jarvis/domain/services/forms"
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
func (service *TemperatureService) StoresTemperature(form *forms.StoresTemperatureForm) {
	_ = form.BuildTemperature()
}
