package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type City struct {
	Name        string
	StateCode   string
	CountryCode string
}

type Permission struct {
	IncludedCountries []string
	ExcludedStates    map[string][]string
}

func main() {
	data := make(map[string]map[string][]string)

	csvfile, err := os.Open("cities.csv")
	if err != nil {
		panic(err)
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.Read() // skip header row

	for {
		row, err := reader.Read()
		if err != nil {
			break
		}
		countryCode := row[0]
		stateCode := row[1]
		city := row[2]
		if _, ok := data[countryCode]; !ok {
			data[countryCode] = make(map[string][]string)
		}
		if _, ok := data[countryCode][stateCode]; !ok {
			data[countryCode][stateCode] = make([]string, 0)
		}
		data[countryCode][stateCode] = append(data[countryCode][stateCode], city)
	}

	permissions := readPermissions()

	city := readCity()

	if isCityAllowed(city, permissions, data) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

// Reads the distributor permissions from standard input.
func readPermissions() Permission {
	var permissions Permission
	permissions.ExcludedStates = make(map[string][]string)

	fmt.Print("Enter distributor permissions: ")
	for {
		var line string
		_, err := fmt.Scanln(&line)
		if err != nil {
			break
		}
		if strings.HasPrefix(line, "INCLUDE:") {
			countryCode := strings.TrimSpace(line[8:])
			permissions.IncludedCountries = append(permissions.IncludedCountries, countryCode)
		} else if strings.HasPrefix(line, "EXCLUDE:") {
			parts := strings.Split(strings.TrimSpace(line[8:]), "-")
			stateCode := parts[0]
			countryCode := parts[1]
			permissions.ExcludedStates[countryCode] = append(permissions.ExcludedStates[countryCode], stateCode)
		}
	}

	return permissions
}

// Reads the city from standard input and returns it as a City struct.
func readCity() City {
	var city City
	fmt.Print("Enter city: ")
	fmt.Scanln(&city.Name, &city.StateCode, &city.CountryCode)
	return city
}

// Returns true if the given city is allowed according to the given permissions and city data.
func isCityAllowed(city City, permissions Permission, data map[string]map[string][]string) bool {
	if !contains(permissions.IncludedCountries, city.CountryCode) {
		return false
	}
	if excludedStates, ok := permissions.ExcludedStates[city.CountryCode]; ok && contains(excludedStates, city.StateCode) {
		return false
	}
	return contains(data[city.CountryCode][city.StateCode], city.Name)
}

// Returns true if the given slice contains the given string.
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
