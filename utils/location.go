package utils

import (
	"encoding/csv"
	"os"

	"github.com/saurabh-sde/challenge2016_saurabh/model"
)

var (
	CityMap     map[string]model.Location
	ProvinceMap map[string]model.Location
	CountryMap  map[string]model.Location
)

func LoadCities(fileName string) ([]model.Location, error) {
	var locations []model.Location
	f, err := os.Open(fileName)
	if err != nil {
		Error("Error opening file: ", err)
		return nil, err
	}
	defer f.Close()

	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		Error(err)
		return nil, err
	}

	// load code and locations
	CityMap = make(map[string]model.Location)
	ProvinceMap = make(map[string]model.Location)
	CountryMap = make(map[string]model.Location)
	for i, row := range rows {
		if i == 0 {
			continue
		}
		// add location
		location := model.Location{
			CityCode:     row[0],
			ProvinceCode: row[1],
			CountryCode:  row[2],
			CityName:     row[3],
			ProvinceName: row[4],
			CountryName:  row[5],
		}
		CityMap[row[0]] = location
		ProvinceMap[row[1]] = location
		CountryMap[row[2]] = location
		locations = append(locations, location)
	}
	return locations, nil
}
