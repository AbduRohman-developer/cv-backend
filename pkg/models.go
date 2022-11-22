package pkg

import (
	"log"
	"net/mail"
)

type DomainUser struct {
	ID          string `json:"id"`
	FullName    string `json:"full_name"`
	ImageURL    string `json:"image_url"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Occupation  string `json:"occupation"`
	Address     string `json:"address"`
	ProfileInfo string `json:"profile_info"`
}

func NewUser() *DomainUser {
	return &DomainUser{}
}

func isValidEmail(email string) error {
	a, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}

	log.Println(a.Name, a.Address)
	return nil
}
