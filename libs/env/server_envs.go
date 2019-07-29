package env

import "fmt"

// ServerEnvs store all env keys for Server configuration
type ServerEnvs struct{}

// Server envs
func Server() *ServerEnvs {
	return &ServerEnvs{}
}

// Host env
func (env *ServerEnvs) Host() string {
	return Env("SERVER_HOST", "0.0.0.0")
}

// Port env
func (env *ServerEnvs) Port() string {
	return Env("SERVER_PORT", "3000")
}

// Addr env
func (env *ServerEnvs) Addr() string {
	return fmt.Sprintf("%s:%s", env.Host(), env.Port())
}
