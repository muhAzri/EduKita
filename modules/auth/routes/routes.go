package routes

import (
	"EduKita/modules/auth/middleware"
	"EduKita/modules/auth/wire"
	firebaseMiddleware "EduKita/modules/firebase/middleware"
	"database/sql"

	"github.com/go-chi/chi"
)

func SetUpAuthRoutes(router chi.Router, db *sql.DB, firebaseMiddleware firebaseMiddleware.FirebaseMiddleware, authMiddleware middleware.AuthMiddleware) {
	authHandler := wire.InitializeAuthHandler(db)

	sessionsRoute := router.Group(nil)
	sessionsRoute.Use(firebaseMiddleware.FirebaseAuthmiddleware)
	sessionsRoute.Post("/sessions", authHandler.Login)

	userRoute := router.Group(nil)
	userRoute.Use(authMiddleware.AuthMiddleware)
	userRoute.Get("/user-short", authHandler.GetShortProfile)

	router.Post("/refresh", authHandler.Refresh)
}
