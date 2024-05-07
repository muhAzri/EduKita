package main

import (
	"EduKita/modules/auth"
	"EduKita/modules/core"
	"EduKita/modules/firebase"
	learningtopics "EduKita/modules/learning_topics"
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
				auth.StartAuthModule(r, db, firebaseMiddleware)
			},
		)
	})

	r.Group(func(r chi.Router) {
		r.Route(
			"/api/v1",
			func(r chi.Router) {
				r.Use(authMiddleware.AuthMiddleware)

				learningtopics.StartLearningTopicsModule(r, db)

			},
		)
	})

	http.ListenAndServe(":8080", r)
}
