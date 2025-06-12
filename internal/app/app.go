package app

import (
	"database/sql"
	"fmt"
	"github.com/Viilih/go-crud-api/internal/api"
	"github.com/Viilih/go-crud-api/internal/store"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {

	pgDb, err := store.OpenDatabase()

	if err != nil {
		return nil, err
	}

	logger := log.New(os.Stdout, "[app] ", log.Ldate|log.Ltime)

	// Stores go here (wtf is that)

	// Handlers here
	workoutHandler := api.NewWorkoutHandler()

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDb,
	}
	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is avaiable\n")
}
