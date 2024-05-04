package model

type LoginModel struct {
	Name           string `json:"name" validate:"required"`
	Email          string `json:"email" validate:"required"`
	ProfilePicture string `json:"profilePicture" `
}
