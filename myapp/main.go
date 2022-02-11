package main

import (
	"github.com/RedHoodJT1988/velocidad"
	"myapp/handlers"
)

type application struct {
	App      *velocidad.Velocidad
	Handlers *handlers.Handlers
}

func main() {
	v := initApplication()
	v.App.ListenAndServe()
}
