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

	for _, c := range countries {
		if c.Name == countryName {
			currentCountry = c
		}
	}

	if currentCountry == nil {
		currentCity = &city{Code: cityCode, Name: cityName}
		currentProvince = &province{Code: provinceCode, Name: provinceName, Cities: []*city{currentCity}}
		currentCountry = &country{Code: countryCode, Name: countryName, Provinces: []*province{currentProvince}}
		countries = append(countries, currentCountry)
		return
	}

	for _, p := range currentCountry.Provinces {
		if p.Name == provinceName {
			currentProvince = p
		}
	}

	if currentProvince == nil {
		currentCity = &city{Code: cityCode, Name: cityName}
		currentProvince = &province{Code: provinceCode, Name: provinceName, Cities: []*city{currentCity}}
		currentCountry.Provinces = append(currentCountry.Provinces, currentProvince)
		return
	}

	for _, ci := range currentProvince.Cities {
		if ci.Name == cityName {
			return
		}
	}

	currentCity = &city{Code: cityCode, Name: cityName}
	currentProvince.Cities = append(currentProvince.Cities, currentCity)

}

func (loc location) String() string {

	return fmt.Sprintf("Country: %s, Province: %s, City: %s", loc.CountryName, loc.ProvinceName, loc.CityName)

}

func getCountry(countryName, provinceName, cityName string) (country, error) {

	for _, c := range countries {

		if countryName != "" && c.Name == countryName {

			if provinceName == "" {
				tempCountry := copyCountry(*c)
				return tempCountry, nil
			}
			outCountry := country{}
			outCountry.Name = c.Name
			outCountry.Code = c.Code

			for _, p := range c.Provinces {
				if p.Name == provinceName {
					if cityName == "" {
						tempProvince := copyProvince(*p)
						outCountry.Provinces = append(outCountry.Provinces, &tempProvince)
						return outCountry, nil
					}

					outProvince := province{}
					outProvince.Name = p.Name
					outProvince.Code = p.Code

					for _, ci := range p.Cities {
						if ci.Name == cityName {
							outCity := city{}
							outCity.Name = ci.Name
							outCity.Code = ci.Code
							outProvince.Cities = append(outProvince.Cities, &outCity)
							outCountry.Provinces = append(outCountry.Provinces, &outProvince)
							return outCountry, nil
						}

					}
				}

			}

		}

	}

	return country{}, errors.New(locationNotFound)

}

func copyCountry(c country) country {
	tempCountry := country{}
	tempCountry.Code = c.Code
	tempCountry.Name = c.Name
	for _, p := range c.Provinces {

		tempProvince := copyProvince(*p)
		tempCountry.Provinces = append(tempCountry.Provinces, &tempProvince)
	}
	return tempCountry
}

func copyProvince(p province) province {

	tempProvince := province{}
	tempProvince.Name = p.Name
	tempProvince.Code = p.Code
	for _, ci := range p.Cities {
		tempCity := city{}
		tempCity.Name = ci.Name
		tempCity.Code = ci.Code
		tempProvince.Cities = append(tempProvince.Cities, &tempCity)

	}
	return tempProvince
}
