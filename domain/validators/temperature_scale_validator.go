package validators

import (
	"github.com/magrathealabs/jarvis/domain/enums"
)

// TemperatureScaleValidator verify time structs
type TemperatureScaleValidator struct {
	Field enums.TemperatureScale
}

// NewTemperatureScaleValidator constructor
func NewTemperatureScaleValidator(field enums.TemperatureScale) *TemperatureScaleValidator {
	return &TemperatureScaleValidator{Field: field}
}

// ValidEnum verify the domain of enums.TemperatureScale
func (validator *TemperatureScaleValidator) ValidEnum() bool {
	return validator.Field == enums.CelsiusTemperaureScale ||
		validator.Field == enums.KelvinTemperaureScale ||
		validator.Field == enums.FahrenheitTemperaureScale
}

// InvalidEnum verify the domain of enums.TemperatureScale
func (validator *TemperatureScaleValidator) InvalidEnum() bool {
	return !validator.ValidEnum()
}
