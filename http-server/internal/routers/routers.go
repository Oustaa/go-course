package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/oustaa/http-server-go/internal/app"
)

func SetupRouters(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/heath-check", app.HeathCheck)
	// workouts routes
	r.Get("/workouts/{id}", app.WorkoutHandler.GetWorkoutById)
	r.Delete("/workouts/{id}", app.WorkoutHandler.DeleteWorkout)
	r.Post("/workouts", app.WorkoutHandler.CreateWorkout)
	r.Put("/workouts/{id}", app.WorkoutHandler.UpdateWorkout)

	return r
}
