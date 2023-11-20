package service

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

const CITIES_DB = "cities.csv"

type App struct {
	Config  *Config
	Logger  *log.Logger
	Dataset map[string][]string
}

type Config struct {
	Port string
}

func NewApp() *App {
	app := &App{
		Config: NewConfig(),
		Logger: log.New(os.Stdout, "challenge2016 - ", log.Ldate|log.Ltime),
	}

	dataSet, err := LoadDataset()
	if err != nil {
		app.Logger.Println("error starting server")
	}
	app.Dataset = dataSet

	return app
}

func NewConfig() *Config {
	return &Config{
		Port: "8080",
	}
}

func LoadDataset() (dataSet map[string][]string, err error) {
	dataSet = make(map[string][]string)

	file, err := os.Open(CITIES_DB)
	if err != nil {
		fmt.Println(err)
		return dataSet, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
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
