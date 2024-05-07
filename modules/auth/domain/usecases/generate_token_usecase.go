package usecases

import (
	"EduKita/modules/auth/data/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type GenerateTokenUsecase interface {
	GenerateJWT(userId string) (model.JwtTokenModel, error)
}

type GenerateTokenUsecaseImpl struct {
	SecretKey        []byte
	RefreshSecretKey []byte
}

func NewGenerateTokenUsecase() *GenerateTokenUsecaseImpl {
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	refreshSecretKey := []byte(os.Getenv("REFRESH_SECRET_KEY"))
	return &GenerateTokenUsecaseImpl{
		SecretKey:        secretKey,
		RefreshSecretKey: refreshSecretKey,
	}
}

func (u *GenerateTokenUsecaseImpl) GenerateJWT(userId string) (model.JwtTokenModel, error) {
	accessClaims := jwt.MapClaims{}
	accessClaims["user_id"] = userId
	accessClaims["exp"] = time.Now().Add(1 * time.Hour).Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessSignedToken, err := accessToken.SignedString(u.SecretKey)
	if err != nil {
		return model.JwtTokenModel{}, err
	}

	refreshClaims := jwt.MapClaims{}
	refreshClaims["user_id"] = userId
	refreshClaims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshSignedToken, err := refreshToken.SignedString(u.RefreshSecretKey)
	if err != nil {
		return model.JwtTokenModel{}, err
	}

	return model.JwtTokenModel{AccessToken: accessSignedToken, RefreshToken: refreshSignedToken}, nil
}
