package learningtopics

import (
	"EduKita/modules/learning_topics/routes"
	"database/sql"

	"github.com/go-chi/chi"
)

func StartLearningTopicsModule(router chi.Router, db *sql.DB) {
	routes.SetUpLearningTopicsRoutes(router, db)
}
