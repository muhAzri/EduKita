package auth

import (
	"EduKita/modules/auth/domain/entity"
	"EduKita/modules/auth/routes"
	"database/sql"

	"github.com/go-chi/chi"
)

func StartAuthModule(router chi.Router, db *sql.DB) {
	entity.MigrateUser(db)

	routes.SetUpAuthRoutes(router, db)
}
