package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

// CityData represents the data structure for city details.
type CityData struct {
	Cities []City `json:"cities"`
}

// City represents details of a city.
type City struct {
	CityCode     string `json:"City Code"`
	ProvinceCode string `json:"Province Code"`
	CountryCode  string `json:"Country Code"`
	CityName     string `json:"City Name"`
	ProvinceName string `json:"Province Name"`
	CountryName  string `json:"Country Name"`
}

// DistributorData represents the data structure for distributor access.
type DistributorData struct {
	Distributors map[string]Distributor `json:"distributors"`
}

// Distributor represents a distributor's access to regions.
type Distributor struct {
	Include  []string `json:"include"`
	Exclude  []string `json:"exclude"`
	Children []string `json:"children"`
}

func main() {
	// Define command-line flags
	inputFlag := flag.String("input", "cities.json", "JSON file containing city data")
	distributorFlag := flag.String("distributor", "", "Distributor name")
	includeFlag := flag.String("include", "", "Regions to include (comma-separated)")
	excludeFlag := flag.String("exclude", "", "Regions to exclude (comma-separated)")
	fromFlag := flag.String("from", "", "Distributor to delegate permissions from")
	toFlag := flag.String("to", "", "Distributor to delegate permissions to")
	accessFlag := flag.Bool("access", false, "Display distributor access")
	regionFlag := flag.String("region", "", "Display distributors for a specific region")
	deleteFlag := flag.String("delete", "", "Delete distributor and its children")

	// Parse command-line arguments
	flag.Parse()

	var cityData []City
	var distributorData *DistributorData
	var err error

	// Load city data from the input JSON file
	cityData, err = loadCityDataFromFile(*inputFlag)
	if err != nil {
		fmt.Println("Error loading city data:", err)
		os.Exit(1)
	}

	// Load or create distributor data
	distributorData = NewDistributorData() // Initialize DistributorData
	err = loadOrCreateDistributorData(distributorData)
	if err != nil {
		fmt.Println("Error loading distributor data:", err)
		os.Exit(1)
	}

	// Perform actions based on command-line flags
	if *distributorFlag != "" {
		if *includeFlag != "" {
			// Add distributor with access
			distributorData.AddDistributor(*distributorFlag, strings.Split(*includeFlag, ","), nil)
		} else if *excludeFlag != "" {
			// Remove access for distributor
			excludedCities := distributorData.ExcludeCities(strings.Split(*excludeFlag, ","), cityData)
			distributorData.RemoveAccess(*distributorFlag, excludedCities)
		}
	}

	if *fromFlag != "" && *toFlag != "" && *includeFlag != "" {
		// Delegate permissions from one distributor to another
		distributorData.DelegatePermissions(*fromFlag, *toFlag, strings.Split(*includeFlag, ","))
	}

	if *accessFlag {
		// Display distributor access
		distributorData.DisplayDistributorAccess(*distributorFlag)
	}

	if *regionFlag != "" {
		// Display distributors for a specific region
		distributorData.DisplayDistributorsForRegion(*regionFlag)
	}

	if *deleteFlag != "" {
		// Delete a distributor and its children
		distributorData.DeleteDistributor(*deleteFlag)
	}

	// Save updated distributor data to the internal JSON file
	err = saveDistributorData(distributorData)
	if err != nil {
		fmt.Println("Error saving distributor data:", err)
	}
}

