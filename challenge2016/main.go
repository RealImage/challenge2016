//Author: Vighnesh Pol	
//Date: 19/11/2023

package main


import (
	"RealImageSolution/handler"
	"RealImageSolution/utils"
	"fmt"
	"os"
	"sync"
)

var (
	Once        sync.Once
	Distributor handler.DistributorInterface = &handler.DistributorsModel{}
)

func init() {
	Once.Do(func() {
		fmt.Println("Loading cities from CSV file...")
		cityLoadStatus, err := Distributor.LoadCitiesFromCSV("./data/cities.csv")
		if err != nil {
			fmt.Println("Error in loading cities from CSV file")
			os.Exit(1)
		}

		if cityLoadStatus {
			fmt.Println("Cities loaded successfully")
		} else {
			fmt.Println("Error in loading cities from CSV file")
			os.Exit(1)
		}
	})
}

func main() {

	var id int = 0

	for {

		fmt.Println("Plase choose an option from below:")
		utils.GetMainMenu()
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Adding a distributor to the list
			fmt.Println("")
			fmt.Println("Add distributor with permission")
			Distributor.AddDistributor(&id)
		case 2:
			// Printing the distributor from the list
			fmt.Println("")
			fmt.Println("List of distributors")
			Distributor.ListDistributors()
		case 3:
			// Checking the distributor permissions
			fmt.Println("")
			fmt.Println("Validate the distributor's permissions")
			// handler.CheckDistributorPermission(&cities, &DistributorsList)
			Distributor.CheckPermission()
		case 4:
			// Creating the network of distributors
			fmt.Println("")
			fmt.Println("Create network of distributors")
			//@TODO: Working on this
			Distributor.CreateSubDistributorNetwork()
		case 5:
			// Get back to the main menu
			fmt.Println("")
			fmt.Println("Back to Main Menu")
			utils.GetMainMenu()
		case 6:
			fmt.Println("Exit")
			os.Exit(0)
		default:
			fmt.Println("Invalid Choice")
		}
	}
}
