package internal

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type CityID string
type StateID string
type CountryID string

type CityName string
type StateName string
type CountryName string

type RegionData struct {
	country CountryName
	state   StateName
	city    CityName
}

// regionDB is a map of country to state to city
var regionDB map[CountryID]map[StateID]map[CityID]RegionData

// NewRegionDB creates a new region database
func NewRegionDB() map[CountryID]map[StateID]map[CityID]RegionData {
	if regionDB == nil {
		regionDB = make(map[CountryID]map[StateID]map[CityID]RegionData)
	}
	return regionDB
}

// AddRegion adds a new region to the database
func AddRegion(cityID, stateID, countryID, cityName, stateName, countryName string) error {
	if countryID == "" || stateID == "" || cityID == "" || countryName == "" || stateName == "" || cityName == "" {
		//return fmt.Errorf("invalid region data: CountryID, StateID, CityID, CountryName, StateName, CityName " +
		//	"cannot be empty")
		fmt.Println("Region data is invalid as some fields are empty. Ignoring region data:", cityID, stateID,
			countryID, cityName, stateName, countryName)
		return nil
	}
	if _, ok := regionDB[CountryID(countryID)]; !ok {
		regionDB[CountryID(countryID)] = make(map[StateID]map[CityID]RegionData)
	}
	if _, ok := regionDB[CountryID(countryID)][StateID(stateID)]; !ok {
		regionDB[CountryID(countryID)][StateID(stateID)] = make(map[CityID]RegionData)
	}
	if _, ok := regionDB[CountryID(countryID)][StateID(stateID)][CityID(cityID)]; !ok {
		regionDB[CountryID(countryID)][StateID(stateID)][CityID(cityID)] = RegionData{
			country: CountryName(countryName),
			state:   StateName(stateName),
			city:    CityName(cityName),
		}
	} else {
		fmt.Println("Ignoring duplicate region data:", cityID, stateID, countryID, cityName, stateName, countryName)
		// return fmt.Errorf("region already exists", cityID, stateID, countryID, cityName, stateName, countryName)
	}
	return nil
}

// IsValidRegion checks if the regionID is valid
func IsValidRegion(regionID string) bool {
	// Empty regionID is invalid
	validRegion := false
	if regionID == "" {
		return false
	}
	regionsSplit := strings.Split(regionID, "-")
	switch len(regionsSplit) {
	case 1:
		//country
		if isValidCountry(regionsSplit[0]) {
			validRegion = true
		}
	case 2:
		//state and country
		if isValidCountry(regionsSplit[1]) && isValidState(regionsSplit[0], regionDB[CountryID(regionsSplit[1])]) {
			validRegion = true
		}
	case 3:
		//city, state and country
		if isValidCountry(regionsSplit[2]) && isValidState(regionsSplit[1], regionDB[CountryID(regionsSplit[2])]) &&
			isValidCity(regionsSplit[0], regionDB[CountryID(regionsSplit[2])][StateID(regionsSplit[1])]) {
			validRegion = true
		}
		validRegion = true
	default:
		//invalid
		return validRegion
	}
	return validRegion
}

// IsValidCountry checks if the country is valid
func isValidCountry(id string) bool {
	if id == "" {
		return false
	}
	if _, ok := regionDB[CountryID(id)]; ok {
		return true
	}
	return false
}

// IsValidState checks if the state is valid
func isValidState(id string, stateDB map[StateID]map[CityID]RegionData) bool {
	if id == "" {
		return false
	}
	if _, ok := stateDB[StateID(id)]; ok {
		return true
	}
	return false
}

// IsValidCity checks if the city is valid
func isValidCity(id string, cityDB map[CityID]RegionData) bool {
	if id == "" {
		return false
	}
	if _, ok := cityDB[CityID(id)]; ok {
		return true
	}
	return false
}

// ReadRegionDataFromRemoteCSV reads region data from a remote CSV file
func ReadRegionDataFromRemoteCSV(url string) error {
	if url == "" {
		return fmt.Errorf("URL cannot be empty")
	}
	NewRegionDB()
	url = "https://raw.githubusercontent.com/RealImage/challenge2016/master/cities.csv"
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching CSV file:", err)
	}
	defer response.Body.Close()

	reader := csv.NewReader(response.Body)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading CSV:", err)
	}
	for i, record := range records {
		if len(record) == 6 && i != 0 {
			err := AddRegion(record[0], record[1], record[2], record[3], record[4], record[5])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// ReadRegionDataFromLocalCSV reads region data from a local CSV file
func ReadRegionDataFromLocalCSV(path string) error {
	//if path == "" {
	//	return fmt.Errorf("file path cannot be empty")
	//}
	NewRegionDB()
	path = "/Users/arb.khan/go/src/github/film-distribution-management/inventory/cities.csv"
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error opening file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading CSV:", err)
	}
	for i, record := range records {
		if len(record) == 6 && i != 0 {
			err := AddRegion(record[0], record[1], record[2], record[3], record[4], record[5])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
