package service

import (
	"github.com/AbduRohman-developer/cv-backend/storage/repository"
)

type service struct {
	repo repository.Repository
}

func New(r repository.Repository) *service {
	return &service{
		repo: r,
	}
}
