package main

import (
	"encoding/csv"
	"log"
	"os"
)

// to initialize regions
func (l *LocationData) Init() {

	l.AvailableLocations = make(map[string]map[string]map[string]struct{}, 0)

	// opening the csv file
	file, err := os.Open("cities.csv")
	// handling the error
	if err != nil {
		// handle the error
		log.Fatal("Error while openting csv ", err)
	}
	defer file.Close()

	locationReader := csv.NewReader(file)
	rowOne := true
	for {
		row, err := locationReader.Read()
		if err != nil {
			// handle error
			break
		}

		if rowOne {
			rowOne = false
			continue
		}

		if _, ok := l.AvailableLocations[row[2]]; !ok {
			l.AvailableLocations[row[2]] = make(map[string]map[string]struct{}, 0)
		}

		if _, ok := l.AvailableLocations[row[2]][row[1]]; !ok {
			l.AvailableLocations[row[2]][row[1]] = make(map[string]struct{}, 0)
		}

		if _, ok := l.AvailableLocations[row[2]][row[1]][row[0]]; !ok {
			l.AvailableLocations[row[2]][row[1]][row[0]] = struct{}{}
		}
	}
}
