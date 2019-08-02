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
	temperature.verifyRecordedAt()
	temperature.verifyTemperatureScale()
	temperature.verifyTemperature()
	temperature.verifyRelativeHumidity()
	temperature.verifyRecordedBy()

	return temperature.Base.Valid()
}

func (temperature *Temperature) verifyRecordedAt() {
	recordedAtValidator := validators.NewTimeValidator(temperature.RecordedAt)

	if recordedAtValidator.InTheFuture() {
		temperature.AppendError(errors.New("recorded_in may not be in the future"))
	}
}

func (temperature *Temperature) verifyTemperatureScale() {
	temperatureScaleValidator := validators.NewTemperatureScaleValidator(temperature.TemperatureScale)

	if temperatureScaleValidator.InvalidEnum() {
		temperature.AppendError(errors.New("temperature scale must be C|F|K"))
	}
}

func (temperature *Temperature) verifyTemperature() {
	temperatureValidator := validators.NewFloat32Validator(temperature.Temperature)

	if !temperatureValidator.InBetween(-1000, 1000) {
		temperature.AppendError(errors.New("unfortunately we haven't met temperatures from other planets yet"))
	}
}

func (temperature *Temperature) verifyRelativeHumidity() {
	relativeHumidityValidator := validators.NewFloat32Validator(temperature.RelativeHumidity)

	if !relativeHumidityValidator.InBetween(0, 1) {
		temperature.AppendError(errors.New("percentage must be between 0 and 1"))
	}
}

func (temperature *Temperature) verifyRecordedBy() {
	recordedByValidator := validators.NewStringValidator(temperature.RecordedBy)

	if recordedByValidator.Nil() {
		temperature.AppendError(errors.New("recorded_by needs to be present"))
	}
}
