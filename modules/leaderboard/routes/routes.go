package routes

import (
	"EduKita/modules/leaderboard/wire"
	"database/sql"

	"github.com/go-chi/chi"
)

func SetUpLeaderboardRoutes(router chi.Router, db *sql.DB) {

	leaderboardHandler := wire.InitializeLeaderboardHandler(db)

	router.Get("/leaderboard", leaderboardHandler.GetLeaderboard)
}
