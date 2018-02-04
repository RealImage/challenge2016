package main

import (
	distributor "./distributor"
	"fmt"
)

var cities []distributor.Cities
var distributorMap map[string]interface{}

func init() {
	cities = distributor.PrepareCitiesJson()
	distributorMap = make(map[string]interface{})
}

func main() {

	for {
		var distType string
		if len(distributorMap) == 0 {
			fmt.Printf("By default you need to create a Direct Distributor initially\n")
			distType = "direct"
		} else {
			distributor.ActionIdentifier(distributorMap, cities)
			distType = distributor.GetDistType()
		}

		if distType == "direct" {
			fmt.Printf("Enter the Distributor name: ")
			permission := distributor.GetInput("add")
			//permission := []string{"DIST1", "INCLUDE: INDIA", "INCLUDE: UNITEDSTATES", "EXCLUDE: KARNATAKA-INDIA", "EXCLUDE: CHENNAI-TAMILNADU-INDIA"}
			valid := distributorMap[permission[0]]
			if valid == nil {
				prepareDistributor(permission)
			} else {
				fmt.Printf("Direct user already exist with this name\n\n")
			}

		} else {
			//valid := distributor.ExistInArray(directUserList, distType)
			valid := distributorMap[distType]
			if valid != nil {
				userCountry, _ := distributor.StringArray(valid.(map[string]interface{}), "countries")
				userStates, _ := distributor.StringArray(valid.(map[string]interface{}), "included_states")
				userCities, _ := distributor.StringArray(valid.(map[string]interface{}), "included_cities")
				if len(userCountry) > 0 {
					fmt.Printf("FYI: You have controlled access in Countries - %v\n", userCountry)
				} else if len(userStates) > 0 {
					fmt.Printf("FYI: You have controlled access in States - %v\n", userStates)
				} else if len(userCities) > 0 {
					fmt.Printf("FYI: You have controlled access in Cities - %v\n", userCities)
				}
				fmt.Printf("Enter the Sub - Distributor name: ")
				permission := distributor.GetInput("add")
				checkexistance := distributorMap[permission[0]]
				if checkexistance == nil {
					prepareSubDistributor(permission, valid.(map[string]interface{}), distType)
				} else {
					fmt.Printf("Disributor name already exist, try with a different name\n")
				}
				//permission := []string{"DIST2", "INCLUDE: KERALA-INDIA", "INCLUDE: PUNJAB-INDIA", "EXCLUDE: GUJARAT-INDIA"}
			} else {
				fmt.Printf("Please enter a valid Parent distributor name\n")
			}
		}
	}

}

/*CRITICAL: SALVA,BN,RO,Salva,Bistrita-Nasaud,Romania*/

func prepareDistributor(permission []string) {

	currentUser := distributor.PrepareRoorUser(permission, cities)
	if currentUser["err"] == nil {
		currentUser["type"] = "direct"
		distributorMap[permission[0]] = currentUser
		fmt.Printf("Distributor created successfully !!\n\n")
	} else {
		fmt.Printf("%v", currentUser["err"])
	}
}

func prepareSubDistributor(permission []string, root map[string]interface{}, parent string) {
	currentUser := distributor.PrepareSubUser(permission, cities, root)
	if currentUser["err"] == nil {
		currentUser["type"] = "indirect"
		currentUser["parent"] = parent
		distributorMap[permission[0]] = currentUser
		fmt.Printf("Sub-Distributor created successfully !!\n")
	} else {
		fmt.Printf("%v", currentUser["err"])
	}
}
