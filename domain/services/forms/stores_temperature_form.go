package forms

import (
	"time"

	"github.com/magrathealabs/jarvis/domain/enums"

	"github.com/magrathealabs/jarvis/domain/models"
)

// StoresTemperatureForm bind temperature model infos to store in metric repostory
type StoresTemperatureForm struct {
	RecordedAt       time.Time              `json:"recorded_at"`
	RecordedBy       string                 `json:"recorded_by"`
	Temperature      float32                `json:"temperature"`
	TemperatureScale enums.TemperatureScale `json:"temperature_scale"`
}

// NewStoresTemperature constructor
func NewStoresTemperature() *StoresTemperatureForm {
	return &StoresTemperatureForm{}
}

// BuildTemperature based in this form
func (form *StoresTemperatureForm) BuildTemperature() *models.Temperature {
	temperature := models.NewTemperature()

	temperature.RecordedAt = form.RecordedAt
	temperature.RecordedBy = form.RecordedBy
	temperature.Temperature = form.Temperature
	temperature.TemperatureScale = form.TemperatureScale

	return temperature
}
