package leaderboard

import (
	"EduKita/modules/leaderboard/routes"
	"database/sql"

	"github.com/go-chi/chi"
)

func StartLeaderboardModule(router chi.Router, db *sql.DB) {

	routes.SetUpLeaderboardRoutes(router, db)
}
