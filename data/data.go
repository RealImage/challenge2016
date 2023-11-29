package data

import (
	"Qcube/models"
	"encoding/csv"
	"log"
	"os"
	"strings"
)

var Locations models.Location

func Load_data() error {
	Locations = make(models.Location)
	log.Println("Loading data")
	file, err := os.Open("data/cities.csv")

	if err != nil {
		log.Println("Error opening file", "cities.csv")
		log.Println(err)
		return err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Println("Error while reading file ", "cities.csv")
		log.Println(err)
		return err
	}

	for _, record := range records[1:] {
		country := strings.ToLower(strings.TrimSpace(string(record[5])))
		stateInCSV := strings.ReplaceAll(strings.ToLower(strings.TrimSpace(string(record[4]))), "-", "+")
		cityInCSV := strings.ReplaceAll(strings.ToLower(strings.TrimSpace(string(record[3]))), "-", "+")
		province, countryexist := Locations[country]
		log.Println("consuming", country, stateInCSV, cityInCSV)
		if countryexist {
			log.Println("Country exist", country)
			statefound := false
			for _, v := range province {
				cities, stateExist := v[strings.ToLower(stateInCSV)]
				if stateExist {
					statefound = true
					log.Println("State exist", stateInCSV)
					v[strings.ToLower(stateInCSV)] = append(cities, strings.ToLower(cityInCSV))

				}
			}
			if !statefound {
				log.Println("State does not exist", stateInCSV)
				state := make(map[string]models.Cities)
				state[strings.ToLower(stateInCSV)] = models.Cities{strings.ToLower(cityInCSV)}
				Locations[country] = append(province, state)
			}

		} else {
			log.Println("Country does not exist", country)
			state := make(map[string]models.Cities)
			state[strings.ToLower(stateInCSV)] = models.Cities{strings.ToLower(cityInCSV)}
			Locations[country] = []models.Province{state}
		}

	}

	log.Println("Data loaded", Locations)
	return nil
}
