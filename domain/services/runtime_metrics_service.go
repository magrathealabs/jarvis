package services

import (
	"time"

	"github.com/magrathealabs/jarvis/domain/models"

	"github.com/magrathealabs/jarvis/domain/repositories"
)

// RuntimeMetricsService handles runtime metrics
type RuntimeMetricsService struct {
	MetricRepository repositories.MetricRepository
}

// NewRuntimeMetricsService constructor
func NewRuntimeMetricsService(metricRepository repositories.MetricRepository) *RuntimeMetricsService {
	return &RuntimeMetricsService{MetricRepository: metricRepository}
}

// StoresEndpointRuntime used in any middleware
func (service *RuntimeMetricsService) StoresEndpointRuntime(url string, startAt time.Time) error {
	ms := service.elapsedTimeInMS(startAt)
	metric := models.NewEndpointRuntime(ms, url)
	return service.MetricRepository.InsertMetric(metric.MetricTag(), metric.MetricValue())
}

func (service *RuntimeMetricsService) elapsedTimeInMS(startAt time.Time) int64 {
	return int64(service.elapsedTime(startAt) / time.Millisecond)
}

func (service *RuntimeMetricsService) elapsedTime(startAt time.Time) time.Duration {
	return time.Since(startAt)
}
