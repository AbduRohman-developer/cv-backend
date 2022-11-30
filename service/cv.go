package service

import "github.com/AbduRohman-developer/cv-backend/storage"

type service struct {
	repo storage.Repository
}

func New(r storage.Repository) *service {
	return &service{
		repo: r,
	}
}
