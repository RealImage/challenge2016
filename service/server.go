package service

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

const CITIES_DB = "cities.csv"

// App - app struct for initialisation
type App struct {
	Config  *Config
	Dataset map[string][]string
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
	dataSet, err := LoadDataset()
	if err != nil {
		log.Println("error starting server")
	}
	app.Dataset = dataSet

	return app
}

func NewConfig() *Config {
	return &Config{
		Port: "8080",
	}
}

// LoadDataset - loading the dataset into memory
func LoadDataset() (dataSet map[string][]string, err error) {
	dataSet = make(map[string][]string)

	file, err := os.Open(CITIES_DB)
	if err != nil {
		log.Println(err)
		return dataSet, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		return dataSet, err
	}

	for count, record := range records {
		if count > 0 {
			key := record[5]
			value := strings.Join(record, "-")
			dataSet[key] = append(dataSet[key], value)
		}
	}

	return dataSet, err
}
