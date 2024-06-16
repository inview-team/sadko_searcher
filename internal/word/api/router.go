package api

import "github.com/go-chi/chi/v5"

func Router(h Handlers) chi.Router {
	r := chi.NewRouter()
	r.Get("/suggestions", h.getListWord)
	r.Post("/add-word", h.downloadWord)
	return r
}
