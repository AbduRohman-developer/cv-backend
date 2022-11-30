package connection

import (
	"errors"
	"fmt"
	"github.com/AbduRohman-developer/cv-backend/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database"          //database is used by migrate
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //connection is used as driver name
	_ "github.com/golang-migrate/migrate/v4/source/file"       //the file package's init function is used by migrate
	"github.com/jmoiron/sqlx"
	"log"
)

func New() *sqlx.DB {
	cfg := config.Get()

	// Connection for connection database with sqlx
	db, err := sqlx.Connect(
		"postges",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB,
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Migration action for automation of database
	if err := migration(cfg, "up"); err != nil {
		log.Fatal(err)
	}
	return db
}

func migration(cfg *config.Config, mType string) error {
	// Generate migration for migrate action
	m, err := migrate.New("file://migrations", fmt.Sprintf(
		"connection://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDB,
	),
	)
	if err != nil {
		return err
	}

	switch mType {
	case "up":
		if err := m.Up(); err != nil && errors.Is(err, migrate.ErrNoChange) {
			return err
		}
	case "down":
		if err := m.Down(); err != nil && errors.Is(err, migrate.ErrNoChange) {
			return err
		}
	default:
		return fmt.Errorf("migration type is not valid, %s", mType)
	}

	return nil
}
