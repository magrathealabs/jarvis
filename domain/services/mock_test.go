package services

// MetricRepositoryMock struct
type MetricRepositoryMock struct {
	InsertMetricReturn error
}

func NewMetricRepositoryMock() *MetricRepositoryMock {
	return &MetricRepositoryMock{}
}

func (mock *MetricRepositoryMock) InsertMetric(tag string, value string) error {
	return mock.InsertMetricReturn
}
