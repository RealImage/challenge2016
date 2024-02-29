package main

import (
	"fmt"
	"os"

	"realImage.com/m/controllers"
)

// Main Function from where user can trigger different flows
func main() {
	controllers.Preprocess()
	for {
		fmt.Println("Select a option to proceed ahead")
		fmt.Println("1. Add Sub-distributor")
		fmt.Println("2. Add Distributor")
		fmt.Println("3. Update Permissions for a distributor")
		fmt.Println("4. Check for access")
		fmt.Println("5. Exit the program")

		var selection string
		fmt.Scanln(&selection)
		switch selection {
		case "1":
			controllers.AddSubDistributor()
		case "2":
			controllers.AddDistributor()
		case "3":
			controllers.UpdatePermissions()
		case "4":
			controllers.CheckForAccess()
		case "5":
			os.Exit(1)
		default:
			continue
		}
	}
}
