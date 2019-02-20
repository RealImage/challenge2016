package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"sort"

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

	_, currentCountry, _ = getCountryFromCountries(countries, countryName)

	if currentCountry == nil {
		currentCity = &city{Code: cityCode, Name: cityName}
		currentProvince = &province{Code: provinceCode, Name: provinceName, Cities: []*city{currentCity}}
		currentCountry = &country{Code: countryCode, Name: countryName, Provinces: []*province{currentProvince}}
		countries = append(countries, currentCountry)
		if len(countries) > 1 {

			sort.SliceStable(countries, func(i, j int) bool {
				if countries[i] != nil && countries[j] != nil {
					return countries[i].Name <= countries[j].Name
				}
				return false
			})

		}
		return
	}

	_, currentProvince, _ = getProvinceFromCountry(currentCountry.Provinces, provinceName)

	if currentProvince == nil {
		currentCity = &city{Code: cityCode, Name: cityName}
		currentProvince = &province{Code: provinceCode, Name: provinceName, Cities: []*city{currentCity}}
		currentCountry.Provinces = append(currentCountry.Provinces, currentProvince)
		if len(currentCountry.Provinces) > 1 {
			sort.SliceStable(currentCountry.Provinces, func(i, j int) bool {
				if currentCountry.Provinces[i] != nil && currentCountry.Provinces[j] != nil {
					return currentCountry.Provinces[i].Name <= currentCountry.Provinces[j].Name
				}
				return false
			})
		}
		return
	}

	_, _, ciok := getCityFromProvince(currentProvince.Cities, cityName)
	if ciok {
		return
	}

	currentCity = &city{Code: cityCode, Name: cityName}
	currentProvince.Cities = append(currentProvince.Cities, currentCity)
	if len(currentProvince.Cities) > 1 {
		sort.SliceStable(currentProvince.Cities, func(i, j int) bool {
			if currentProvince.Cities[i] != nil && currentProvince.Cities[j] != nil {
				return currentProvince.Cities[i].Name <= currentProvince.Cities[j].Name
			}
			return false
		})
	}

}

func (loc location) String() string {

	return fmt.Sprintf("Country: %s, Province: %s, City: %s", loc.CountryName, loc.ProvinceName, loc.CityName)

}

func getCountryFromCountries(inCountries []*country, countryName string) (int, *country, bool) {
	if inCountries != nil {
		length := len(inCountries)
		i := sort.Search(length, func(i int) bool {
			if inCountries[i] != nil {
				return inCountries[i].Name >= countryName
			}
			return false
		})

		if i < length && inCountries[i].Name == countryName {
			return i, inCountries[i], true
		}
	}
	return 0, nil, false

}

func getProvinceFromCountry(inProvinces []*province, provinceName string) (int, *province, bool) {

	if inProvinces != nil {

		length := len(inProvinces)
		i := sort.Search(length, func(i int) bool {
			if inProvinces[i] != nil {
				return inProvinces[i].Name >= provinceName
			}
			return false
		})

		if i < length && inProvinces[i].Name == provinceName {
			return i, inProvinces[i], true
		}
	}
	return 0, nil, false

}

func getCityFromProvince(inCities []*city, cityName string) (int, *city, bool) {

	if inCities != nil {

		length := len(inCities)
		i := sort.Search(length, func(i int) bool {
			if inCities[i] != nil {
				return inCities[i].Name >= cityName
			}
			return false
		})

		if i < length && inCities[i].Name == cityName {
			return i, inCities[i], true
		}
	}
	return 0, nil, false
}

func getCountry(countryName, provinceName, cityName string) (country, error) {

	_, c, _ := getCountryFromCountries(countries, countryName)
	if c == nil {
		return country{}, errors.New(locationNotFound)
	}

	if provinceName == "" {
		return copyCountry(*c), nil
	}

	outCountry := country{}
	outCountry.Name = c.Name
	outCountry.Code = c.Code

	_, p, _ := getProvinceFromCountry(c.Provinces, provinceName)
	if p == nil {
		return country{}, errors.New(locationNotFound)
	}

	if cityName == "" {
		tempProvince := copyProvince(*p)
		outCountry.Provinces = append(outCountry.Provinces, &tempProvince)
		return outCountry, nil

	}
	outProvince := province{}
	outProvince.Name = p.Name
	outProvince.Code = p.Code

	_, ci, _ := getCityFromProvince(p.Cities, cityName)
	if ci == nil {
		return country{}, errors.New(locationNotFound)
	}

	outCity := city{}
	outCity.Name = ci.Name
	outCity.Code = ci.Code
	outProvince.Cities = append(outProvince.Cities, &outCity)
	outCountry.Provinces = append(outCountry.Provinces, &outProvince)
	return outCountry, nil

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
