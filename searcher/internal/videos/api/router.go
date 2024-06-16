package api

import (
	"github.com/go-chi/chi/v5"
)

func Router(h Handlers) chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.filterVectorID)
	return r
}
