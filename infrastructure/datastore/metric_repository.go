package datastore

import (
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
	return repository.conn.SimpleSend(tag, value)
}
