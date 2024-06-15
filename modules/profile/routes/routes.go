package routes

import (
	"EduKita/modules/profile/wire"
	"database/sql"

	"github.com/go-chi/chi"
)

func SetUpProfileRoutes(router chi.Router, db *sql.DB) {

	profileHandler := wire.InitializeProfileHandler(db)

	router.Get("/profile", profileHandler.GetProfile)
}
