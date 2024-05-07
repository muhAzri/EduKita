package usecases

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

type ValidateTokenUsecase interface {
	ValidateJWT(JwtToken string) (*jwt.Token, error)
}

type ValidateTokenUsecaseImpl struct {
	SecretKey        []byte
	RefreshSecretKey []byte
}

func NewValidateTokenUsecase() *ValidateTokenUsecaseImpl {
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	refreshSecretKey := []byte(os.Getenv("REFRESH_SECRET_KEY"))
	return &ValidateTokenUsecaseImpl{secretKey, refreshSecretKey}
}

func (u *ValidateTokenUsecaseImpl) ValidateJWT(JwtToken string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(JwtToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return u.SecretKey, nil
	})
	if err != nil {
		return parsedToken, err
	}

	return parsedToken, nil
}
