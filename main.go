package main

import (
	"awesomeProject43/handler"
	"awesomeProject43/models"
	"bufio"
	"fmt"
	"os"
)

func main() {
	distributors := make(map[string]models.Permissions)

	for {
		fmt.Println("1. Add Distributor with Permissions")
		fmt.Println("2. Check Permission for a Distributor")
		fmt.Println("3. Exit")
		fmt.Print("Select an option: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			handler.AddDistributor(&distributors)
		case "2":
			handler.CheckPermission(distributors)
		case "3":
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
