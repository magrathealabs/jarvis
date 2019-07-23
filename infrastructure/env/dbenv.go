package env

// DbEnv store DB_* vars
type DbEnv struct {
	*Env
}

// NewDbEnv constructor
func NewDbEnv() *DbEnv {
	return &DbEnv{Env: New()}
}

// Engine is DB_ENGINE
func (env *DbEnv) Engine() string {
	return env.Read("DB_ENGINE", "sqlite")
}
