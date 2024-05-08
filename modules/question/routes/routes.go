package routes

import (
	"EduKita/modules/question/wire"
	"database/sql"

	"github.com/go-chi/chi"
)

func SetUpLearningTopicsRoutes(router chi.Router, db *sql.DB) {

	questionHandler := wire.InitializeQuestionHandler(db)

	router.Get("/questions/{learning_topic_id}", questionHandler.GetQuestionByLearningTopic)
	router.Get("/questions/{learning_topic_id}/quiz", questionHandler.Get10RandomQuestionByLearningTopic)
}
