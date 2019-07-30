package enums

// TemperatureScale enum
type TemperatureScale string

// TemperaureScales
const (
	CelsiusTemperaureScale    = TemperatureScale("C")
	FahrenheitTemperaureScale = TemperatureScale("F")
	KelvinTemperaureScale     = TemperatureScale("K")
)
