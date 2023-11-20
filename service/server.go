package service

import (
	"log"
)

const CITIES_DB = "cities.csv"

// App - app struct for initialisation
type App struct {
	Config *Config
}

// Config - config struct for several configuration for app
type Config struct {
	Port string
}

func NewApp() *App {
	app := &App{
		Config: NewConfig(),
	}

	//load the dataset into memory
	err := LoadDataset()
	if err != nil {
		log.Println("error starting server")
	}

	return app
}

func NewConfig() *Config {
	return &Config{
		Port: "8080",
	}
}
