package main

// @AUTHOR: Yash Chauhan @iyashjayesh
// @Language: Golang
// @Last Updated: 24 April 2023

import (
	"RealImageSolution/handler"
	"RealImageSolution/models"
	"RealImageSolution/utils"
	"fmt"
	"os"
	"sync"
)

var (
	cities []models.City
	Once   sync.Once
)

func init() {
	Once.Do(func() {
		fmt.Println("Loading cities from CSV file...")
		cities = handler.ReadCitiesCSV("./data/cities.csv")
	})
}

func main() {

	fmt.Println("#############################|Real Image Challenge CLI TOOL|#############################")
	fmt.Println("#############################|Author: Yash Chauhan|#############################")
	fmt.Println(" 						   ")
	var id int = 0
	var DistributorsList []models.Distributor

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
			handler.AddDistributor(&id, &DistributorsList)
		case 2:
			// Printing the distributor from the list
			fmt.Println("")
			fmt.Println("#### PRINTING THE DISTRIBUTOR LIST ####")
			handler.PrintDistList(&DistributorsList)
		case 3:
			// Checking the distributor permissions
			fmt.Println("")
			fmt.Println("#### CHECKING THE DISTRIBUTOR PERMISSIONS ####")
			handler.CheckDistributorPermission(&cities, &DistributorsList)
		case 4:
			// Creating the network of distributors
			fmt.Println("")
			fmt.Println("#### CREATING THE NETWORK OF DISTRIBUTORS ####")
			handler.DistributorNetwork(&DistributorsList)
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
