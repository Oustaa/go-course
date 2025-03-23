package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/oustaa/http-server-go/internal/app"
)

func SetupRouters(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/heath-check", app.HeathCheck)

	return r
}
