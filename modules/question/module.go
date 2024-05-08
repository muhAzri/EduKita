package question

import (
	"EduKita/modules/question/routes"
	"database/sql"

	"github.com/go-chi/chi"
)

func StartQuestionModule(r chi.Router, db *sql.DB) {

	routes.SetUpLearningTopicsRoutes(r, db)
}
