package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct {
}

func NewWorkoutHandler() *WorkoutHandler {
	workoutHandler := WorkoutHandler{}

	return &workoutHandler
}

func (wh *WorkoutHandler) GetWorkoutById(w http.ResponseWriter, r *http.Request) {
	paramsWorkoutID := chi.URLParam(r, "id")

	if paramsWorkoutID == "" {
		http.NotFound(w, r)
		return
	}

	workoutId, err := strconv.ParseInt(paramsWorkoutID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "workout id is %d\n", workoutId)
}

func (wh *WorkoutHandler) DeleteWorkout(w http.ResponseWriter, r * http.Request){
    paramsWorkoutID := chi.URLParam(r, "id")
    if paramsWorkoutID == "" {
        http.NotFound(w, r)
        return
    }

    workoutId, err := strconv.ParseInt(paramsWorkoutID, 10, 64)
    if err != nil {
        http.NotFound(w, r)
        return
    }

    fmt.Fprintf(w, "the workout with id %d was deleted with success", workoutId)
}

func (wh *WorkoutHandler) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Workout created with success\n")
}
