package main

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"
)

type DistributorNetwork struct {
	cityData map[string]map[string]map[string]struct{}
	distributorData map[string]Distributor
}

type Distributor struct {
	Parent string `json:"parent"`
	Included map[string]struct{} `json:"included"`
	Excluded map[string]struct{} `json:"excluded"`
}

/* 
LoadCities will load and make the data set to be ready with data loaded from csv
*/
func (dn *DistributorNetwork) LoadCities(){
	
	tmpCityData := make(map[string]map[string]map[string]struct{})
	
	// load countries and assign the data
	file, err := os.Open("./cities.csv")
	if err != nil {
		LogError(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	
	firstRow := true
	for {
		row, err := csvReader.Read()
		
		if err == io.EOF {
			break
		}

		if err != nil {
			LogError(err)
		}
		
		if firstRow {
			firstRow = false
			continue
		}

		if _, countryExists := tmpCityData[string(row[2])]; !countryExists {
			tmpCityData[string(row[2])] = map[string]map[string]struct{} {
				string(row[1]) : {
					string(row[0]) : {},
				},
			}
			continue
		}
		
		if _, proviceExists := tmpCityData[string(row[2])][string(row[1])]; !proviceExists {
			tmpCityData[string(row[2])][string(row[1])] = map[string]struct{} {
				string(row[0]) : {},
			}
			continue
		}

		tmpCityData[string(row[2])][string(row[1])][string(row[0])] = struct{}{}
	}
	
	dn.cityData = tmpCityData
}

/* 
AddDistributor callback excute the procedures and validations to add a new distributor to the data set
*/
func (dn *DistributorNetwork) AddDistributor(distID string, parentDistID string, includedRegions []string, excludedRegions []string) error {
	if distID == "" {
		return errors.New("Distributor id is mandatory")
	}

	if _, exists := dn.distributorData[distID]; exists {
		return errors.New("Distributor already exists")
	}
	
	if parentDistID != "" {
		if _, exists := dn.distributorData[parentDistID]; !exists {
			return errors.New("Parent distributor does not exist")
		}
	}

	tmpIncluded := make(map[string]struct{}, 0)

	for _, includedItem := range includedRegions {

		var isValidRegion, country, provice, city = dn.getValidRegion(includedItem) 

		if !isValidRegion {
			return errors.New("Invalid region code")
		}

		for validItem := range tmpIncluded {
			if validItem == includedItem {
				return errors.New("Duplicate region is not permitted")
			}
			if strings.HasSuffix(validItem, includedItem){
				return errors.New("Region overlap is not permitted")
			}
			if strings.HasSuffix(includedItem, validItem){
				return errors.New("Region overlap is not permitted")
			}
		} 

		if parentDistID != "" {
			isValid := dn.validateParent(parentDistID, country, provice, city)
			if !isValid {
				return errors.New("One of the provided included regions is not a subset of parent distributor")
			}
		}

		tmpIncluded[includedItem] = struct{}{}
	}

	tmpExcluded := make(map[string]struct{}, 0)

	for _, excludedItem := range excludedRegions {

		var isValidRegion, country, provice, city = dn.getValidRegion(excludedItem) 

		if !isValidRegion {
			return errors.New("Invalid region code")
		}

		isInclSubset := false
		
		if _, exists := tmpIncluded[dn.makeRegionSlug(country)]; exists {
			isInclSubset = true
		}
		
		if _, exists := tmpIncluded[dn.makeRegionSlug(provice, country)]; exists {
			isInclSubset = true
		}
		
		if _, exists := tmpIncluded[dn.makeRegionSlug(city, provice, country)]; exists {
			isInclSubset = false
		}
		
		if !isInclSubset {
			return errors.New("One of the provided excluded regions is not a subset of the included regions")
		}

		for validItem := range tmpExcluded {
			if validItem == excludedItem {
				return errors.New("Duplicate region is not permitted")
			}
			if strings.HasSuffix(validItem, excludedItem){
				return errors.New("Region overlap is not permitted")
			}
			if strings.HasSuffix(excludedItem, validItem){
				return errors.New("Region overlap is not permitted")
			}
		}

		tmpExcluded[excludedItem] = struct{}{}
	}

	dn.distributorData[distID] = Distributor{Included: tmpIncluded, Excluded: tmpExcluded, Parent: parentDistID}

	return nil
}

/* 
IsDistributorPermitted will validate the distributor access to the provided region
*/
func (dn *DistributorNetwork) IsDistributorPermitted(distID string, region string) (bool, error) {

	if _, exists := dn.distributorData[distID]; !exists {
		return false, errors.New("Distributor not exists")
	}
	
	var isValidRegion, country, provice, _ = dn.getValidRegion(region) 
	
	if !isValidRegion {
		return false, errors.New("Invalid region code")
	}
	
	distributor := dn.distributorData[distID]

	if len(distributor.Excluded) > 0 {
		if _, exists := distributor.Excluded[region]; exists {
			return false, nil
		}
	
		if country != "" {
			if _, exists := distributor.Excluded[country]; exists {
				return false, nil
			}
		}
		
		if country != "" && provice != "" {
			if _, exists := distributor.Excluded[dn.makeRegionSlug(provice, country)]; exists {
				return false, nil
			}
		}
	}

	allExcluded := dn.getAllExcluded(distributor.Parent)

	if len(allExcluded) > 0 {
		if _, exists := allExcluded[region]; exists {
			return false, nil
		}
	
		if country != "" {
			if _, exists := allExcluded[country]; exists {
				return false, nil
			}
		}
		
		if country != "" && provice != "" {
			if _, exists := allExcluded[dn.makeRegionSlug(provice, country)]; exists {
				return false, nil
			}
		}
	}

	if len(distributor.Included) > 0 {
		if country != "" {
			if _, exists := distributor.Included[country]; exists {
				return true, nil
			}
		}
		
		if country != "" && provice != "" {
			if _, exists := distributor.Included[dn.makeRegionSlug(provice, country)]; exists {
				return true, nil
			}
		}
		
		if _, exists := distributor.Included[region]; exists {
			return true, nil
		}
	}
	
	return false, nil
}

/* 
getAllExcluded is a sort of recursive function which will return the excluded regions of self and upline parents'
*/
func (dn *DistributorNetwork) getAllExcluded(distID string) (map[string]struct{}) {
	distributor := dn.distributorData[distID]
	excludedRegions := dn.distributorData[distID].Excluded
	if distributor.Parent != "" {
		parentExcluded := dn.getAllExcluded(distributor.Parent)
		for key := range parentExcluded {
			excludedRegions[key] = struct{}{}
		}
	}
	return excludedRegions
}

/* 
getValidRegion will check the provide region information against available data set
*/
func (dn *DistributorNetwork) getValidRegion(regionSlug string) (isValid bool, country string, province string, city string) {
	
	if len(regionSlug) == 0 {
		return false, "", "", ""
	}
	
	regionSplit := strings.Split(regionSlug, "-")
	
	if len(regionSplit) > 3 {
		return false, "", "", ""
	}
	
	switch len(regionSplit) {
	case 3:
		// country-provice-city
		if _, exists := dn.cityData[regionSplit[2]][regionSplit[1]][regionSplit[0]]; !exists {
			return false, "", "", ""
		}
		return true, regionSplit[2], regionSplit[1], regionSplit[0]
	case 2:
		// country-provice
		if _, exists := dn.cityData[regionSplit[1]][regionSplit[0]]; !exists {
			return false, "", "", ""
		}
		return true, regionSplit[1], regionSplit[0], ""
	case 1:
		// country
		if _, exists := dn.cityData[regionSplit[0]]; !exists {
			return false, "", "", ""
		}
		return true, regionSplit[0], "", ""
	}
		
	return false, "", "", ""
}

/* 
makeRegionSlug will make the region slug with default seperator(hyphen)
*/
func (dn *DistributorNetwork) makeRegionSlug(data ...string) string {
	return strings.Join(data, "-") 
}

/* 
validateParent will validate the given region code of a distributor against its parent's regions
*/
func (dn *DistributorNetwork) validateParent(parentDistID string, country string, province string, city string) bool {

	if len(dn.distributorData[parentDistID].Included) == 0 {
		return false
	}

	if country != "" {
		if _, exists := dn.distributorData[parentDistID].Excluded[country]; exists {
			return false
		}
	}
	
	if country != "" && province != "" {
		if _, exists := dn.distributorData[parentDistID].Excluded[dn.makeRegionSlug(province,country)]; exists {
			return false
		}
	}
	
	if country != "" && province != "" && city != "" {
		if _, exists := dn.distributorData[parentDistID].Excluded[dn.makeRegionSlug(city,province,country)]; exists {
			return false
		}
	}

	if country != "" {
		if _, exists := dn.distributorData[parentDistID].Included[country]; exists {
			return true
		}
	}

	if country != "" && province != "" {
		if _, exists := dn.distributorData[parentDistID].Included[dn.makeRegionSlug(province,country)]; exists {
			return true
		}
	}

	if country != "" && province != "" && city != "" {
		if _, exists := dn.distributorData[parentDistID].Included[dn.makeRegionSlug(city,province,country)]; exists {
			return true
		}
	}

	return false
}

/* 
NewDistributorNetwork will create a DistributorNetwork struct and initialize the required data
*/
func NewDistributorNetwork() *DistributorNetwork {

	distributorNetworkObj := new(DistributorNetwork)

	distributorNetworkObj.LoadCities()

	distributorNetworkObj.distributorData = make(map[string]Distributor, 0)

	return distributorNetworkObj
}