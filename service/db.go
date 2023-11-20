package service

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

var dataSet map[string][]string

// LoadDataset - loading the dataset into memory
func LoadDataset() (err error) {
	dataSet = make(map[string][]string)

	file, err := os.Open(CITIES_DB)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		return err
	}

	for count, record := range records {
		if count > 0 {
			key := record[5]
			value := strings.Join(record, "-")
			dataSet[key] = append(dataSet[key], value)
		}
	}

	return err
}

// ValidateRegion - func to check if the regions mentioned are available or not
func ValidateRegion(regions []string) bool {
	count := 0
	for _, region := range regions {
		substr := strings.Split(region, "-")
		record := dataSet[getCountry(substr)]
		if record != nil {
			for _, data := range record {
				if strings.Contains(data, region) {
					count++
					break
				}
			}
		}
	}
	return count == len(regions)
}

// getCountry - function to extract country because it is better to search this way
func getCountry(region []string) (country string) {
	lastIndex := 0
	if len(region) > 1 {
		lastIndex = len(region) - 1
	}
	country = region[lastIndex]
	return country
}
