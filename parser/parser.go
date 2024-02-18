package parser

import (
	"challenge2016/dto" // Importing DTO package for data transfer objects
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

func ParseCSVFile(csvFilePath string) ([]dto.Country, error) {
	locationsFile, err := os.OpenFile(csvFilePath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer locationsFile.Close()

	locations := []*dto.Location{}

	if err := gocsv.UnmarshalFile(locationsFile, &locations); err != nil { // Load clients from file
		return nil, err
	}

	groupedData := make([]dto.Country, 0)

	for _, location := range locations {
		countryName := strings.ToUpper(location.CountryName)
		stateName := strings.ToUpper(location.ProvinceName)
		cityName := strings.ToUpper(location.CityName)

		var countryIndex int
		countryExists := false
		for i, country := range groupedData {
			if strings.EqualFold(country.Name, countryName) {
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
				if strings.EqualFold(state.Name, stateName) {
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
