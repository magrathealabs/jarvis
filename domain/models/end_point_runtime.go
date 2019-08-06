package models

import (
	"encoding/json"
	"fmt"

	"github.com/krakenlab/ternary"
)

// EndpointRuntime struct
type EndpointRuntime struct {
	Base

	Route string `json:"route"`
	Value int64  `json:"value"`
}

// NewEndpointRuntime constructor
func NewEndpointRuntime(value int64, route string) *EndpointRuntime {
	return &EndpointRuntime{Value: value, Route: route}
}

// ToJSON marshal
func (endpointRuntime *EndpointRuntime) ToJSON() string {
	data, err := json.Marshal(endpointRuntime)
	ternary.Func(err == nil, func() {}, func() { panic(err) })()
	return string(data)
}

// MetricValue transformation
func (endpointRuntime *EndpointRuntime) MetricValue() string {
	return fmt.Sprintf("%d", endpointRuntime.Value)
}

// MetricTag route transformation
func (endpointRuntime *EndpointRuntime) MetricTag() string {
	return fmt.Sprintf("end_point_runtime.%s", endpointRuntime.clearMetricTag(endpointRuntime.Route))
}
