package routes

import (
	"EduKita/modules/auth/wire"
	"EduKita/modules/firebase/middleware"
	"database/sql"

	"github.com/go-chi/chi"
)

func SetUpAuthRoutes(router chi.Router, db *sql.DB, firebaseMiddleware middleware.FirebaseMiddleware) {
	authHandler := wire.InitializeAuthHandler(db)

	r := router.Group(nil)
	r.Use(firebaseMiddleware.FirebaseAuthmiddleware)
	r.Post("/sessions", authHandler.Login)
}
