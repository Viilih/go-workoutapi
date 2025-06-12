package main

import (
	"net/http"
	"time"

	"github.com/Viilih/go-crud-api/internal/app"
	"github.com/Viilih/go-crud-api/internal/routes"
)

func main() {
	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}
	defer app.DB.Close() // At the end of the execution it calls the defer function (pretty cool stuff)
	app.Logger.Println("Starting application")
	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
