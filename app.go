package main

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

type App struct {
	env *Environment
}

func NewApp(env *Environment) *App {
	return &App{env: env}
}

func (app App) Run() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signals
		log.Print("Shutting down")
		os.Exit(0)
	}()

	log.Print("Starting app")
	log.Error().Msg(http.ListenAndServe(app.hostPort(), nil).Error())
}

func (app App) hostPort() string {
	return net.JoinHostPort(app.env.Host, app.env.Port)
}
