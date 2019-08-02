package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/krakenlab/ternary"
	"github.com/magrathealabs/jarvis/domain/enums"
	"github.com/magrathealabs/jarvis/domain/exceptions"
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

// MetricValue to store in MetricRepository
func (temperature *Temperature) MetricValue() string {
	return fmt.Sprintf("%f", temperature.Temperature)
}

// MetricTag to store in MetricRepository
func (temperature *Temperature) MetricTag() string {
	return fmt.Sprintf("temperature.%s.%s", temperature.TemperatureScale, temperature.RecordedBy)
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
		temperature.AppendError(exceptions.ErrorRecordedInMeyNotBeInTheFuture)
	}
}

func (temperature *Temperature) verifyTemperatureScale() {
	temperatureScaleValidator := validators.NewTemperatureScaleValidator(temperature.TemperatureScale)

	if temperatureScaleValidator.InvalidEnum() {
		temperature.AppendError(exceptions.ErrorTemperatureScaleMstBeCFK)
	}
}

func (temperature *Temperature) verifyTemperature() {
	temperatureValidator := validators.NewFloat32Validator(temperature.Temperature)

	if !temperatureValidator.InBetween(-1000, 1000) {
		temperature.AppendError(exceptions.ErrorUnfortunatelyWeHeventMetTemperaturesFromOtherPlanetsYet)
	}
}

func (temperature *Temperature) verifyRelativeHumidity() {
	relativeHumidityValidator := validators.NewFloat32Validator(temperature.RelativeHumidity)

	if !relativeHumidityValidator.InBetween(0, 1) {
		temperature.AppendError(exceptions.ErrorPercentageMustBeBetween0and1)
	}
}

func (temperature *Temperature) verifyRecordedBy() {
	recordedByValidator := validators.NewStringValidator(temperature.RecordedBy)

	if recordedByValidator.Nil() {
		temperature.AppendError(exceptions.ErrorRecordedByNeedsToBePresent)
	}
}
