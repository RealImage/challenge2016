package main

import (
	"challenge2016/dto"        // Importing DTO package for data transfer objects
	"challenge2016/input"      // Importing input package for user input handling
	"challenge2016/parser"     // Importing parser package for parsing CSV files
	"challenge2016/permission" // Importing permission package for permission checking
	"challenge2016/validator"  // Importing validator package for data validation
	"fmt"
	"log"
	"strings"
)

func main() {
	var distributorInformation []dto.Distributor
	groupedData, err := parser.ParseCSVFile("cities.csv") // Parsing the CSV file containing city data
	if err != nil {
		log.Fatalf("Error parsing CSV file: %v", err) // Exiting if there's an error in parsing CSV
	}

	for {
		choice := input.PromptMenu() // Asking the user for their choice
		switch choice {
		case "Create a new distributor":
			distributorData := input.PromptDistributorData(false)                                                      // Getting data for a new distributor and send false for this the distributor Data
			errorRes := validator.ValidateDistributorData(distributorData, groupedData, distributorInformation, false) // Validating distributor data
			if len(errorRes) > 0 {
				fmt.Println(strings.Join(errorRes, "\n"))
				continue
			}
			distributorInformation = append(distributorInformation, distributorData) // Appending the new distributor to the list
		case "Create a sub distributor":
			subDistributorData := input.PromptDistributorData(true)                                                      // Getting data for a new sub-distributor and send true for this the sub-distributor Data
			errorRes := validator.ValidateDistributorData(subDistributorData, groupedData, distributorInformation, true) // Validating sub-distributor data
			if len(errorRes) > 0 {
				fmt.Println(strings.Join(errorRes, "\n"))
				continue
			}
			distributorInformation = append(distributorInformation, subDistributorData) // Appending the new sub-distributor to the list
		case "Check permission for a distributor":
			checkPermissionData := input.PromptCheckPermissionData()                                                    // Getting data to check permission
			errorRes := validator.ValidateCheckPermissionData(checkPermissionData, groupedData, distributorInformation) // Validating permission check data
			if len(errorRes) > 0 {
				fmt.Println(strings.Join(errorRes, "\n"))
				continue
			}
			checkPermissionResult := permission.CheckPermission(checkPermissionData.DistributorName, checkPermissionData.Regions, "Check Permission", distributorInformation) // Checking permission
			fmt.Println("Check Permission Result:", checkPermissionResult)
		case "View Distributors information":
			for _, distributor := range distributorInformation {
				fmt.Printf("Name: %s, Include: %v, Exclude: %v, Parent: %s\n", distributor.Name, distributor.Include, distributor.Exclude, distributor.Parent)
			}
		case "Exit":
			fmt.Println("Exiting the program")
			return // Exiting the program
		}
	}
}
