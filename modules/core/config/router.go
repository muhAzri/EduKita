package config

import (
	"github.com/MadAppGang/httplog"
	"github.com/go-chi/chi"
)

func InitializeRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(httplog.Logger)

	return router
}
