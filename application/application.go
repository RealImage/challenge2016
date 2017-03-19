package application

import (
	//"../distribution"

	"fmt"
	"os"
)

func RunApplication() {
	/*
		1. Create Distributor
		2. Include/Exclude locations
		3. Verify Distribution location
		4. Exit
		Choose an action:

		1. Create Distributor
		Distributor name:
		Is he a sub-distributor? (yes/no):
		What is the name of parent distributor?

		2. Include/Exclude locations:
		Distributor name:
		Include/Exclude? (i/e):
		Location Code (Ex: IN, TN-IN, CENAI-TN-IN):

		3. Verify Distribution Area
		Distributor name:
		Location Code (Ex: IN, TN-IN, CENAI-TN-IN):
	*/

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
