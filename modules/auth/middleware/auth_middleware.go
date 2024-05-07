package middleware

import (
	"EduKita/modules/auth/domain/usecases"
	"EduKita/modules/core/constants"
	"EduKita/modules/core/response"
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

type AuthMiddleware struct {
	ValidateTokenUsecase usecases.ValidateTokenUsecase
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		ValidateTokenUsecase: usecases.NewValidateTokenUsecase(),
	}
}

func (am *AuthMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.BuildResponseFailure(http.StatusUnauthorized, "Authorization header is missing", w)
			return
		}

		headerParts := strings.Split(authHeader, " ")

		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			response.BuildResponseFailure(http.StatusUnauthorized, "Invalid authorization header format", w)
			return
		}

		token := headerParts[1]

		jwtToken, err := am.ValidateTokenUsecase.ValidateJWT(token)

		if err != nil {
			response.BuildResponseFailure(http.StatusUnauthorized, err.Error(), w)
			return
		}

		claim, ok := jwtToken.Claims.(jwt.MapClaims)
		if !ok || !jwtToken.Valid {
			response.BuildResponseFailure(http.StatusUnauthorized, "Invalid token", w)
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, constants.UserIDKey, claim["user_id"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
