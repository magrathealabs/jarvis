package models

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/krakenlab/ternary"
)

// EndPointResponseTime struct
type EndPointResponseTime struct {
	Base

	Route string  `json:"route"`
	Value float64 `json:"value"`
}

// NewEndPointResponseTime constructor
func NewEndPointResponseTime(value float64, route string) *EndPointResponseTime {
	return &EndPointResponseTime{Value: value, Route: route}
}

// ToJSON marshal
func (endPointResponseTime *EndPointResponseTime) ToJSON() string {
	data, err := json.Marshal(endPointResponseTime)
	ternary.Func(err == nil, func() {}, func() { panic(err) })()
	return string(data)
}

// MetricValue transformation
func (endPointResponseTime *EndPointResponseTime) MetricValue() string {
	return fmt.Sprintf("%f", endPointResponseTime.Value)
}

// MetricTag route transformation
func (endPointResponseTime *EndPointResponseTime) MetricTag() string {
	metricRoute := strings.Replace(endPointResponseTime.Route, "/", "_", -1)
	metricRoute = strings.TrimPrefix(metricRoute, "_")
	metricRoute = strings.TrimSuffix(metricRoute, "_")
	return fmt.Sprintf("end_point_response_time.%s", metricRoute)
}
