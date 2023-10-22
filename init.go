package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Load region data from CSV file
func LoadRegionData() []Region {
	file, err := os.Open("cities.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return nil
	}

	var regions []Region
	for _, record := range records {
		region := Region{
			CityCode:     record[0],
			ProvinceCode: record[1],
			CountryCode:  record[2],
			CityName:     record[3],
			ProvinceName: record[4],
			CountryName:  record[5],
		}
		regions = append(regions, region)
	}

	return regions
}