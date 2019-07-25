package env

// DatabaseEnvs store all env keys for database configuration
type DatabaseEnvs struct{}

// Database envs
func Database() *DatabaseEnvs {
	return &DatabaseEnvs{}
}

// Engine env
func (env *DatabaseEnvs) Engine() string {
	return Env("DB_ENGINE")
}
