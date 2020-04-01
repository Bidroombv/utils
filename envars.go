package utils

import (
	"os"
	"strconv"
	"strings"
)

// GetEnv looks up for an environment string or set it to default
func GetEnv(key string, defaultValue string) string {
	if value, status := os.LookupEnv(key); status {
		// XXX work around a docker-compose bug:
		// https://github.com/docker/compose/issues/2854
		// Maybe it will be fixed one day:
		// https://github.com/docker/compose/pull/6923
		if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) {
			value = strings.TrimPrefix(value, `"`)
			value = strings.TrimSuffix(value, `"`)
		}

		return value
	}
	return defaultValue
}

// GetEnvInt looks up for an environment int or set it to default
func GetEnvInt(key string, defaultValue int) int {
	if value, status := os.LookupEnv(key); status {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// GetEnvBool looks up for an environment bool or set it to default
func GetEnvBool(key string, defaultValue bool) bool {
	if value, status := os.LookupEnv(key); status {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}
