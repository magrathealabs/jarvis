package models

import (
	"encoding/json"

	"github.com/krakenlab/ternary"
)

// EndPointResponseTime struct
type EndPointResponseTime struct {
	Base

	Route string  `json:"route"`
	Value float32 `json:"value"`
}

// NewEndPointResponseTime constructor
func NewEndPointResponseTime(value float32, route string) *EndPointResponseTime {
	return &EndPointResponseTime{Value: value, Route: route}
}

// ToJSON marshal
func (EndPointResponseTime *EndPointResponseTime) ToJSON() string {
	data, err := json.Marshal(EndPointResponseTime)
	ternary.Func(err == nil, func() {}, func() { panic(err) })()
	return string(data)
}