// loadCityDataFromFile loads city data from a JSON file.
func loadCityDataFromFile(filename string) ([]City, error) {
	data := make([]City, 0)

	// Load data from the JSON file
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data as a slice of City objects
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewDistributorData() *DistributorData {
	return &DistributorData{
		Distributors: make(map[string]Distributor),
	}
}

// loadOrCreateDistributorData loads distributor data from the internal JSON file
// or creates a new one if it doesn't exist.
func loadOrCreateDistributorData(data *DistributorData) error {
	// Check if the internal JSON file exists
	if _, err := os.Stat("internal.json"); os.IsNotExist(err) {
		// Create a new distributor with access to all regions
		data.AddDistributor("default", []string{"*"}, nil)
	} else {
		// Load data from the internal JSON file
		fileData, err := os.ReadFile("internal.json")
		if err != nil {
			return err
		}

		// Unmarshal the JSON data as DistributorData
		err = json.Unmarshal(fileData, data)
		if err != nil {
			return err
		}
	}

	return nil
}

// AddDistributor adds a distributor with access to specific regions.
func (data *DistributorData) AddDistributor(name string, include []string, children []string) {
	distributor := Distributor{
		Include:  include,
		Exclude:  nil,
		Children: children,
	}
	data.Distributors[name] = distributor
}

// RemoveAccess removes access to specific regions for a distributor.
func (data *DistributorData) RemoveAccess(name string, exclude []string) {
	distributor, exists := data.Distributors[name]
	if !exists {
		return
	}
	distributor.Exclude = append(distributor.Exclude, exclude...)
	data.Distributors[name] = distributor
}

// ExcludeCities excludes all cities in the specified regions.
func (data *DistributorData) ExcludeCities(regions []string, cityData []City) []string {
	var excludedCities []string
	for _, city := range cityData {
		for _, region := range regions {
			if city.CountryCode == region || city.ProvinceCode == region || city.CityCode == region {
				excludedCities = append(excludedCities, city.CityCode)
				break
			}
		}
	}
	return excludedCities
}

// DelegatePermissions delegates permissions from one distributor to another for specific regions.
func (data *DistributorData) DelegatePermissions(fromName, toName string, include []string) {
	from, fromExists := data.Distributors[fromName]
	to, toExists := data.Distributors[toName]

	if !fromExists || !toExists {
		fmt.Println("Invalid distributor names")
		return
	}

	// Verify that 'from' has permissions for 'include'
	if !data.CanDistribute(from.Include, from.Exclude, include) {
		fmt.Println("Invalid permissions for delegation")
		return
	}

	// Delegate permissions to 'to'
	to.Include = append(to.Include, include...)
	data.Distributors[toName] = to
}

// CanDistribute checks if a distributor can distribute to specific regions.
func (data *DistributorData) CanDistribute(include, exclude, regions []string) bool {
	// Implement the logic to check if a distributor can distribute to regions
	// based on their inclusion and exclusion lists.
	// Return true if they can distribute, otherwise false.
	for _, region := range regions {
		// Check if the region is excluded
		for _, ex := range exclude {
			if region == ex {
				return false
			}
		}
		// Check if the region is included
		if len(include) == 0 || include[0] == "*" {
			return true // No specific inclusions mean access to all regions
		}
		for _, inc := range include {
			if region == inc {
				return true
			}
		}
	}
	return false
}

// DisplayDistributorAccess displays the access regions for a distributor.
func (data *DistributorData) DisplayDistributorAccess(name string) {
	distributor, exists := data.Distributors[name]
	if !exists {
		fmt.Println("Distributor not found")
		return
	}

	fmt.Printf("%s has access to: %s\n", name, strings.Join(distributor.Include, ", "))
}

// DisplayDistributorsForRegion displays the distributors that have access to a specific region.
func (data *DistributorData) DisplayDistributorsForRegion(region string) {
	var distributors []string
	for name, distributor := range data.Distributors {
		if data.CanDistribute(distributor.Include, distributor.Exclude, []string{region}) {
			distributors = append(distributors, name)
		}
	}

	if len(distributors) == 0 {
		fmt.Printf("No distributor has access to %s\n", region)
	} else {
		fmt.Printf("%s is accessible by the following distributors: %s\n", region, strings.Join(distributors, ", "))
	}
}

// DeleteDistributor deletes a distributor and its children from the system.
func (data *DistributorData) DeleteDistributor(name string) {
	// Check if the distributor exists
	_, exists := data.Distributors[name]
	if !exists {
		fmt.Println("Distributor not found")
		return
	}

	// Delete the distributor and its children
	delete(data.Distributors, name)
}

// saveDistributorData saves distributor data to the internal JSON file.
func saveDistributorData(data *DistributorData) error {
	// Marshal data to JSON format
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Save data to the internal JSON file
	err = os.WriteFile("internal.json", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
