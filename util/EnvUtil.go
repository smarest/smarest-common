package util

import (
	"errors"
	"os"
)

const (
	PARAM_FROM_URL = "frm"
)

// GetEnvDefault get required environment
func GetEnvDefault(name string, defaultValue string) string {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env
	}
	return defaultValue
}

// GetEnvRequired get required environment
func GetEnvRequired(name string) (string, error) {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env, nil
	}
	return "", errors.New("not found env: " + name)
}
