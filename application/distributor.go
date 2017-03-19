package application

import (
	"../distribution"

	"fmt"
)

var distributors map[int]*distribution.Distributor

func init() {
	distributors = make(map[int]*distribution.Distributor)
}

func createDistributor() {
	var distributorName, isSubDistributor string
	var parentDistributor *distribution.Distributor

	fmt.Printf("\nDistributor name: ")
	fmt.Scan(&distributorName)

	if len(distributors) > 0 {
		for {
			fmt.Print("Is he a sub-distributor? (y/n): ")
			fmt.Scan(&isSubDistributor)
			if isSubDistributor != "y" && isSubDistributor != "n" {
				fmt.Println("!! Invalid option. Enter y or n.")
			} else {
				break
			}
		}
		if isSubDistributor == "y" {
			fmt.Print("Please select a parent distributor...")
			parentDistributor = distributors[chooseDistributor()]
		} else {
			parentDistributor = nil
		}
	}

	distributor := &distribution.Distributor{}
	distributor.Initialize(distributorName, parentDistributor)
	distributorNumber := len(distributors) + 1
	distributors[distributorNumber] = distributor

	fmt.Println("Created Distributor !!")
}

func chooseDistributor() (number int) {
	for {
		listDistributors()
		fmt.Print("\nChoose distributor number (Ex: 1, 2, 3, ..): ")
		fmt.Scan(&number)
		_, ok := distributors[number]
		if ok {
			return
		}
		fmt.Println("!! Invalid distributor number")
	}
}

func listDistributors() {
	fmt.Println("\nHere is the list of distributors...")
	for i, distributor := range distributors {
		fmt.Printf("[%v] %v\n", i, distributor.Name)
	}
}
