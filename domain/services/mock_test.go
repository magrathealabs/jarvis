package services

import "time"

// MetricRepositoryMock struct
type MetricRepositoryMock struct {
	InsertMetricReturn   error
	InsertMetricAtReturn error
}

func NewMetricRepositoryMock() *MetricRepositoryMock {
	return &MetricRepositoryMock{}
}

func (mock *MetricRepositoryMock) InsertMetric(tag string, value string) error {
	return mock.InsertMetricReturn
}

func (mock *MetricRepositoryMock) InsertMetricAt(tag string, value string, at time.Time) error {
	return mock.InsertMetricReturn
}
