package distributor

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

var distributors map[string]Distributor

var distributorsLocations map[string]*ValidDistributionLocations

var filepath string

type Distributor struct {
	ParentDistributor string     `json:"parentDistributor"`
	Name              string     `json:"distributorName"`
	IncludedLocations []Location `json:"includedLocations"`
	ExcludedLocations []Location `json:"excludedLocations"`
}

type Location struct {
	CityCode    string `json:"cityCode"`
	StateCode   string `json:"stateCode"`
	CountryCode string `json:"countryCode"`
}

type ValidDistributionLocations struct {
	Name              string
	ParentDistributor string
	Locations         CountryMap
}

type DistributorPermissions struct {
	Name     string     `json:"distributorName"`
	Location []Location `json:"location"`
}

func init() {
	distributors = make(map[string]Distributor)
	distributorsLocations = make(map[string]*ValidDistributionLocations)
	flag.StringVar(&filepath, "file", "cities.csv", "CSV File Path")
}

func CreateDistributor(w http.ResponseWriter, req *http.Request) {
	distributorsData := []Distributor{}
	err := json.NewDecoder(req.Body).Decode(&distributorsData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, distributor := range distributorsData {
		// Checking distributor is present or not
		distributorPresent, errMsg := checkDistributorExists(distributor)
		if distributorPresent {
			http.Error(w, "ERROR: "+errMsg, http.StatusInternalServerError)
		}

		// Checking distributor's parent is valid or not
		parentPresent, errMsg := validateDistributorParent(distributor)
		if !parentPresent {
			http.Error(w, "ERROR: "+errMsg, http.StatusInternalServerError)
		}

		distLocations := &ValidDistributionLocations{
			Name:              distributor.Name,
			ParentDistributor: distributor.ParentDistributor,
		}
		// Get all valid locations for a distributor
		distLocations.getDistributorValidLocations(distributor)
		distributorsLocations[distLocations.Name] = distLocations

		if errMsg == "" {
			distributors[distributor.Name] = distributor
		}
	}
	resp, _ := json.Marshal(distributors)
	w.Write(resp)
}

func VerifyDistributorRights(w http.ResponseWriter, req *http.Request) {
	distributorsData := []DistributorPermissions{}
	err := json.NewDecoder(req.Body).Decode(&distributorsData)
	if err != nil {
		fmt.Println(err)
	}
	locationRights := make(map[string]map[string]string)

	for _, distributor := range distributorsData {
		locationData := make(map[string]string)
		distributorData, isDistributorAvailable := distributorsLocations[distributor.Name]
		if isDistributorAvailable {
			for _, location := range distributor.Location {
				countryCode := location.CountryCode
				stateCode := location.StateCode
				cityCode := location.CityCode
				isValid := false

				// Adding location data for response
				locationName := cityCode + "-" + stateCode + "-" + countryCode
				if cityCode == "" && stateCode == "" && countryCode == "" {
					errMsg := "Distributor " + distributor.Name + " doesn't have a valid location."
					http.Error(w, "ERROR: "+errMsg, http.StatusInternalServerError)
				} else if cityCode == "" && stateCode == "" {
					locationName = countryCode
				} else if cityCode == "" {
					locationName = stateCode + "-" + countryCode
				}
				locationData[locationName] = "NO"

				_, isCountryAvailable := distributorData.Locations[countryCode]
				if isCountryAvailable {
					_, isStateAvailable := distributorData.Locations[countryCode][stateCode]
					if isStateAvailable || stateCode == "" {
						_, isCityAvailable := distributorData.Locations[countryCode][stateCode][cityCode]
						if isCityAvailable || cityCode == "" {
							isValid = true
						}
					}
				}

				if isValid == true {
					locationData[locationName] = "YES"
				}
				locationRights[distributor.Name] = locationData
			}
		} else {
			errMsg := "Distributor " + distributor.Name + " not available"
			http.Error(w, "ERROR: "+errMsg, http.StatusInternalServerError)
		}
	}
	resp, _ := json.Marshal(locationRights)
	w.Write(resp)
}
