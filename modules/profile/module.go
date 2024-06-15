package profile

import (
	"EduKita/modules/profile/routes"
	"database/sql"

	"github.com/go-chi/chi"
)

func StartProfileModule(router chi.Router, db *sql.DB) {

	routes.SetUpProfileRoutes(router, db)
}
