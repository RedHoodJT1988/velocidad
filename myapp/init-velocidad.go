package main

import (
	"log"
	"myapp/handlers"
	"os"

	"github.com/RedHoodJT1988/velocidad"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init velocidad
	vel := &velocidad.Velocidad{}
	err = vel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	vel.AppName = "myapp"

	myHandlers := &handlers.Handlers{
		App: vel,
	}

	vel.InfoLog.Println("Debug is set to", vel.Debug)

	app := &application{
		App:      vel,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()

	return app
}
