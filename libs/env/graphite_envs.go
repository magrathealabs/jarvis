package env

import (
	"strconv"

	"github.com/krakenlab/ternary"
)

// GraphiteEnvs store all env keys for Graphite configuration
type GraphiteEnvs struct{}

// Graphite envs
func Graphite() *GraphiteEnvs {
	return &GraphiteEnvs{}
}

// Host env
func (env *GraphiteEnvs) Host() string {
	return Env("GRAPHITE_HOST", "localhost")
}

// Port env
func (env *GraphiteEnvs) Port() string {
	return Env("GRAPHITE_PORT", "2003")
}

// PortAsInt get Port env and convert to int
func (env *GraphiteEnvs) PortAsInt() int {
	port := env.Port()

	portAsInt, err := strconv.Atoi(port)

	return ternary.Int(err == nil, portAsInt, 2003)
}
