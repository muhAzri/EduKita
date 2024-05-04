package main

import (
	"EduKita/modules/auth"
	"EduKita/modules/core"
	"EduKita/modules/firebase"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r, db := core.StartCore()
	firebaseMiddleware := firebase.StartFirebaseModule()

	r.Group(func(r chi.Router) {
		r.Use(firebaseMiddleware.AuthMiddleware)
		r.Route(
			"/api/v1",
			func(r chi.Router) {
				// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				// 	response.BuildResponseSuccess(http.StatusOK, "Welcome to EduKita", "success", map[string]interface{}{"message": "Welcome to EduKita"}, w)
				// })
				auth.StartAuthModule(r, db)

			},
		)
	})

	http.ListenAndServe(":8080", r)
}
