package model

// DomainUser struct for using in database
type DomainUser struct {
	ID          string `db:"id"`
	FullName    string `db:"full_name"`
	ImageURL    string `db:"image_url"`
	PhoneNumber string `db:"phone_number"`
	Email       string `db:"email"`
	Occupation  string `db:"occupation"`
	Address     string `db:"address"`
	ProfileInfo string `db:"profile_info"`
}
