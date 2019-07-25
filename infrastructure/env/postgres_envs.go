package env

// PostgresEnvs store all env keys for postgres configuration
type PostgresEnvs struct{}

// Postgres envs
func Postgres() *PostgresEnvs {
	return &PostgresEnvs{}
}

// Host env
func (env *PostgresEnvs) Host() string {
	return Env("POSTGRES_HOST", "localhost")
}

// Port env
func (env *PostgresEnvs) Port() string {
	return Env("POSTGRES_PORT", "5432")
}

// User env
func (env *PostgresEnvs) User() string {
	return Env("POSTGRES_USER", "postgres")
}

// Dbname env
func (env *PostgresEnvs) Dbname() string {
	return Env("POSTGRES_DBNAME", "test")
}

// Pass env
func (env *PostgresEnvs) Pass() string {
	return Env("POSTGRES_PASS", "")
}
