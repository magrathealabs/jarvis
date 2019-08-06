package datastore

import (
	"time"

	"github.com/krakenlab/ternary"
	"github.com/magrathealabs/jarvis/domain/repositories"
	"github.com/magrathealabs/jarvis/libs/env"
	"github.com/marpaia/graphite-golang"
)

// MetricRepository send metrics to graphite
type MetricRepository struct {
	conn *graphite.Graphite
}

// NewMetricRepository constructor
func NewMetricRepository(host string, port int) repositories.MetricRepository {
	conn, err := graphite.NewGraphite(host, port)

	ternary.Func(err == nil, func() {}, func() { panic(err) })()

	return &MetricRepository{conn: conn}
}

// NewMetricRepositoryFromEnv using env package
func NewMetricRepositoryFromEnv() repositories.MetricRepository {
	return NewMetricRepository(env.Graphite().Host(), env.Graphite().PortAsInt())
}

// InsertMetric into graphite
func (repository *MetricRepository) InsertMetric(tag string, value string) error {
	return repository.InsertMetricAt(tag, value, time.Now())
}

// InsertMetricAt custom time into graphite
func (repository *MetricRepository) InsertMetricAt(tag string, value string, at time.Time) error {
	metric := graphite.NewMetric(tag, value, at.UTC().Unix())
	return repository.conn.SendMetric(metric)
}
