package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"

	"os"
)

func prepareAllLocations() error {

	var count int
	file, err := os.Open("cities1.csv")
	if err != nil {
		return errors.New("error in opening cities file: " + err.Error())

	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	csvReader.Read()
	count++

	for {
		line, err := csvReader.Read()

		if err != nil {
			if err == io.EOF {
				break
			}
			return errors.New("error in reading cities file: " + err.Error())
		}
		count++

		if len(line) != 6 {
			return fmt.Errorf("Not enough parameters in line: %d", count)
		}

		for i, field := range line {
			if field == "" {
				return fmt.Errorf("Field %d is empty in line: %d", i+1, count)
			}
		}

		addToLocation(line)

	}

	return nil
}

func addToLocation(line []string) {

	countryCode, provinceCode, cityCode := line[2], line[1], line[0]
	countryName, provinceName, cityName := line[5], line[4], line[3]

	var currentCountry *country
	var currentProvince *province
	var currentCity *city

	for _, country := range countries {
		if country.Name == countryName {
			currentCountry = country
		}
	}

	if currentCountry == nil {
		currentCity = &city{Code: cityCode, Name: cityName}
		currentProvince = &province{Code: provinceCode, Name: provinceName, Cities: []*city{currentCity}}
		currentCountry = &country{Code: countryCode, Name: countryName, Provinces: []*province{currentProvince}}
		countries = append(countries, currentCountry)
		return
	}

	for _, province := range currentCountry.Provinces {
		if province.Name == provinceName {
			currentProvince = province
		}
	}

	if currentProvince == nil {
		currentCity = &city{Code: cityCode, Name: cityName}
		currentProvince = &province{Code: provinceCode, Name: provinceName, Cities: []*city{currentCity}}
		currentCountry.Provinces = append(currentCountry.Provinces, currentProvince)
		return
	}

	for _, city := range currentProvince.Cities {
		if city.Name == cityName {
			return
		}
	}

	currentCity = &city{Code: cityCode, Name: cityName}
	currentProvince.Cities = append(currentProvince.Cities, currentCity)

}
