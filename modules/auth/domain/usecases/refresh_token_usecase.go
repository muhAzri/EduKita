package usecases

import (
	"EduKita/modules/auth/data/model"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type RefreshTokenUsecase interface {
	RefreshAccessToken(JwtToken string) (model.JwtTokenModel, error)
}

type RefreshTokenUsecaseImpl struct {
	SecretKey        []byte
	RefreshSecretKey []byte
}

func NewRefreshTokenUsecase() *RefreshTokenUsecaseImpl {
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	refreshSecretKey := []byte(os.Getenv("REFRESH_SECRET_KEY"))
	return &RefreshTokenUsecaseImpl{secretKey, refreshSecretKey}
}

func (u *RefreshTokenUsecaseImpl) RefreshAccessToken(RefrrefeshToken string) (model.JwtTokenModel, error) {
	token, err := jwt.Parse(RefrrefeshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return u.RefreshSecretKey, nil
	})
	if err != nil {
		return model.JwtTokenModel{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return model.JwtTokenModel{}, errors.New("invalid refresh token")
	}

	accessClaims := jwt.MapClaims{}
	accessClaims["user_id"] = claims["user_id"]
	accessClaims["exp"] = time.Now().Add(1 * time.Hour).Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessSignedToken, err := accessToken.SignedString(u.SecretKey)
	if err != nil {
		return model.JwtTokenModel{}, err
	}

	return model.JwtTokenModel{AccessToken: accessSignedToken, RefreshToken: RefrrefeshToken}, nil
}
