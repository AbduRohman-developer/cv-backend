package storage

import (
	"github.com/AbduRohman-developer/cv-backend/pkg/connection"
	"github.com/jmoiron/sqlx"
)

type postgres struct {
	DB *sqlx.DB
}

func NewPostgres() *postgres {
	return &postgres{
		DB: connection.New(),
	}
}
