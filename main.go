package main

import (
	"challenge2016/dto"        // Importing DTO package for data transfer objects
	"challenge2016/input"      // Importing input package for user input handling
	"challenge2016/parser"     // Importing parser package for parsing CSV files
	"challenge2016/permission" // Importing permission package for permission checking
	"challenge2016/validator"  // Importing validator package for data validation
	"fmt"
	"log"
)

func main() {
	var distributorInformation []dto.Distributor
	groupedData, err := parser.ParseCSVFile("cities.csv") // Parsing the CSV file containing city data
	if err != nil {
		log.Fatalf("Error parsing CSV file: %v", err) // Exiting if there's an error in parsing CSV
	}

	for {
		choice := input.AskNextQuestion() // Asking the user for their choice
		switch choice {
		case "1":
			distributorData := input.GetDistributorData()                                                       // Getting data for a new distributor
			errorRes := validator.ValidateDistributorData(distributorData, groupedData, distributorInformation) // Validating distributor data
			if len(errorRes) > 0 {
				fmt.Println(errorRes)
				continue
			}
			distributorObject := input.CreateNewDistributor(distributorData)           // Creating a new distributor object
			distributorInformation = append(distributorInformation, distributorObject) // Appending the new distributor to the list
		case "2":
			subDistributorData := input.GetSubDistributorData()                                                       // Getting data for a new sub-distributor
			errorRes := validator.ValidateSubDistributorData(subDistributorData, groupedData, distributorInformation) // Validating sub-distributor data
			if len(errorRes) > 0 {
				fmt.Println(errorRes)
				continue
			}
			subDistributorObject := input.CreateNewDistributor(subDistributorData)        // Creating a new sub-distributor object
			distributorInformation = append(distributorInformation, subDistributorObject) // Appending the new sub-distributor to the list
		case "3":
			checkPermissionData := input.GetCheckPermissionData()                                                       // Getting data to check permission
			errorRes := validator.ValidateCheckPermissionData(checkPermissionData, groupedData, distributorInformation) // Validating permission check data
			if len(errorRes) > 0 {
				fmt.Println(errorRes)
				continue
			}
			checkPermissionResult := permission.CheckPermission(checkPermissionData.DistributorName, checkPermissionData.Regions, "Check Permission", distributorInformation) // Checking permission
			fmt.Println("Check Permission Result:", checkPermissionResult)
		case "4":
			input.DisplayDistributorInformation(distributorInformation) // Displaying distributor information
		case "5":
			fmt.Println("Exiting the program")
			return // Exiting the program
		}
	}
}
