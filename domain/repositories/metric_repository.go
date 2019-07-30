package repositories

import (
	"time"

	"github.com/magrathealabs/jarvis/domain/models"
)

// MetricRepository interface
type MetricRepository interface {
	InsertTemperature(time *time.Time, temperature *models.Temperature) error
}
