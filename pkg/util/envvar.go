package util

import "os"

func ReadEnvVar(envVariable, defaultValue string) string {
	data, found := os.LookupEnv(envVariable)
	if !found {
		return defaultValue
	}
	return data
}
