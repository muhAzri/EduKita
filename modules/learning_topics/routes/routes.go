package routes

import (
	"EduKita/modules/learning_topics/wire"
	"database/sql"

	"github.com/go-chi/chi"
)

func SetUpLearningTopicsRoutes(router chi.Router, db *sql.DB) {

	learningTopicsHandler := wire.InitializeLearningTopicsHandler(db)

	router.Get("/learning-topics", learningTopicsHandler.GetAllLearningTopics)
}
