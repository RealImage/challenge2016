package main

import (
	distributor "./distributor"
	//"bufio"
	//"encoding/json"
	"fmt"
	//"os"
	//"os/exec"
	//"strconv"
	//"strings"
)

var cities []distributor.Cities
var distributorMap map[string]interface{}
var directUserList []string
var indirectUserList []string

func init() {
	cities = distributor.PrepareCitiesJson()
	distributorMap = make(map[string]interface{})
}

func main() {

	for {
		var distType string
		if len(directUserList) == 0 {
			fmt.Printf("By defaudirectUserListlt you need to create a Direct Distributor initially\n")
			distType = "direct"
		} else {
			distributor.ActionIdentifier(directUserList, indirectUserList, distributorMap, cities)
			distType = distributor.GetDistType()
		}

		if distType == "direct" {
			fmt.Printf("Enter the Distributor name: ")
			permission := distributor.GetInput()
			//permission := []string{"DIST1", "INCLUDE: INDIA", "INCLUDE: UNITEDSTATES", "EXCLUDE: KARNATAKA-INDIA", "EXCLUDE: CHENNAI-TAMILNADU-INDIA"}
			valid := distributor.ExistInArray(directUserList, permission[0])
			if valid == "" {
				prepareDirectUser(permission)
			} else {
				fmt.Printf("Direct user already exist with this name\n")
			}

		} else {
			valid := distributor.ExistInArray(directUserList, distType)
			if valid != "" {
				userCountry, _ := distributor.StringArray(distributorMap[valid].(map[string]interface{}), "countries")
				distributor.CallClear()
				fmt.Printf("FYI: You have controlled access in Countries - %v\n", userCountry)
				fmt.Printf("Enter the Sub - Distributor name: ")
				permission := distributor.GetInput()
				//permission := []string{"DIST2", "INCLUDE: KERALA-INDIA", "INCLUDE: PUNJAB-INDIA", "EXCLUDE: GUJARAT-INDIA"}

				prepareInDirectUser(permission, distributorMap[valid].(map[string]interface{}), valid)
			}
		}
	}

}

/*CRITICAL: SALVA,BN,RO,Salva,Bistrita-Nasaud,Romania*/

func prepareDirectUser(permission []string) {

	currentUser := distributor.PrepareRoorUser(permission, cities)
	if currentUser["err"] == nil {
		directUserList = append(directUserList, permission[0])
		currentUser["type"] = "direct"
		distributorMap[permission[0]] = currentUser
		fmt.Printf("Distributor creates successfully !!\n")
	} else {
		distributor.CallClear()
		fmt.Printf("%v", currentUser["err"])
	}

	//fmt.Printf("%v", distributorMap[permission[0]])
}

func prepareInDirectUser(permission []string, root map[string]interface{}, parent string) {
	currentUser := distributor.PrepareSubUser(permission, cities, root)
	if currentUser["err"] == nil {
		currentUser["type"] = "indirect"
		currentUser["parent"] = parent
		indirectUserList = append(indirectUserList, permission[0])
		distributorMap[permission[0]] = currentUser
		fmt.Printf("Sub-Distributor creates successfully !!\n")
	} else {
		distributor.CallClear()
		fmt.Printf("%v", currentUser["err"])
	}

	//fmt.Printf("%v", distributorMap[permission[0]])
}
