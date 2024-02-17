package main

import (
	"challenge2016/dto"           // Importing DTO package for data transfer objects
	"challenge2016/input"         // Importing input package for user input handling
	Parser "challenge2016/parser" // Importing parser package for parsing CSV files
	"challenge2016/permission"    // Importing permission package for permission checking
	"challenge2016/validator"     // Importing validator package for data validation
	"fmt"
	"log"
	"strings"
)

func main() {
	var distributorInformation []dto.Distributor
	groupedData, err := Parser.ParseCSVFile("cities.csv") // Parsing the CSV file containing city data
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
			distributorInformation = append(distributorInformation, dto.Distributor{
				Name:    strings.ToUpper(distributorData.Name),
				Include: distributorData.Include,
				Exclude: distributorData.Exclude,
				Parent:  strings.ToUpper(distributorData.Parent),
			}) // Appending the new distributor to the list
		case "2":
			subDistributorData := input.GetSubDistributorData()                                                       // Getting data for a new sub-distributor
			errorRes := validator.ValidateSubDistributorData(subDistributorData, groupedData, distributorInformation) // Validating sub-distributor data
			if len(errorRes) > 0 {
				fmt.Println(errorRes)
				continue
			}
			distributorInformation = append(distributorInformation, dto.Distributor{
				Name:    strings.ToUpper(subDistributorData.Name),
				Include: subDistributorData.Include,
				Exclude: subDistributorData.Exclude,
				Parent:  strings.ToUpper(subDistributorData.Parent),
			}) // Appending the new sub-distributor to the list
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
			for _, distributor := range distributorInformation {
				fmt.Printf("Name: %s, Include: %v, Exclude: %v, Parent: %s\n", distributor.Name, distributor.Include, distributor.Exclude, distributor.Parent)
			}
		case "5":
			fmt.Println("Exiting the program")
			return // Exiting the program
		}
	}
}
