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

func NewDB() *sqlx.DB {
	cfg := config.Get()
	url := NewURL(cfg, "postgres")

	// Connection for connection database with sqlx
	db, err := sqlx.Connect("postgres", url)
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
	url := NewURL(cfg, "migration")
	// Generate migration for migrate action
	m, err := migrate.New("file://migrations", url)
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
