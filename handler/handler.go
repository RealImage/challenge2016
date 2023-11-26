package handler

import (
	"awesomeProject43/models"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isSubset(child, parent []string) bool {
	set := make(map[string]bool)
	for _, value := range parent {
		set[value] = true
	}
	for _, value := range child {
		if _, ok := set[value]; !ok {
			return false
		}
	}
	return true
}

func AddDistributor(distributors *map[string]models.Permissions) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Distributor Name: ")
	scanner.Scan()
	name := scanner.Text()

	distributor := models.Distributor{
		Name:              name,
		PermittedPlaces:   make([]string, 0),
		RestrictedPlaces:  make([]string, 0),
		AuthorizedRegions: make(map[string]bool),
		Parent:            nil,
	}

	for {
		fmt.Print("Enter permission (INCLUDE/EXCLUDE): REGION  e.g., CHICAGO-ILLINOIS-UNITEDSTATES or press 4 to finish: ")
		scanner.Scan()
		permission := scanner.Text()

		if permission == "4" {
			break
		}

		data := strings.Split(permission, ":")
		if len(data) != 2 {
			fmt.Println("Invalid input. Please follow the format: INCLUDE: REGION or EXCLUDE: REGION")
			continue
		}

		action := strings.TrimSpace(data[0])
		region := strings.TrimSpace(data[1])

		switch action {
		case "INCLUDE":
			distributor.PermittedPlaces = append(distributor.PermittedPlaces, region)
		case "EXCLUDE":
			distributor.RestrictedPlaces = append(distributor.RestrictedPlaces, region)
		default:
			fmt.Println("Invalid action. Please use INCLUDE or EXCLUDE.")
		}
	}

	// Build authorized regions map
	for _, region := range distributor.PermittedPlaces {
		distributor.AuthorizedRegions[region] = true
	}
	for _, region := range distributor.RestrictedPlaces {
		distributor.AuthorizedRegions[region] = false
	}

	fmt.Print("Enter parent distributor name (leave blank if none): ")
	scanner.Scan()
	parentName := scanner.Text()
	if parentName != "" {
		if parent, ok := (*distributors)[parentName]; ok {
			// Check if child's permissions are a subset of the parent's permissions
			if !isSubset(distributor.PermittedPlaces, parent.GetPermittedPlaces()) {
				fmt.Println("Invalid permissions. Child's permissions must be a subset of the parent's permissions.")
				return
			}

			distributor.Parent = parent
			// Set child's permissions to be a subset of the parent's permissions
			distributor.PermittedPlaces = append(distributor.PermittedPlaces, parent.GetPermittedPlaces()...)
			distributor.RestrictedPlaces = append(distributor.RestrictedPlaces, parent.GetRestrictedPlaces()...)
		} else {
			fmt.Println("Parent distributor not found.")
			return
		}
	}

	(*distributors)[name] = distributor
	fmt.Println("Distributor added successfully.")
}

func CheckPermission(distributors map[string]models.Permissions) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Distributor Name: ")
	scanner.Scan()
	name := scanner.Text()

	if distributor, ok := distributors[name]; ok {
		fmt.Print("Enter query (e.g., CHICAGO-ILLINOIS-UNITEDSTATES): ")
		scanner.Scan()
		query := scanner.Text()

		if hasPermission(distributor, query) {
			fmt.Println("Permission granted.")
		} else {
			fmt.Println("Permission denied.")
		}
	} else {
		fmt.Println("Distributor not found.")
	}
}

// hasPermission checks if the distributor has permission for the given region
func hasPermission(distributor models.Permissions, region string) bool {
	// Split the region into country, state, and city
	country, state, city := parseRegion(region)
	// Check if the distributor has permission for the country
	if !contains(distributor.GetPermittedPlaces(), country) {
		return false
	}

	// Check if the distributor has restriction for the state or city
	if contains(distributor.GetRestrictedPlaces(), state+"-"+country) {
		return false
	}

	if city != "" && contains(distributor.GetRestrictedPlaces(), city+"-"+state+"-"+country) {
		return false
	}

	// Check parent's permissions
	if distributor.GetParent() != nil {
		return hasPermission(distributor.GetParent(), region)
	}

	return true
}

// contains checks if a string is present in a slice of strings
func contains(slice []string, s string) bool {
	for _, value := range slice {
		// Check if every part in the value is present in the query
		if s == value {
			return true
		}
	}
	return false
}

// parseRegion splits the region into country, state, and city
func parseRegion(region string) (country, state, city string) {
	parts := strings.Split(region, "-")
	switch len(parts) {
	case 1:
		country = parts[0]
	case 2:
		state, country = parts[0], parts[1]
	case 3:
		city, state, country = parts[0], parts[1], parts[2]
	}
	return country, state, city
}
