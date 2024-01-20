package model

type AuthenticationInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
