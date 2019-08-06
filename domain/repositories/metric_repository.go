package repositories

import "time"

// MetricRepository interface
type MetricRepository interface {
	InsertMetric(tag string, value string) error
	InsertMetricAt(tag string, value string, at time.Time) error
}
