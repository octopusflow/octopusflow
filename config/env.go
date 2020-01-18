package config

import "os"

const AppEnv = "APP_ENV"

func GetEnv() string {
	if os.Getenv(AppEnv) == "" {
		return "dev"
	}
	return os.Getenv(AppEnv)
}
