package main

import (
	"os"

	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	app := NewApp(host, port)
	app.Run()
}
