package distributor

type CountryMap map[string]StateMap

type StateMap map[string]CityMap

type CityMap map[string]bool

var locations CountryMap

// Get all valid locations for a distributor
func (distLocations *ValidDistributionLocations) getDistributorValidLocations(distributor Distributor) {
	distLocations.Locations = make(CountryMap)
	locations, _ = loadCSVData()

	for _, includedLocation := range distributor.IncludedLocations {
		countryCode := includedLocation.CountryCode
		stateCode := includedLocation.StateCode
		cityCode := includedLocation.CityCode
		distLocations.addCountryData(countryCode, stateCode, cityCode)
	}

	for _, excludedLocation := range distributor.ExcludedLocations {
		countryCode := excludedLocation.CountryCode
		stateCode := excludedLocation.StateCode
		cityCode := excludedLocation.CityCode
		distLocations.deleteLocationData(countryCode, stateCode, cityCode)
	}

	if distributor.ParentDistributor != "none" && distributor.ParentDistributor != "" {
		distLocations.deleteParentsExcludedLocations(distributor)
	}
}

func (distLocations *ValidDistributionLocations) deleteParentsExcludedLocations(distributor Distributor) {
	if distributor.ParentDistributor != "none" && distributor.ParentDistributor != "" {
		parentDistributor := distributors[distributor.ParentDistributor]
		for _, excludedLocation := range parentDistributor.ExcludedLocations {
			countryCode := excludedLocation.CountryCode
			stateCode := excludedLocation.StateCode
			cityCode := excludedLocation.CityCode
			distLocations.deleteLocationData(countryCode, stateCode, cityCode)
		}
		if parentDistributor.ParentDistributor != "none" && parentDistributor.ParentDistributor != "" {
			distLocations.deleteParentsExcludedLocations(parentDistributor)
		}
	}
}

func (distLocations *ValidDistributionLocations) addCountryData(countryCode, stateCode, cityCode string) {
	if countryCode == "*" || countryCode == "" {
		distLocations.Locations = locations
	} else if countryCode != "" {
		_, isCountryAvailable := locations[countryCode]
		if isCountryAvailable {
			_, isCountryInitialized := distLocations.Locations[countryCode]
			if !isCountryInitialized {
				distLocations.Locations[countryCode] = make(StateMap)
			}
			distLocations.addStateData(countryCode, stateCode, cityCode)
		}
	}
}

func (distLocations *ValidDistributionLocations) addStateData(countryCode, stateCode, cityCode string) {
	if stateCode == "*" || stateCode == "" {
		distLocations.Locations[countryCode] = locations[countryCode]
	} else if stateCode != "" {
		_, isStateAvailable := locations[countryCode][stateCode]
		if isStateAvailable {
			_, isStateInitialized := distLocations.Locations[countryCode][stateCode]
			if !isStateInitialized {
				distLocations.Locations[countryCode][stateCode] = make(CityMap)
			}
			distLocations.addCityData(countryCode, stateCode, cityCode)
		}
	}
}

func (distLocations *ValidDistributionLocations) addCityData(countryCode, stateCode, cityCode string) {
	if cityCode == "*" || cityCode == "" {
		distLocations.Locations[countryCode][stateCode] = locations[countryCode][stateCode]
	} else if cityCode != "" {
		_, isCityAvailable := locations[countryCode][stateCode][cityCode]
		if isCityAvailable {
			distLocations.Locations[countryCode][stateCode][cityCode] = true
		}
	}
}

func (distLocations *ValidDistributionLocations) deleteLocationData(countryCode, stateCode, cityCode string) {
	if countryCode == "*" || countryCode == "" {
		distLocations.Locations = make(CountryMap)
	}
	_, isCountryAvailable := distLocations.Locations[countryCode]
	if isCountryAvailable {
		if stateCode == "*" || stateCode == "" {
			delete(distLocations.Locations, countryCode)
		}
		_, isStateAvailable := distLocations.Locations[countryCode][stateCode]
		if isStateAvailable {
			if cityCode == "*" || cityCode == "" {
				delete(distLocations.Locations[countryCode], stateCode)
			}
			_, isCityAvailable := distLocations.Locations[countryCode][stateCode][cityCode]
			if isCityAvailable {
				delete(distLocations.Locations[countryCode][stateCode], cityCode)
			}
		}
	}
	return
}
