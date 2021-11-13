package env

import "os"

type Environment struct{}

func NewEnvironment() *Environment {
	return &Environment{}
}

func (e Environment) Get(key string) string {
	return os.Getenv(key)
}

func (e Environment) IsDevelopment() bool {
	return e.Get("MODE") == "development"
}

func (e Environment) IsProduction() bool {
	return e.Get("MODE") == "production"
}

func (e Environment) IsTest() bool {
	return e.Get("MODE") == "test"
}
