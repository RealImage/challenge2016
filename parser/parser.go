package parser

import (
	"challenge2016/dto" // Importing DTO package for data transfer objects
	"encoding/csv"
	"os"
	"strings"
)

// The function `ParseCSVFile` takes a CSV file path as input, reads the file, and parses the data into
// a structured format representing countries, states, and cities.
func ParseCSVFile(csvFilePath string) ([]dto.Country, error) {
	file, err := os.Open(csvFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	groupedData := make([]dto.Country, 0)

	for _, row := range records {
		countryName := strings.ToUpper(row[5])
		stateName := strings.ToUpper(row[4])
		cityName := strings.ToUpper(row[3])

		var countryIndex int
		countryExists := false
		for i, country := range groupedData {
			if country.Name == countryName {
				countryIndex = i
				countryExists = true
				break
			}
		}

		if !countryExists {
			newCountry := dto.Country{
				Name: countryName,
				States: []dto.State{
					{
						Name:   stateName,
						Cities: []dto.City{{Name: cityName}},
					},
				},
			}
			groupedData = append(groupedData, newCountry)
		} else {
			var stateIndex int
			stateExists := false
			for i, state := range groupedData[countryIndex].States {
				if state.Name == stateName {
					stateIndex = i
					stateExists = true
					break
				}
			}

			if !stateExists {
				newState := dto.State{
					Name:   stateName,
					Cities: []dto.City{{Name: cityName}},
				}
				groupedData[countryIndex].States = append(groupedData[countryIndex].States, newState)
			} else {
				groupedData[countryIndex].States[stateIndex].Cities = append(groupedData[countryIndex].States[stateIndex].Cities, dto.City{Name: cityName})
			}
		}
	}

	return groupedData, nil
}
