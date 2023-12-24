package utils

import (
	"dis1/models"
	"encoding/csv"
	"os"
)

var locations map[models.Location]struct{}

func init() {
	locations = make(map[models.Location]struct{})
}

// returns the enclosing region for the given region and field.
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

// reads locations from the CSV file and returns a map.
func ReadLocations(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

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

		// Constructing a key for the map based on the chosen inquiry field.
		key := location

		locations[key] = struct{}{}
	}

	return nil
}
