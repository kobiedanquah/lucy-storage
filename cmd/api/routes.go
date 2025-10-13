package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func loadRoutes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	
	mux.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("service online"))
	})

	return mux
}
