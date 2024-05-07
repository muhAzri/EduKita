package middleware

import (
	"EduKita/modules/core/constants"
	"EduKita/modules/core/response"
	"context"
	"fmt"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
)

type FirebaseMiddleware struct {
	AuthClient auth.Client
}

func (fm *FirebaseMiddleware) FirebaseAuthmiddleware(next http.Handler) http.Handler {
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

		decoded, err := fm.AuthClient.VerifyIDToken(context.Background(), token)
		if err != nil {
			fmt.Println(err)
			if strings.Contains(err.Error(), "ID token has expired") {
				response.BuildResponseFailure(http.StatusUnauthorized, "Token has expired", w)
			} else {
				response.BuildResponseFailure(http.StatusUnauthorized, "Invalid Token", w)
			}
			return
		}

		ctx := context.WithValue(r.Context(), constants.FirebaseIDKey, decoded.UID)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
