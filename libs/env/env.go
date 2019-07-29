package env

import (
	"os"

	"github.com/krakenlab/ternary"
)

// Env based in key. First OS Values.
func Env(key string, defaults ...string) string {
	value := os.Getenv(key)
	return ternary.String(isBlank(value), firstValidDefault(defaults), value)
}

func isBlank(value string) bool {
	return value == ""
}

func firstValidDefault(defaults []string) string {
	for _, def := range defaults {
		if !isBlank(def) {
			return def
		}
	}

	return ""
}
