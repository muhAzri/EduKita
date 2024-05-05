package auth

import (
	"EduKita/modules/auth/routes"
	"database/sql"

	"github.com/go-chi/chi"
)

func StartAuthModule(router chi.Router, db *sql.DB) {

	routes.SetUpAuthRoutes(router, db)
}
