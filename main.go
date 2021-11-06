package main

import (
	"log"

	"github.com/rs/zerolog"
)

func main() {
	env := NewEnvironment()
	if env.IsDevelopment() {
		log.SetOutput(zerolog.NewConsoleWriter())
	}

	app := NewApp(env)
	app.Run()
}
