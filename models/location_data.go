package models

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)


func CheckCode(code string) bool{


	dir, err := os.Getwd()
	
	fmt.Println(dir)
	filePath := filepath.Join(dir, "cities.csv")

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return false
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return false
	}

	// Create a slice to store city names
	var cities,states,countries []string

	// Iterate through the records and extract city names
	for _, record := range records {
		cityCode := record[0]
		provinceCode := record[1]
		countryCode := record[2]
		// Format the city name
		var city string= fmt.Sprintf("%s-%s-%s", cityCode, provinceCode, countryCode)
		var state string= fmt.Sprintf("%s-%s", provinceCode, countryCode)
		var country string= fmt.Sprint(countryCode)

		// Add the city name to the slice
		cities = append(cities, city)
		states = append(states, state)
		countries=append(countries, country)

	}
	
	for _, city := range cities {
		if city==code{
			return true
		}
	}
	for _, state := range states {
		if state==code{
			return true
		}
	}

	for _, country := range countries {
		if country==code{
			return true
		}
	}
	
	return false
}
