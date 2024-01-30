package main

import (
	"log"

	"github.com/dahaev/bottesting/internal/app"
)

func main() {
	application, err := app.New()
	if err != nil {
		log.Fatal("cannot start application", err)
	}
	application.Start()
}
