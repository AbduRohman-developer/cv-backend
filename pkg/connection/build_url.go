package connection

import (
	"fmt"
	"github.com/AbduRohman-developer/cv-backend/config"
)

func NewURL(cfg *config.Config, typeString string) string {
	switch typeString {
	case "postgres":
		return fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB,
		)
	case "migration":
		return fmt.Sprintf(
			"connection://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDB,
		)
	default:
		return fmt.Sprintf("there is no url for this type, %s", typeString)
	}
}
