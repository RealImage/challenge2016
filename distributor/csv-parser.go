package distributor

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

func loadCSVData() (CountryMap, error) {
	csvFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	locations := make(CountryMap)
	for {
		csvData, err := reader.Read()
		if err == io.EOF {
			// To break the unending loop
			break
		} else if err != nil {
			return nil, err
		}
		cityCode := csvData[0]
		stateCode := csvData[1]
		countryCode := csvData[2]
		_, isCountryPresent := locations[countryCode]
		if !isCountryPresent {
			locations[countryCode] = make(StateMap)
		}

		_, isStatePresent := locations[countryCode][stateCode]
		if !isStatePresent {
			locations[countryCode][stateCode] = make(CityMap)
		}

		locations[countryCode][stateCode][cityCode] = false
	}
	return locations, nil
}
