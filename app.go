package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
)

type App struct {
	env    *Environment
	logger zerolog.Logger
}

func NewApp(env *Environment) *App {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	if env.IsDevelopment() {
		cw := zerolog.NewConsoleWriter()
		cw.TimeFormat = time.RFC822Z
		logger = zerolog.New(cw).With().Timestamp().Logger()
	}
	return &App{env: env, logger: logger}
}

func (app App) Run() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signals
		app.logger.Print("Shutting down")
		os.Exit(0)
	}()

	fmt.Println("Server: https://go-todo.myshopify.io")

	app.logger.Info().
		Str("host_port", app.hostPort()).
		Str("mode", app.env.Mode).
		Msg("listening for connections")

	err := http.ListenAndServe(app.hostPort(), nil)
	app.logger.Error().Msg(err.Error())
}

func (app App) hostPort() string {
	return net.JoinHostPort(app.env.Host, app.env.Port)
}
