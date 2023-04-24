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
		cities = utils.ReadCitiesCSV("./data/cities.csv")
	})
}

func main() {
	fmt.Println("Real Image Challenge CLI TOOL")
	var id int = 0
	var DistributorsList []models.Distributor

	for {
		utils.GetMainMenu()
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			handler.AddDistributor(&id, &DistributorsList)
		case 2:
			// Print the distributor from the list
			handler.PrintDistList(&DistributorsList)
		case 3:
		case 4:
			handler.CheckDistributorPermission(&cities, &DistributorsList)
		case 5:
			utils.GetMainMenu()
		case 6:
			fmt.Println("Exiting from the CLI Tool")
			os.Exit(0)
		default:
			fmt.Println("Invalid Choice")
		}
	}
}
