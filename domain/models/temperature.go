package models

import (
	"encoding/json"
	"time"

	"github.com/krakenlab/ternary"
	"github.com/magrathealabs/jarvis/domain/enums"
)

// Temperature struct
type Temperature struct {
	Base

	RecordedBy       string                 `json:"recorded_by"`
	RecordedAt       time.Time              `json:"recorded_at"`
	Temperature      float32                `json:"temperature"`
	TemperatureScale enums.TemperatureScale `json:"temperature_scale"`
	RelativeHumidity float32                `json:"relative_humidity"`
}

// NewTemperature constructor
func NewTemperature() *Temperature {
	return &Temperature{Temperature: 0, TemperatureScale: enums.CelsiusTemperaureScale}
}

// ToJSON marshal
func (temperature *Temperature) ToJSON() string {
	data, err := json.Marshal(temperature)
	ternary.Func(err == nil, func() {}, func() { panic(err) })()
	return string(data)
}
