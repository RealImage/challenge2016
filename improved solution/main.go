package main

// @AUTHOR: Yash Chauhan @iyashjayesh
// @Language: Golang
// @Last Updated: 26 April 2023
// @Description: This is the main file for the CLI tool for the Real Image Challenge. (Improved Solution)

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

	fmt.Println("#############################|Real Image Challenge CLI TOOL|#############################")
	fmt.Println("#############################|Author: Yash Chauhan|#############################")
	fmt.Println(" 						   ")
	var id int = 0

	for {

		fmt.Println("######## MAIN MENU ########")
		utils.GetMainMenu()
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Adding a distributor to the list
			fmt.Println("")
			fmt.Println("#### ADDING A DISTRIBUTOR WITH PERMISSIONS ####")
			Distributor.AddDistributor(&id)
		case 2:
			// Printing the distributor from the list
			fmt.Println("")
			fmt.Println("#### PRINTING THE DISTRIBUTOR LIST ####")
			Distributor.ListDistributors()
		case 3:
			// Checking the distributor permissions
			fmt.Println("")
			fmt.Println("#### CHECKING THE DISTRIBUTOR PERMISSIONS ####")
			// handler.CheckDistributorPermission(&cities, &DistributorsList)
			Distributor.CheckPermission()
		case 4:
			// Creating the network of distributors
			fmt.Println("")
			fmt.Println("#### CREATING THE NETWORK OF DISTRIBUTORS ####")
			//@TODO: Working on this
			// handler.DistributorNetwork(&DistributorsList)
		case 5:
			// Get back to the main menu
			fmt.Println("")
			fmt.Println("#### GETTING BACK TO THE MAIN MENU ####")
			utils.GetMainMenu()
		case 6:
			fmt.Println("Exiting from the CLI Tool")
			os.Exit(0)
		default:
			fmt.Println("Invalid Choice")
		}
	}
}
