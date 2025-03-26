package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/oustaa/http-server-go/internal/api"
	"github.com/oustaa/http-server-go/internal/store"
	"github.com/oustaa/http-server-go/migrations"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	// DB             *sql.DB
}

func NewApplication() (*Application, error) {
    // pgDB, err := store.Open()

	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// handlers
	workoutHandler := api.NewWorkoutHandler()

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		// DB:             pgDB,
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
