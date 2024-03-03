package main

import (
	"fmt"
	"image-challenge/processor"
	"os"
)

func main() {
	processor.Preprocess()

	options := map[string]func(){
		"1": processor.AddSubDistributor,
		"2": processor.AddDistributor,
		"3": processor.UpdatePermissions,
		"4": processor.CheckForAccess,
		"5": func() {
			os.Exit(1)
		},
	}

	for {
		printMenu()
		selection := getUserInput()
		if action, ok := options[selection]; ok {
			action()
		} else {
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}

func printMenu() {
	fmt.Println("Select an option to proceed:")
	fmt.Println("1. Add Sub-distributor")
	fmt.Println("2. Add Distributor")
	fmt.Println("3. Update Permissions for a distributor")
	fmt.Println("4. Check for access")
	fmt.Println("5. Exit the program")
}

func getUserInput() string {
	var selection string
	fmt.Scanln(&selection)
	return selection
}
