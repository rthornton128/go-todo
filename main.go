package main

import "github.com/rthornton128/go-todo/pkg/env"

func main() {
	env := env.NewEnvironment()
	app := NewApp(env)

	app.Run()
}
