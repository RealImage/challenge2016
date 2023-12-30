package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// City
type City struct {
	Code         string
	Province     string
	Country      string
	Name         string
	ProvinceCode string
	CountryCode  string
}

// Distributor Permission
type Distributor struct {
	name    string
	include []string // to be in format either of IN or IN-KA or IN-KA-BANGL
	exclude []string // to be in format either of IN or IN-KA or IN-KA-BANGL
}

var CityData []City

// ReadCities reads city data from a CSV file.
func ReadCities(filePath string) ([]City, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cities []City
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) == 6 {
			city := City{
				Code:         fields[0],
				ProvinceCode: fields[1],
				CountryCode:  fields[2],
				Name:         fields[3],
				Province:     fields[4],
				Country:      fields[5],
			}
			cities = append(cities, city)
		}
	}

	return cities, nil
}

const Format = "(comma separated and in format of country or country-state or country-state-city,  e.g., IN or IN-UP or IN-UP-ZMANI )"

// create permission
func createDistributor() {
	var distributor Distributor
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter distributor name: ")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())
	fmt.Print("Enter included regions " + Format + " : ")
	scanner.Scan()
	includedRegions := strings.Split(strings.TrimSpace(scanner.Text()), ",")
	//validateRegions()
	fmt.Print("Enter excluded regions " + Format + " : ")
	scanner.Scan()
	excludedRegions := strings.Split(strings.TrimSpace(scanner.Text()), ",")
	//validateRegions()
	distributor.name = name
	distributor.include = includedRegions
	distributor.exclude = excludedRegions
	saveDistributorToCSV(distributor)
}

func saveDistributorToCSV(distributor Distributor) {
	// Open the CSV file in append mode or create it if it doesn't exist
	file, err := os.OpenFile("distributors.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)

	// // Write the header to the CSV file
	// header := []string{"Name", "Include", "Exclude"}
	// if err := writer.Write(header); err != nil {
	// 	fmt.Println("Error writing header:", err)
	// 	return
	// }

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	if fileInfo.Size() == 0 {
		// File is empty, write the header
		header := []string{"Name", "Include", "Exclude"}
		if err := writer.Write(header); err != nil {
			fmt.Println("Error writing header:", err)
			return
		}
	}

	// Convert include and exclude slices to a comma-separated string
	includeStr := strings.Join(distributor.include, ",")
	excludeStr := strings.Join(distributor.exclude, ",")

	// Write the distributor data to the CSV file
	row := []string{distributor.name, includeStr, excludeStr}
	if err := writer.Write(row); err != nil {
		fmt.Println("Error writing row:", err)
		return
	}
	// Flush and close the CSV writer
	writer.Flush()

	// Check for errors during the writing process
	if err := writer.Error(); err != nil {
		fmt.Println("Error writing CSV:", err)
		return
	}
}

func fetchAllDistributorsFromCSV(filePath string) ([]Distributor, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	header, err := reader.Read()
	if err != nil {
		return nil, err
	}

	expectedHeader := []string{"Name", "Include", "Exclude"}
	if !areEqual(header, expectedHeader) {
		return nil, fmt.Errorf("Invalid CSV format. Expected header: %v", expectedHeader)
	}

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var distributors []Distributor
	for _, record := range records {
		distributor := Distributor{
			name:    record[0],
			include: strings.Split(record[1], ","),
			exclude: strings.Split(record[2], ","),
		}
		distributors = append(distributors, distributor)
	}

	return distributors, nil
}

func areEqual(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
}

func showDistributorList() ([]string, error) {
	var distributorList []string
	distributors, err := fetchAllDistributorsFromCSV("distributors.csv")
	if err != nil {
		return nil, err
	}
	for _, distributor := range distributors {
		distributorList = append(distributorList, distributor.name)
	}
	return distributorList, nil
}

func checkAuthorization(distributorName string, regionToCheck string) (bool, error) {
	distributor, err := fetchPermissionForADistributorFromCSV(distributorName)
	if err != nil {
		return false, err
	}
	permissions := addRemainingPermissionForADistributor(distributor)

	for _, item := range permissions {
		if startsWith(item, regionToCheck) {
			return true, nil
		}
	}
	return false, nil

}

func fetchPermissionForADistributorFromCSV(distributorName string) (Distributor, error) {
	distributors, err := fetchAllDistributorsFromCSV("distributors.csv")
	if err != nil {
		return Distributor{}, err
	}
	for _, distributor := range distributors {
		if distributor.name == distributorName {
			return distributor, nil // TODO: Currently there will only be a single entry for a distributor
		}
	}
	return Distributor{}, nil
}

func addRemainingPermissionForADistributor(distributor Distributor) []string {
	var permissibleRegions []string
	for _, region := range distributor.include {
		countryCode, stateCode, city := getRegions(region)

		if countryCode != "" && stateCode != "" && city != "" {
			permissibleRegions = append(permissibleRegions, countryCode+"-"+stateCode+"-"+city)
		} else if countryCode != "" && stateCode != "" {
			var allCities = []string{}
			allCities = fetchCities(stateCode)
			for _, element := range allCities {
				permissibleRegions = append(permissibleRegions, countryCode+"-"+stateCode+"-"+element)
			}
		} else if countryCode != "" {
			var allCities = []string{}
			allStates := fetchStates(countryCode)
			for eachState := range allStates {
				allCities = fetchCities(allStates[eachState])
				for _, element := range allCities {
					permissibleRegions = append(permissibleRegions, countryCode+"-"+allStates[eachState]+"-"+element)
				}
			}

		}

	}

	permissibleRegions = removeExcludedPermission(distributor, permissibleRegions)

	return permissibleRegions

}

func getRegions(input string) (string, string, string) {
	parts := strings.Split(input, "-")

	switch len(parts) {
	case 1:
		return parts[0], "", ""
	case 2:
		return parts[0], parts[1], ""
	case 3:
		return parts[0], parts[1], parts[2]
	default:
		return "", "", ""
	}
}

func removeExcludedPermission(distributor Distributor, permissibleRegions []string) []string {
	for _, region := range distributor.exclude {
		permissibleRegions = removeLocationPermission(permissibleRegions, region)
	}
	return permissibleRegions
}

func removeLocationPermission(permissibleRegions []string, locationPrefix string) []string {
	var result []string
	for _, s := range permissibleRegions {
		if !startsWith(s, locationPrefix) {
			result = append(result, s)
		}
	}
	return result
}

func startsWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func fetchAllCountries() []string {
	// Create a map to store unique CountryCodes
	uniqueCodes := make(map[string]bool)

	// Iterate through the slice of City
	for _, city := range CityData {
		uniqueCodes[city.CountryCode] = true
	}

	// Extract unique CountryCodes from the map
	var result []string
	for code := range uniqueCodes {
		result = append(result, code)
	}

	return result
}

func fetchCities(stateCode string) []string {
	uniqueCodes := make(map[string]bool)

	for _, city := range CityData {
		if city.ProvinceCode == stateCode {
			uniqueCodes[city.Code] = true
		}
	}

	result := make([]string, 0, len(uniqueCodes))
	for code := range uniqueCodes {
		result = append(result, code)
	}

	return result
}

func fetchStates(countryCode string) []string {
	provinceCodes := make(map[string]struct{})

	for _, city := range CityData {
		if city.CountryCode == countryCode {
			provinceCodes[city.ProvinceCode] = struct{}{}
		}
	}

	uniqueCodes := make([]string, 0, len(provinceCodes))
	for code := range provinceCodes {
		uniqueCodes = append(uniqueCodes, code)
	}

	return uniqueCodes
}

// isCountry checks whether entered region is a country
func isCountry(region string) bool {
	regionPart := strings.Split(region, "-")
	if len(regionPart) > 1 {
		return false
	}
	allCountries := fetchAllCountries()
	for _, element := range allCountries {
		if element == regionPart[0] {
			return true
		}
	}
	return false
}

func initialise() {
	cities, err := ReadCities("cities.csv")
	if err != nil {
		fmt.Println("Error reading cities CSV:", err)
		return
	}
	CityData = cities
}

func addDistributorPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Do you want to add a distributor ? (Y/N)")
	scanner.Scan()
	addDistributor := strings.TrimSpace(scanner.Text())
	if addDistributor == "Y" {
		for {
			createDistributor()
			fmt.Print("Do you want to add more ? (Y/N)")
			scanner.Scan()
			addMore := strings.TrimSpace(scanner.Text())
			if addMore == "N" {
				break
			}
		}
	}
	return
}

func checkAuthorizationPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Do you want to check region authorization for a distributor? (Y/N)")
	scanner.Scan()
	checkAuth := strings.TrimSpace(scanner.Text())
	if checkAuth == "Y" {
		for {
			distributorsList, err := showDistributorList()
			if err != nil {
				fmt.Print(err.Error())
				return
			}
			fmt.Println(strings.Join(distributorsList, "\n"))
			fmt.Printf("Enter distributor : ")
			scanner.Scan()
			distributor := strings.TrimSpace(scanner.Text())
			fmt.Printf("Enter regions "+Format+" to check for distributor %s: ", distributor)
			scanner.Scan()
			regionToCheck := strings.TrimSpace(scanner.Text())

			isAllowed, err := checkAuthorization(distributor, regionToCheck)
			if err != nil {
				fmt.Print(err.Error())
				return
			}
			if isAllowed {
				fmt.Println("ALLOWED")
			} else {
				fmt.Println("NOT ALLOWED")
			}
			fmt.Print("Do you want to check more ? (Y/N)")
			scanner.Scan()
			checkMore := strings.TrimSpace(scanner.Text())
			if checkMore == "N" {
				break
			}
		}
	}
}

func main() {

	initialise()
	addDistributorPrompt()
	checkAuthorizationPrompt()

	// isAllowed, _ = checkAuthorization("Distributor2", "IN")
	// fmt.Println(isAllowed)
	// isAllowed, _ = checkAuthorization("Distributor2", "IN-WB")
	// fmt.Println(isAllowed)
	// isAllowed, _ = checkAuthorization("Distributor2", "IN-UP-ZMANI")
	// fmt.Println(isAllowed)
	// isAllowed, _ = checkAuthorization("Distributor2", "IN-WB-TUFGA")
	// fmt.Println(isAllowed)
}
