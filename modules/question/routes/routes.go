package routes

import (
	"EduKita/modules/question/wire"
	"database/sql"

	"github.com/go-chi/chi"
)

func SetUpLearningTopicsRoutes(router chi.Router, db *sql.DB) {

	questionHandler := wire.InitializeQuestionHandler(db)
	answerHandler := wire.InitializeAnswerHandler(db)
	adminHandler := wire.InitializeAdminQuestionHandler(db)

	router.Get("/questions/{learning_topic_id}", questionHandler.GetQuestionByLearningTopic)
	router.Get("/questions/{learning_topic_id}/quiz", questionHandler.Get10RandomQuestionByLearningTopic)
	router.Post("/questions/answer", answerHandler.AnswerQuestion)

	router.Get("/questions/admin/{question_id}", adminHandler.GetQuestion)
	router.Delete("/questions/admin/{question_id}", adminHandler.DeleteQuestion)
	router.Post("/questions/admin", adminHandler.CreateQuestion)
	router.Put("/questions/admin/{question_id}", adminHandler.UpdateQuestion)
}
