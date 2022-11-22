package config

import (
	"github.com/spf13/cast"
	"os"
)

type Config struct {
	Host             string
	Port             int
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresDB       string
	PostgresPassword string
}

// Get gives instance of Config struct with configuration values
func Get() *Config {
	return &Config{
		Host:             cast.ToString(getOrReturnDefault("HOST", "localhost")),
		Port:             cast.ToInt(getOrReturnDefault("PORT", 80)),
		PostgresHost:     cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost")),
		PostgresPort:     cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432)),
		PostgresUser:     cast.ToString(getOrReturnDefault("POSTGRES_USER", "akbarshoh")),
		PostgresDB:       cast.ToString(getOrReturnDefault("POSTGRES_DB", "cv_info")),
		PostgresPassword: cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "")),
	}
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
