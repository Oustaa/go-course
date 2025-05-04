package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/oustaa/url-shortner/internal/views"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", views.MainView)

	return r
}
