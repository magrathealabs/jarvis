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

// PostgresHost env
func (env *DatabaseEnvs) PostgresHost() string {
	return Env("DB_POSTGRES_HOST", "localhost")
}
