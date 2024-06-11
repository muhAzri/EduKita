package model

type RefreshModel struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
