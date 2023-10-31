package util

import (
	"challenge2016/model"
	"strings"
)

// CheckIfStateExcluded returns true if all states are excluded or all the cities are excluded, else false
func CheckIfStateExcluded(state string, country *model.Country) bool {
	// return true if country since all the states under a country are excluded
	if len(country.States) == 0 {
		return true
	}

	// return true if all cities under a state are excluded
	if state, ok := country.States[state]; ok {
		return len(state.Cities) == 0
	}

	return false
}

// CheckIfStateIncluded returns true if all states are included or given state is existed in IncludedStates map, else false
func CheckIfStateIncluded(state string, IncludedStates map[string]*model.State) bool {
	// returns true as all the states are included
	if len(IncludedStates) == 0 {
		return true
	}

	// returns true if the given state is included in IncludedStates map
	if _, ok := IncludedStates[state]; ok {
		return true
	}

	return false
}

// CheckIfCountryExcluded returns true if country existed in excludedCountry map, else return false
func CheckIfCountryExcluded(country string, excludedCountry map[string]*model.Country) bool {

	// return false if excludedCountry is empty since its not explicilty excluded
	if len(excludedCountry) == 0 {
		return false
	}

	if country, ok := excludedCountry[country]; ok {
		return len(country.States) == 0
	}

	return false
}

// CheckIfCountryIncluded returns true if existed in includedCountry map, else returns false
func CheckIfCountryIncluded(country string, includedCountry map[string]*model.Country) bool {
	if _, ok := includedCountry[country]; ok {
		return true
	}
	return false
}

// checkCityExisted returns true if inputcity exists in cities list else false
func CheckCityExisted(inputCity string, cities []string) bool {
	for _, city := range cities {
		if city == inputCity {
			return true
		}
	}
	return false
}

func ConvertStringToLowerCase(input string) string {
	return strings.ToLower(strings.ReplaceAll(input, " ", ""))
}
