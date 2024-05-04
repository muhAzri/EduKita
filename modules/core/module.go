package core

import (
	"EduKita/modules/core/config"
	"database/sql"

	"github.com/go-chi/chi"
)

func StartCore() (*chi.Mux, *sql.DB) {
	config.InitializeEnvironment()

	router := config.InitializeRouter()

	db, err := config.InitializeDatabase()

	if err != nil {
		panic(err)
	}

	return router, db
}
