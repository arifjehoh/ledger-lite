package config

import (
	"os"
)

type Config struct {
	PORT string
}

func LoadConfig() Config {
	return Config{
		PORT: ":" + getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
