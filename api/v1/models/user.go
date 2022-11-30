package models

type UserModel struct {
	ID          string `json:"id"`
	FullName    string `json:"full_name"`
	ImageURL    string `json:"image_url"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Occupation  string `json:"occupation"`
	Address     string `json:"address"`
	ProfileInfo string `json:"profile_info"`
}
