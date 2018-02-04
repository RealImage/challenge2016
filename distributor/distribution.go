package distribution

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Cities struct {
	AreaCode    string `json:"area_code"`
	StateCode   string `json:"state_code"`
	CountryCode string `json:"country_code"`
	Area        string `json:"area"`
	State       string `json:"state"`
	Country     string `json:"country`
}

func PrepareCitiesJson(fileName string) []Cities {
	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var cities []Cities
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		cities = append(cities, Cities{
			AreaCode:    line[0],
			StateCode:   line[1],
			CountryCode: line[2],
			Area:        line[3],
			State:       line[4],
			Country:     line[5],
		})
	}
	return cities
}

func getCountryNames(cities []Cities) []string {
	//collect all the Country names and return as an array
	var countries []string
	for _, item := range cities {
		countries = append(countries, item.Country)
	}

	return removeDuplicates(countries)
}

func excludeAreaName(countries []string, cities []Cities, excludeState []string) []string {
	var allowedAreas []string
	var allAreas []string

	/*If a distributor has access to a country, ultimately he is subjected to have access too all the cities corresponding to the Country*/
	for _, country := range countries {
		for _, item := range cities {
			if item.Country == country {
				allAreas = append(allAreas, fmt.Sprintf("%v_%v_%v", item.Area, item.State, item.Country))
			}
		}
	}

	/*If the distributed has excluded state then sub user should definitely not have authority to access corresponding cities*/
	if len(excludeState) > 0 {
		for _, need := range allAreas {
			for _, item := range excludeState {
				state := strings.Split(need, "_")
				if state[1] != item {
					allowedAreas = append(allowedAreas, need)
				}
			}
		}
	} else {
		return removeDuplicates(allAreas)
	}

	return removeDuplicates(allowedAreas)
}

func excludeStateName(countries []string, cities []Cities) []string {
	/*having access to a country is enough for a ditributor to exclude corresponding states*/
	var allowedState []string
	for _, country := range countries {
		for _, item := range cities {
			if item.Country == country {
				allowedState = append(allowedState, fmt.Sprintf("%v_%v", item.State, item.Country))
			}
		}
	}

	return removeDuplicates(allowedState)
}

func extendStateAccess(cities []Cities, countries []string) []string {
	/*Direct distributor can choose any state provided he haven't included/excluded corresponding country*/
	var unselectedState []string

	for _, item := range cities {
		unselectedState = append(unselectedState, fmt.Sprintf("%v_%v", item.State, item.Country))

	}

	for _, country := range countries {
		for _, item := range cities {
			if item.Country == country {
				unselectedState = Remove(unselectedState, fmt.Sprintf("%v_%v", item.State, item.Country))
			}
		}
	}

	return removeDuplicates(unselectedState)
}

func extendAreaAccess(cities []Cities, countries []string, states []string) []string {
	/*Direct distributor can choose any city provided he haven't included/excluded corresponding country/state*/
	var unSelectedCity []string

	for _, item := range cities {
		unSelectedCity = append(unSelectedCity, fmt.Sprintf("%v_%v_%v", item.Area, item.State, item.Country))
	}
	/*Area access should not be permitted in case if the user is prohibited from any State or Country*/
	/*if len(countries) > 0 && len(states) > 0 {
		for _, country := range countries {
			for _, item := range cities {
				if item.Country == country {
					for _, state := range states {
						stateName := fmt.Sprintf("%v_%v", item.State, item.Country)
						if stateName == state {
							unSelectedCity = Remove(unSelectedCity, fmt.Sprintf("%v_%v_%v", item.Area, item.State, item.Country))
						}
					}
				}
			}
		}
	}*/

	if len(countries) > 0 {
		for _, item := range cities {
			for _, country := range countries {
				if item.Country == country {
					unSelectedCity = Remove(unSelectedCity, fmt.Sprintf("%v_%v_%v", item.Area, item.State, item.Country))
				}
			}
		}
	}

	if len(states) > 0 {
		for _, item := range cities {
			for _, state := range states {
				stateName := fmt.Sprintf("%v_%v", item.State, item.Country)
				if stateName == state {
					unSelectedCity = Remove(unSelectedCity, fmt.Sprintf("%v_%v_%v", item.Area, item.State, item.Country))
				}
			}
		}
	}

	return removeDuplicates(unSelectedCity)
}

