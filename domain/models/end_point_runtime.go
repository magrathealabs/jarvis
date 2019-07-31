package models

import (
	"encoding/json"
	"fmt"
	"strings"

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
func (EndpointRuntime *EndpointRuntime) ToJSON() string {
	data, err := json.Marshal(EndpointRuntime)
	ternary.Func(err == nil, func() {}, func() { panic(err) })()
	return string(data)
}

// MetricValue transformation
func (EndpointRuntime *EndpointRuntime) MetricValue() string {
	return fmt.Sprintf("%d", EndpointRuntime.Value)
}

// MetricTag route transformation
func (EndpointRuntime *EndpointRuntime) MetricTag() string {
	metricRoute := strings.Replace(EndpointRuntime.Route, "/", "_", -1)
	metricRoute = strings.TrimPrefix(metricRoute, "_")
	metricRoute = strings.TrimSuffix(metricRoute, "_")
	metricRoute = ternary.String(metricRoute == "", "root", metricRoute)

	return fmt.Sprintf("end_point_runtime.%s", metricRoute)
}
