package models

import (
	"encoding/json"

	"github.com/krakenlab/ternary"
	"github.com/magrathealabs/jarvis/domain/enums"
)

// Temperature struct
type Temperature struct {
	Base

	Value float32                `json:"value"`
	Scale enums.TemperatureScale `json:"scale"`
}

// NewTemperature constructor
func NewTemperature(value float32, scale enums.TemperatureScale) *Temperature {
	return &Temperature{Value: value, Scale: scale}
}

// ToJSON marshal
func (temperature *Temperature) ToJSON() string {
	data, err := json.Marshal(temperature)
	ternary.Func(err == nil, func() {}, func() { panic(err) })()
	return string(data)
}
