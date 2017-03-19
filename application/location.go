package application

import (
	"fmt"
)

func updateLocation() {
	distributor := distributors[chooseDistributor()]

	var updateType string
	for {
		fmt.Print("Include/Exclude? (i/e): ")
		fmt.Scan(&updateType)
		if updateType != "i" && updateType != "e" {
			fmt.Println("!! Invalid option. Enter i or e.")
		} else {
			break
		}
	}

	var location string
	fmt.Print("Location Code (Ex: IN, TN-IN, CENAI-TN-IN): ")
	fmt.Scan(&location)

	if updateType == "i" {
		err := distributor.Include(location)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		distributor.Exclude(location)
	}

	fmt.Println("Location is updated !!")
}

func verifyLocation() {
	distributor := distributors[chooseDistributor()]

	var location string
	fmt.Print("Location Code (Ex: IN, TN-IN, CENAI-TN-IN): ")
	fmt.Scan(&location)

	if distributor.CanDistribute(location) {
		fmt.Printf("YES, %v can distribute in %v\n", distributor.Name, location)
	} else {
		fmt.Printf("NO, %v cannot distribute in %v\n", distributor.Name, location)
	}
}
