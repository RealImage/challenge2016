package main

import (
	"Qcube/data"
	"Qcube/models"
	"fmt"
	"io"
	"log"
	"os"
)

var DistributorMap = make(map[string]models.Distributor)

func init() {

	// print log messages
	log.SetOutput(io.Discard)
	log.SetPrefix("DEBUG: ")

	data.Load_data()
}

func main() {

	fmt.Println("Welcome to cinema business system")
	for {
		fmt.Println("------------Main Menu----------------")
		fmt.Println(" 1.Add Distributor \n 2.Add SubDistributor \n 3.View Distributors \n 4.Check Distributor Authorized in location \n 5.Exit ")
		var choice string
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			AddDistributor()
		case "2":
			AddSubDistributor()
		case "3":
			ViewDistributors()
		case "4":
			CheckDistributorByLocation()
		case "5":
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	}

}
