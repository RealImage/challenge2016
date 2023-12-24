package utils

import (
	"RealImage/models"
	"encoding/csv"
	"os"
)

// Data structure to store all the data from the CSV
var locations map[models.Location]struct{}

func init() {
	locations = make(map[models.Location]struct{})
}

// GetEnclosingRegion returns the enclosing region for the given location.
func GetEnclosingRegion(location models.Location) models.Location {
	if location.CityCode != "" && location.CityName != "" {
		return models.Location{
			ProvinceCode: location.ProvinceCode,
			ProvinceName: location.ProvinceName,
			CountryCode:  location.CountryCode,
			CountryName:  location.CountryName,
		}
	}
	if location.ProvinceCode != "" && location.ProvinceName != "" {
		return models.Location{
			CountryCode: location.CountryCode,
			CountryName: location.CountryName,
		}
	}
	return models.Location{}
}

// ReadLocations reads locations from the CSV file and returns a map.
func ReadLocations(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Assuming the CSV has columns for City Code, Province Code, Country Code,
	// City Name, Province Name, and Country Name.
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// Constructing location using the required columns.
		location := models.Location{
			CityCode:     record[0],
			ProvinceCode: record[1],
			CountryCode:  record[2],
			CityName:     record[3],
			ProvinceName: record[4],
			CountryName:  record[5],
		}

		key := location
		locations[key] = struct{}{}
	}
	return nil
}
