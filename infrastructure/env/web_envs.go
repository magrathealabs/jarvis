package env

// WebEnvs store all env keys for web configuration
type WebEnvs struct{}

// Web envs
func Web() *WebEnvs {
	return &WebEnvs{}
}

// Port env
func (env *WebEnvs) Port() string {
	return Env("WEB_PORT", "3030")
}

// Interface env
func (env *WebEnvs) Interface() string {
	return Env("WEB_INTERFACE", "0.0.0.0")
}
