package config

import (
	_ "github.com/joho/godotenv/autoload"
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

	AccessSigningKey  string
	RefreshSigningKey string

	CasbinModelConfigPath string
	CasbinModelPath       string

	Environment string

	AccessKeyExpireDays  int
	RefreshKeyExpireDays int
}

// Get gives instance of Config struct with configuration values
func Get() *Config {
	return &Config{
		Host:                  cast.ToString(getOrReturnDefault("HOST", "localhost")),
		Port:                  cast.ToInt(getOrReturnDefault("PORT", 80)),
		PostgresHost:          cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost")),
		PostgresPort:          cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432)),
		PostgresUser:          cast.ToString(getOrReturnDefault("POSTGRES_USER", "akbarshoh")),
		PostgresDB:            cast.ToString(getOrReturnDefault("POSTGRES_DB", "cv_info")),
		PostgresPassword:      cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "")),
		AccessSigningKey:      cast.ToString(getOrReturnDefault("ACCESS_SIGNING_KEY", "")),
		CasbinModelConfigPath: cast.ToString(getOrReturnDefault("CASBIN_MODEL_PATH", "./config/rbac_model.conf")),
		CasbinModelPath:       cast.ToString(getOrReturnDefault("CASBIN_POLICY_PATH", "./config/models.csv")),
		Environment:           cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop")),
		AccessKeyExpireDays:   cast.ToInt(getOrReturnDefault("ACCESS_KEY_EXPIRE_DAYS", 7)),
		RefreshKeyExpireDays:  cast.ToInt(getOrReturnDefault("REFRESH_EXPIRE_DAYS", 3)),
		RefreshSigningKey:     cast.ToString(getOrReturnDefault("REFRESH_SIGNING_KEY", "")),
	}
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
