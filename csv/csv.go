package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Location struct {
	CityCode     string
	ProvinceCode string
	CountryCode  string
	CityName     string
	ProvinceName string
	CountryName  string
}

type LocationSlice []*Location

type Csv struct {
	FileName string `json:"file_name"`
}

func (c *Csv) Read() (map[string]*Location, error) {

	file, err := os.Open(c.FileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var locations LocationSlice
	for _, record := range records[1:] {
		location := &Location{
			CityCode:     record[0],
			ProvinceCode: record[1],
			CountryCode:  record[2],
			CityName:     record[3],
			ProvinceName: record[4],
			CountryName:  record[5],
		}
		locations = append(locations, location)
	}
	return locations.Parse(), nil
}

func (l LocationSlice) Parse() map[string]*Location {

	regions := make(map[string]*Location)

	for i, line := range l {
		if i == 0 { // Skip header
			continue
		}
		key := fmt.Sprintf("%s-%s-%s", line.CountryCode, line.ProvinceCode, line.CityCode)
		regions[key] = line
	}
	return regions
}
