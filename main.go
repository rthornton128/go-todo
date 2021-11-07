package main

func main() {
	env := NewEnvironment()
	app := NewApp(env)

	app.Run()
}
