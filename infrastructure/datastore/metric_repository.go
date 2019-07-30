package datastore

import (
	"fmt"
	"time"

	"github.com/krakenlab/ternary"
	"github.com/magrathealabs/jarvis/domain/models"
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

// InsertTemperature into graphite
func (repository *MetricRepository) InsertTemperature(time *time.Time, temperature *models.Temperature) error {
	return nil
}

// InsertEndPointResponseTime into graphite
func (repository *MetricRepository) InsertEndPointResponseTime(endPointResponseTime *models.EndPointResponseTime) error {
	return repository.conn.SimpleSend(fmt.Sprintf("end_point_response_time.%s", endPointResponseTime.Route), fmt.Sprintf("%f", endPointResponseTime.Value))
}
