package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/krakenlab/ternary"
	"github.com/magrathealabs/jarvis/domain/enums"
	"github.com/magrathealabs/jarvis/domain/validators"
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

// Valid model
func (temperature *Temperature) Valid() bool {
	recordedAtValidator := validators.NewTimeValidator(temperature.RecordedAt)
	temperatureScaleValidator := validators.NewTemperatureScaleValidator(temperature.TemperatureScale)
	temperatureValidator := validators.NewFloat32Validator(temperature.Temperature)
	relativeHumidityValidator := validators.NewFloat32Validator(temperature.RelativeHumidity)
	recordedByValidator := validators.NewStringValidator(temperature.RecordedBy)

	if recordedAtValidator.InTheFuture() {
		temperature.AppendError(errors.New("recorded_in may not be in the future"))
	}

	if temperatureScaleValidator.InvalidEnum() {
		temperature.AppendError(errors.New("temperature scale must be C|F|K"))
	}

	if !temperatureValidator.InBetween(-1000, 1000) {
		temperature.AppendError(errors.New("unfortunately we haven't met temperatures from other planets yet"))
	}

	if !relativeHumidityValidator.InBetween(0, 1) {
		temperature.AppendError(errors.New("percentage must be between 0 and 1"))
	}

	if recordedByValidator.Nil() {
		temperature.AppendError(errors.New("recorded_by needs to be present"))
	}

	return temperature.Base.Valid()
}
