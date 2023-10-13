package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// To store DistributionInformation which contains name of the dsitributor, city ,state and country they wanted to distribute
type DistributorInformation struct {
	distributorName string
	city            string
	state           string
	country         string
}

// main is the entry point of the application.
func main() {
	// Load permission data from the "permission.csv" file.
	authorizations, err := LoadPermissionData("permission.csv")
	if err != nil {
		fmt.Println("Not able to load permission data, Please check the CSV File")
		os.Exit(0)
	}
	//If Args are less then expected then exit the application
	if len(os.Args) != 3 {
		fmt.Println("Not enough arguments to proceed")
		os.Exit(0)
	}
	// Retrieve distributor name and location from command line arguments.
	inputDistributorName := os.Args[1]
	inputLocation := os.Args[2]

	// Split the input location into city, state, and country.
	city, state, country := splitLocation(inputLocation)

	// Create a DistributorInformation object based on the input data.
	distributorInput := DistributorInformation{
		distributorName: inputDistributorName,
		city:            city,
		state:           state,
		country:         country,
	}

	// Retrieve the distributor information from the loaded authorizations.
	distributorName := authorizations[inputDistributorName]

	// Initialize a flag to check if there is a match.
	var IsMatched bool

Distributor:
	// Iterate through the "INCLUDE" section of distributor information.
	for _, data := range distributorName["INCLUDE"] {
		// Check if the input distributor's country matches the data's country.
		includeMatch := findMatch(data, distributorInput)
		if includeMatch {
			IsMatched = true
			break
		}
	}

	// Iterate through the "EXCLUDE" section of distributor information.
	for _, data := range distributorName["EXCLUDE"] {
		excludeMatch := findMatch(data, distributorInput)
		// If any one of them is excluded then exit the loop
		if excludeMatch {
			IsMatched = false
			break
		}
	}

	// Iterate through other distributor keys in the authorizations (excluding "INCLUDE" and "EXCLUDE").
	for key := range distributorName {
		if key != "INCLUDE" && key != "EXCLUDE" {
			// Update distributorName and jump to the "Distributor" label for further checks.
			distributorName = authorizations[key]
			goto Distributor
		}
	}

	// Determine and print the final permission status.
	if IsMatched {
		fmt.Println("Permission: YES")
	} else {
		fmt.Println("Permission: NO")
	}
}

func findMatch(data, distributorInput DistributorInformation) bool {

	var isMatched bool
	// Initialize a variable to determine which fields to consider for matching.
	var LenghtToConsider int
	// Determine the length to consider based on available data (country, state, city).
	if data.city != "" && data.state != "" {
		LenghtToConsider = 3
	} else if data.state != "" {
		LenghtToConsider = 2
	} else {
		LenghtToConsider = 1
	}
	// Check for exclusions based on the length to consider.
	if LenghtToConsider == 3 { //country state city
		isMatched = (distributorInput.city == data.city && distributorInput.state == data.state && distributorInput.country == data.country)
	} else if LenghtToConsider == 2 { //state city
		isMatched = (distributorInput.state == data.state && distributorInput.country == data.country)
	} else { //city
		isMatched = distributorInput.country == data.country
	}
	return isMatched
}

// splitLocation takes a string representing a location in the format "city-state-country"
// and splits it into three components: city, state, and country. If any of these components
// is missing in the input string, the corresponding variable will be an empty string.
// The function returns these three components as separate strings.
func splitLocation(location string) (string, string, string) {
	// Split the input string using the "-" delimiter and store the parts in an array.
	parts := strings.Split(location, "-")
	// Initialize variables to store the city, state, and country components.
	city, state, country := "", "", ""
	// Check the number of parts obtained from the split operation.
	if len(parts) == 3 {
		// If there are three parts, assign them to the city, state, and country variables.
		city, state, country = parts[0], parts[1], parts[2]
	} else if len(parts) == 2 {
		// If there are two parts, assign them to the state and country variables.
		state, country = parts[0], parts[1]
	} else {
		// If there is only one part, assign it to the country variable.
		country = parts[0]
	}
	// Return the city, state, and country components as separate strings.
	return city, state, country
}

// LoadPermissionData loads permission data from a CSV file located at the specified filePath.
// It parses the CSV file, extracts distributor information, and organizes it into a hierarchical
// map structure based on distributor names, distributor levels, and location.
// The function returns a map containing the parsed data, organized by distributor hierarchy,
// and an error if any issues occur during file reading or parsing.
func LoadPermissionData(filePath string) (map[string]map[string][]DistributorInformation, error) {
	// Open the CSV file at the provided filePath.
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a CSV reader to read records from the file.
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Initialize a map to store distributor information in a hierarchical structure.
	distributionMap := make(map[string]map[string][]DistributorInformation)

	// Loop through the CSV records and process distributor information.
	for _, record := range records {
		var dp DistributorInformation

		// Skip the header row.
		if record[0] == "Distributor" {
			continue
		}

		// Split the location information into city, state, and country.
		city, state, country := splitLocation(record[2])

		// Split the distributor record by "<" to identify distributor levels.
		distributorRecord := strings.Split(record[0], "<")

		if len(distributorRecord) == 3 {
			distributionMap[distributorRecord[0]] = make(map[string][]DistributorInformation)
			distributionMap[distributorRecord[0]][distributorRecord[1]] = nil
			distributionMap[distributorRecord[0]][distributorRecord[2]] = nil
		} else if len(distributorRecord) == 2 {
			distributionMap[distributorRecord[0]] = make(map[string][]DistributorInformation)
			distributionMap[distributorRecord[0]][distributorRecord[1]] = nil
		}

		// Set the distributor name and location information in the distributor record.
		record[0] = distributorRecord[0]
		dp.city = city
		dp.state = state
		dp.country = country

		if _, ok := distributionMap[record[0]]; ok {
			if _, ok := distributionMap[record[0]][record[1]]; ok {
				distributionMap[record[0]][record[1]] = append(distributionMap[record[0]][record[1]], dp)
			} else {
				distributionMap[record[0]][record[1]] = []DistributorInformation{dp}
			}
		} else {
			distributionMap[record[0]] = make(map[string][]DistributorInformation)
			distributionMap[record[0]][record[1]] = []DistributorInformation{dp}
		}
	}

	return distributionMap, nil
}
