package main

import "os"

type Environment struct {
	DataSourceName string // Database source name string; example "mysql
	Host           string // Host to listen on: "localhost"
	Mode           string // Server mode; example: "production" or "development"
	Port           string // Port to listen on: "80"
}

func NewEnvironment() *Environment {
	dataSourceName := os.Getenv("DSN")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	mode, ok := os.LookupEnv("MODE")
	if !ok {
		mode = "production"
	}

	return &Environment{
		DataSourceName: dataSourceName,
		Host:           host,
		Mode:           mode,
		Port:           port,
	}
}

func (e Environment) IsDevelopment() bool {
	return e.Mode == "development"
}

func (e Environment) IsProduction() bool {
	return e.Mode == "production"
}
