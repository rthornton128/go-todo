package main

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rthornton128/go-todo/pkg/env"
	"github.com/rthornton128/go-todo/pkg/store"
)

type App struct {
	*env.Environment

	db     *store.Database
	logger zerolog.Logger
}

func NewApp(env *env.Environment) *App {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	if env.IsDevelopment() {
		cw := zerolog.NewConsoleWriter()
		cw.TimeFormat = time.RFC822Z
		logger = zerolog.New(cw).With().Timestamp().Logger()
	}

	db, err := store.NewDatabase(env.Get("DSN"))
	if err != nil {
		logger.Fatal().Err(err).Msg("opening database")
	}
	logger.Info().Msg("connected to DB")

	return &App{Environment: env, db: db, logger: logger}
}

func (app App) Run() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signals
		app.logger.Print("Shutting down")
		os.Exit(0)
	}()

	app.logger.Info().
		Str("host_port", app.hostPort()).
		Str("mode", app.Get("MODE")).
		Msg("listening for connections")

	err := http.ListenAndServe(app.hostPort(), nil)
	app.logger.Error().Msg(err.Error())
}

func (app App) hostPort() string {
	return net.JoinHostPort(app.Get("HOST"), app.Get("PORT"))
}
