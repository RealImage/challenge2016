package csvhandler

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type State struct {
	cities map[string]bool
}

type Country struct {
	states map[string]*State
}

type CsvFileContent struct {
	CityMap map[string]*Country
}

const filePath string = "cities.csv"

func (c *CsvFileContent) LoadCsv() error {

	file, err := os.Open(filePath)
	if err != nil {
		return errors.New("error while reading file")
	}

	defer file.Close()

	fileReader := csv.NewReader(file)

	_, err = fileReader.Read()
	if err != nil {
		return errors.New("error while reading file")
	}

	for {
		record, err := fileReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return errors.New("error while reading file")
		}

		countryName := strings.ToLower(strings.ReplaceAll(record[5], " ", ""))
		stateName := strings.ToLower(strings.ReplaceAll(record[4], " ", ""))
		cityName := strings.ToLower(strings.ReplaceAll(record[3], " ", ""))

		// Check if the country exists
		country, exists := c.CityMap[countryName]
		if !exists {
			country = &Country{
				states: make(map[string]*State),
			}
			c.CityMap[countryName] = country

		}

		// Check if the state exists
		state, exists := c.CityMap[countryName].states[stateName]
		if !exists {
			state = &State{
				cities: make(map[string]bool),
			}
			country.states[stateName] = state
		}

		// Add the city
		state.cities[cityName] = true
	}

	return nil

}

func (c *CsvFileContent) ValidateInputRegion(countryName, state, city string) error {

	if countryName == "" {
		return fmt.Errorf("countryName should not be empty")
	}

	country, countryExists := c.CityMap[countryName]
	if !countryExists {
		return fmt.Errorf("No country exists with given name: %s", countryName)
	}

	if state == "" {
		return nil // No need to check states or cities if state is not provided.
	}

	stateData, stateExists := country.states[state]
	if !stateExists {
		return fmt.Errorf("No state exists with given name: %s", state)
	}

	if city == "" {
		return nil // No need to check cities if city is not provided.
	}

	if _, ok := stateData.cities[city]; !ok {
		return fmt.Errorf("No city exists with given name: %s", city)
	}

	return nil
}
