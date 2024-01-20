package model

type ContactInput struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone"`
	PhoneCC     string `json:"phonecc"`
	Address     string `json:"street"`
	City        string `json:"city"`
	Zip         string `json:"zip"`
	Country     string `json:"country"`
}
