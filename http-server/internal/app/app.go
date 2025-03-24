package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/oustaa/http-server-go/internal/api"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// handlers
	workoutHandler := api.NewWorkoutHandler()

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
	}

	return app, nil
}

func (a *Application) HeathCheck(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	name := queryParams.Get("name")

	fmt.Println(name)
	fmt.Println(queryParams.Encode())

	fmt.Fprint(w, "Status is available")
}
