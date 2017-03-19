package application

import (
	"fmt"
	"os"
)

func RunApplication() {
	for {
		action := getAction()
		switch action {
		case 1:
			createDistributor()
		case 2:
			listDistributors()
		case 3:
			updateLocation()
		case 4:
			verifyLocation()
		case 5:
			os.Exit(0)
		}
	}
}

func getAction() (action int) {
	for {
		fmt.Print(`
1. Create Distributor
2. List Distributors
3. Include/Exclude locations
4. Verify Distribution location
5. Exit

Choose an action (Enter any action from 1 to 5): `)
		fmt.Scan(&action)
		if action >= 1 && action <= 5 {
			return
		}
		fmt.Println("!! Invalid action")
	}
}
