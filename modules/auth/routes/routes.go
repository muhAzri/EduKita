package routes

import (
	"EduKita/modules/auth/wire"
	"database/sql"

	"github.com/go-chi/chi"
)

func SetUpAuthRoutes(router chi.Router, db *sql.DB) {
	authHandler := wire.InitializeAuthHandler(db)

	router.Post("/sessions", authHandler.Login)
}
