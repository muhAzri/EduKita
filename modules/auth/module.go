package auth

import (
	"EduKita/modules/auth/middleware"
	"EduKita/modules/auth/routes"
	firebaseMiddleware "EduKita/modules/firebase/middleware"
	"database/sql"

	"github.com/go-chi/chi"
)

func StartAuthModule(router chi.Router, db *sql.DB, firebaseMiddleware firebaseMiddleware.FirebaseMiddleware) {

	routes.SetUpAuthRoutes(router, db, firebaseMiddleware)
}

func StartAuthMiddleware() middleware.AuthMiddleware {

	middleware := middleware.NewAuthMiddleware()

	return *middleware

}
