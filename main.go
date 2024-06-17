package main

import (
	"EduKita/modules/auth"
	"EduKita/modules/core"
	"EduKita/modules/firebase"
	"EduKita/modules/leaderboard"
	learningtopics "EduKita/modules/learning_topics"
	"EduKita/modules/profile"
	"EduKita/modules/question"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r, db := core.StartCore()
	firebaseMiddleware := firebase.StartFirebaseModule()
	authMiddleware := auth.StartAuthMiddleware()

	r.Group(func(r chi.Router) {
		r.Route(
			"/api/v1/auth",
			func(r chi.Router) {
				auth.StartAuthModule(r, db, firebaseMiddleware, authMiddleware)
			},
		)
	})

	r.Group(func(r chi.Router) {
		r.Route(
			"/api/v1",
			func(r chi.Router) {
				r.Use(authMiddleware.AuthMiddleware)

				learningtopics.StartLearningTopicsModule(r, db)
				question.StartQuestionModule(r, db)
				profile.StartProfileModule(r, db)
				leaderboard.StartLeaderboardModule(r, db)

			},
		)
	})

	http.ListenAndServe(":8080", r)
}
