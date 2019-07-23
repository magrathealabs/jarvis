package env

import (
	"os"

	"github.com/krakenlab/ternary"
)

// Env base structure
type Env struct{}

// New env object
func New() *Env {
	return &Env{}
}

func (env *Env) Read(key string, def string) string {
	res := os.Getenv(key)
	return ternary.String(res == "", def, res)
}
