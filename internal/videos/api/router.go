package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Router(h Handlers) chi.Router {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: false,
	}))
	r.Post("/search", h.filterVectorID)
	return r
}
