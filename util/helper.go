package util

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func PopulateMapFromCSV(filePath string) (map[string]map[string]map[string]struct{}, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	//  parse the CSV file
	reader := csv.NewReader(file)

	//create a nested map structure
	data := make(map[string]map[string]map[string]struct{})

	//  read each line of the csv file
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading CSV: %v", err)
		}

		// Extract values from the CSV record
		countryName := strings.ToLower(record[5])
		provinceName := strings.ToLower(record[4])
		cityName := strings.ToLower(record[3])

		// Create map if not exists
		if data[countryName] == nil {
			data[countryName] = make(map[string]map[string]struct{})
		}
		if data[countryName][provinceName] == nil {
			data[countryName][provinceName] = make(map[string]struct{})
		}

		// Add city to the map
		data[countryName][provinceName][cityName] = struct{}{}
	}

	return data, nil
}

func IsValidCountry(countryName string, data map[string]map[string]map[string]struct{}) bool {
	_, exists := data[strings.ToLower(countryName)]
	return exists
}

func IsValidProvince(countryName, provinceName string, data map[string]map[string]map[string]struct{}) bool {
	_, exists := data[strings.ToLower(countryName)][strings.ToLower(provinceName)]
	return exists
}

func IsValidCity(countryName, provinceName, cityName string, data map[string]map[string]map[string]struct{}) bool {
	_, exists := data[strings.ToLower(countryName)][strings.ToLower(provinceName)][strings.ToLower(cityName)]
	return exists
}
