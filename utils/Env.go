package utils

import "os"

// GetEnv egts the env value or instead puts the default value
func GetEnv(name, defaultValue string) string {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env
	}
	return defaultValue
}