func subAllowedStateNames(userAccess map[string]interface{}, cities []Cities) []string {
	/*Depending on the distributor access state will be prepared*/
	countryAccess, _ := userAccess["countries"].([]string)
	includeState, _ := userAccess["included_states"].([]string)
	excludeState, _ := userAccess["excluded_states"].([]string)

	var allowedStates []string

	for _, cnt := range countryAccess {
		for _, item := range cities {
			if item.Country == cnt {
				allowedStates = append(allowedStates, fmt.Sprintf("%v_%v", item.State, item.Country))
			}
		}
	}

	for _, state := range includeState {
		allowedStates = append(allowedStates, state)
	}

	if len(excludeState) > 0 {
		for _, allowed := range allowedStates {
			for _, restricted := range excludeState {
				if allowed == restricted {
					allowedStates = Remove(allowedStates, restricted)
				}
			}
		}
	}

	return removeDuplicates(allowedStates)
}

func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] != true {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

/*TODO: Optimize this function*/
func subAllowedAreaNames(userAccess map[string]interface{}, cities []Cities, include []string, exclude []string) []string {
	countryAccess, _ := userAccess["countries"].([]string)
	includeArea, _ := userAccess["included_cities"].([]string)
	includeState, _ := userAccess["included_states"].([]string)
	excludeState, _ := userAccess["excluded_states"].([]string)
	excludeArea, _ := userAccess["excluded_cities"].([]string)

	var allowedAreas []string

	if len(include) > 0 {
		includeState = append(includeState, include...)
	}

	if len(exclude) > 0 {
		excludeState = append(excludeState, exclude...)
	}

	/*prepare a list with all the City names*/
	for _, cnt := range countryAccess {
		for _, item := range cities {
			if item.Country == cnt {
				allowedAreas = append(allowedAreas, fmt.Sprintf("%v_%v_%v", item.Area, item.State, item.Country))
			}
		}
	}

	/*based on the parent distributor remove the city name*/

	for _, inc := range includeArea {
		allowedAreas = append(allowedAreas, inc)
	}

	for _, ins := range includeState {
		for _, item := range cities {
			custom := strings.Split(ins, "_")
			if item.State == custom[0] {
				allowedAreas = append(allowedAreas, fmt.Sprintf("%v_%v_%v", item.Area, item.State, item.Country))
			}
		}
	}

	for _, exs := range excludeState {
		for _, item := range cities {
			custom := strings.Split(exs, "_")
			if item.State == custom[0] {
				allowedAreas = Remove(allowedAreas, fmt.Sprintf("%v_%v_%v", item.Area, item.State, item.Country))
			}
		}
	}

	for _, exa := range excludeArea {
		for _, item := range allowedAreas {
			if item == exa {
				allowedAreas = Remove(allowedAreas, item)
			}
		}
	}

	return removeDuplicates(allowedAreas)
}

func Remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func PrepareRootUser(input []string, cities []Cities) map[string]interface{} {
	currentUser := make(map[string]interface{})
	var countries []string
	var excludeStateAccess []string
	var excludeAreaAccess []string
	var rootStateAccess []string
	var rootAreaAccess []string

	for _, line := range input {
		lowerLine := strings.ToLower(line)
		if strings.Contains(lowerLine, "exclude") && strings.Contains(lowerLine, ":") {
			custom := strings.Split(line, ":")
			custom1 := custom[1]
			if strings.Contains(custom1, "_") {
				detect := strings.Split(custom1, "_")
				if len(detect) == 2 {
					allowedStates := excludeStateName(countries, cities)
					stateName := ExistInArray(allowedStates, custom1)
					if stateName != "" {
						excludeStateAccess = append(excludeStateAccess, stateName)
					} else {
						currentUser["err"] = fmt.Sprintf("[%v] Not permitted, please try again\n", line)
						return currentUser
					}
				} else if len(detect) == 3 {
					allowedAreas := excludeAreaName(countries, cities, excludeStateAccess)
					areaName := ExistInArray(allowedAreas, custom1)
					if areaName != "" {
						excludeAreaAccess = append(excludeAreaAccess, areaName)
					} else {
						currentUser["err"] = fmt.Sprintf("[%v] Not permitted, please try again\n", line)
						return currentUser
					}
				}
			} else {
				currentUser["err"] = fmt.Sprintf("[%v] Not permitted to exclude Country, please try again\n", line)
			}
		} else if strings.Contains(lowerLine, "include") && strings.Contains(lowerLine, ":") {
			allCountries := getCountryNames(cities)
			custom := strings.Split(line, ":")
			custom1 := custom[1]
			if strings.Contains(custom1, "_") {
				detect := strings.Split(custom1, "_")
				if len(detect) == 2 {
					allowedStates := extendStateAccess(cities, countries)
					stateName := ExistInArray(allowedStates, custom1)
					if stateName != "" {
						rootStateAccess = append(rootStateAccess, stateName)
					} else {
						currentUser["err"] = fmt.Sprintf("[%v] Not permitted, please try again\n", line)
						return currentUser
					}
				} else if len(detect) == 3 {
					allowedAreas := extendAreaAccess(cities, countries, excludeStateAccess)
					areaName := ExistInArray(allowedAreas, custom1)
					if areaName != "" {
						rootAreaAccess = append(rootAreaAccess, areaName)
					} else {
						currentUser["err"] = fmt.Sprintf("[%v] Not permitted, please try again\n", line)
						return currentUser
					}
				}
			} else {
				countryName := ExistInArray(allCountries, custom1)
				if countryName != "" {
					countries = append(countries, countryName)
				} else {
					currentUser["err"] = fmt.Sprintf("[%v] Not permitted, please try again\n", line)
					return currentUser
				}
			}
		} else {
			currentUser["err"] = fmt.Sprintf("[%v] Not permitted - mention the INCLUDE/EXCLUDE operation\n", line)
		}
	}

	currentUser["countries"] = removeDuplicates(countries)
	currentUser["excluded_states"] = removeDuplicates(excludeStateAccess)
	currentUser["excluded_cities"] = removeDuplicates(excludeAreaAccess)
	currentUser["included_states"] = removeDuplicates(rootStateAccess)
	currentUser["included_cities"] = removeDuplicates(rootAreaAccess)

	return currentUser
}

func PrepareSubUser(input []string, cities []Cities, root map[string]interface{}) map[string]interface{} {
	currentUser := make(map[string]interface{})
	var excludeStateAccess []string
	var excludeAreaAccess []string
	var includedStateAccess []string
	var excludedStateAccess []string

	for _, line := range input {
		lowerLine := strings.ToLower(line)
		if strings.Contains(lowerLine, "exclude") && strings.Contains(lowerLine, ":") {
			custom := strings.Split(line, ":")
			custom1 := custom[1]
			//custom1, _ := strings.TrimSpace(tmp)
			if strings.Contains(custom1, "_") {
				detect := strings.Split(custom1, "_")
				if len(detect) == 2 {
					allowedStates := subAllowedStateNames(root, cities)
					stateName := ExistInArray(allowedStates, custom1)
					if stateName != "" {
						excludeStateAccess = append(excludeStateAccess, stateName)
					} else {
						currentUser["err"] = fmt.Sprintf("[%v] Not permitted, please try again\n", line)
						return currentUser
					}
				} else if len(detect) == 3 {
					allowedAreas := subAllowedAreaNames(root, cities, includedStateAccess, excludedStateAccess)
					areaName := ExistInArray(allowedAreas, custom1)
					if areaName != "" {
						excludeAreaAccess = append(excludeAreaAccess, areaName)
					} else {
						currentUser["err"] = fmt.Sprintf("[%v] Not permitted, please try again\n", line)
						return currentUser
					}
				}
			} else {
				currentUser["err"] = fmt.Sprintf("[%v] Not permitted to exclude Country, please try again\n", line)
			}
		} else if strings.Contains(lowerLine, "include") && strings.Contains(lowerLine, ":") {
			custom := strings.Split(line, ":")
			custom1 := custom[1]
			if strings.Contains(custom1, "_") {
				detect := strings.Split(custom1, "_")
				if len(detect) == 2 {
					allowedStates := subAllowedStateNames(root, cities)
					stateName := ExistInArray(allowedStates, custom1)
					if stateName != "" {
						includedStateAccess = append(includedStateAccess, stateName)
					} else {
						currentUser["err"] = fmt.Sprintf("[%v] Not permitted, please try again\n", line)
						return currentUser
					}
				} else if len(detect) == 3 {
					allowedAreas := subAllowedAreaNames(root, cities, includedStateAccess, excludedStateAccess)
					areaName := ExistInArray(allowedAreas, custom1)
					if areaName != "" {
						excludedStateAccess = append(excludedStateAccess, areaName)
					} else {
						currentUser["err"] = fmt.Sprintf("[%v] Not permitted, please try again\n", line)
						return currentUser
					}
				}
			} else {
				currentUser["err"] = fmt.Sprintf("[%v] Not permitted to include Country, please try again\n", line)
			}
		} else {
			currentUser["err"] = fmt.Sprintf("[%v] Not permitted - mention the INCLUDE/EXCLUDE operation\n", line)
		}
	}

	currentUser["excluded_states"] = removeDuplicates(excludeStateAccess)
	currentUser["excluded_cities"] = removeDuplicates(excludeAreaAccess)
	currentUser["included_states"] = removeDuplicates(includedStateAccess)
	currentUser["included_cities"] = removeDuplicates(excludedStateAccess)
	return currentUser
}

func ExistInArray(listOfItems []string, name string) string {
	for _, item := range listOfItems {
		if strings.EqualFold(strings.Replace(item, " ", "", -1), strings.Replace(name, " ", "", -1)) {
			return item
		}
	}
	return ""
}
