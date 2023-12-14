package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// Permission represents the include and exclude permissions for a distributor
type Permission struct {
	Include []string
	Exclude []string
}

// Distributor represents a distributor with its permissions
type Distributor struct {
	Name       string
	Permission Permission
	Child      *Distributor // Pointer to a child distributor
}

// Function to check if a distributor has permission for a given region
func hasPermission(distributor *Distributor, region string) bool {
	// Check for include permissions
	for _, include := range distributor.Permission.Include {
		if strings.HasPrefix(region, include) {
			return true
		}
	}

	// Check for exclude permissions
	for _, exclude := range distributor.Permission.Exclude {
		if strings.HasPrefix(region, exclude) {
			return false
		}
	}

	// Check permissions of child distributor recursively
	if distributor.Child != nil {
		return hasPermission(distributor.Child, region)
	}

	return false
}

func main() {
	// Read CSV file containing cities data
	file, err := os.Open("cities.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	// Parse CSV data
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV records:", err)
		return
	}

	// Build the permissions hierarchy
	rootDistributor := &Distributor{}
	currentDistributor := rootDistributor

	for _, record := range records {
		distributorName := record[0]
		include := record[1]
		exclude := record[2]

		newDistributor := &Distributor{
			Name: distributorName,
			Permission: Permission{
				Include: strings.Split(include, ","),
				Exclude: strings.Split(exclude, ","),
			},
		}

		if currentDistributor.Name == "" {
			rootDistributor = newDistributor
		} else {
			currentDistributor.Child = newDistributor
		}

		currentDistributor = newDistributor
	}

	// Example usage
	fmt.Println("DISTRIBUTOR1 has permission for CHICAGO-ILLINOIS-UNITEDSTATES:",
		hasPermission(rootDistributor, "CHICAGO-ILLINOIS-UNITEDSTATES"))

	fmt.Println("DISTRIBUTOR1 has permission for CHENNAI-TAMILNADU-INDIA:",
		hasPermission(rootDistributor, "CHENNAI-TAMILNADU-INDIA"))

	fmt.Println("DISTRIBUTOR1 has permission for BANGALORE-KARNATAKA-INDIA:",
		hasPermission(rootDistributor, "BANGALORE-KARNATAKA-INDIA"))
}
