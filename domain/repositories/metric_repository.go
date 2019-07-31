package repositories

// MetricRepository interface
type MetricRepository interface {
	InsertMetric(tag string, value string) error
}
