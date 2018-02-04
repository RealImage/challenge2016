package distribution

import (
	"bufio"
	"fmt"
	"os"
	//"os/exec"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func GetInput(action string) []string {
	var status string
	var inputLines []string

	for status != "pass" {
		scanner.Scan()
		text := scanner.Text()
		if text == "" {
			if len(inputLines) > 1 {
				status = "pass"
			} else {
				fmt.Printf("You are not allowed to create a user with no rules/actual rules of parent distributor\n")
			}
		} else {
			if text != "" {
				inputLines = append(inputLines, text)
			}
		}
		if len(inputLines) == 1 && action == "add" {
			fmt.Printf("Please list the permissions,\n")
		} else if len(inputLines) == 1 && action == "check" {
			return inputLines
		}

	}

	return inputLines

}

func GetDistType() string {
	var status string
	var text string

	for status != "pass" {
		fmt.Printf("\nEnter the distributor type (Direct/Sub): ")
		scanner.Scan()
		text = strings.ToLower(scanner.Text())
		if text == "direct" {
			status = "pass"
		} else if text == "sub" {
			fmt.Printf("\nEnter the Parent - Distributor name: ")
			scanner.Scan()
			text = scanner.Text()
			return text
		}

	}
	return text
}

func StringArray(data map[string]interface{}, key string) ([]string, string) {

	if data != nil {
		userObj := data[key]
		switch t := userObj.(type) {
		case []string:
			return t, ""
		case []interface{}:
			items := make([]string, 0)
			for i := range t {
				items = append(items, fmt.Sprintf("%s", t[i]))
			}
			return items, ""
		}

	}
	return nil, ""
}

/*func CallClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}*/

func ActionIdentifier(distributorMap map[string]interface{}, cities []Cities) {
	var status string
	var text string

	for status != "pass" {
		fmt.Printf("1. Continue adding distributors?\n2. View the existing distributors?\nPlease select your choice: ")
		scanner.Scan()
		text = scanner.Text()
		if text == "1" || text == "2" {
			if text == "1" {
				status = "pass"
			} else {
				//CallClear()
				fmt.Printf("Direct Distributors:\n %v", printDistributor(distributorMap, "direct"))
				fmt.Printf("\n\nSub Distributors:\n %v", printDistributor(distributorMap, "indirect"))

				fmt.Printf("\nPlease enter any Distributor name to check permission: ")
				scanner.Scan()
				givenName := scanner.Text()
				findUser := distributorMap[givenName].(map[string]interface{})
				if findUser != nil {
					fmt.Printf("\nPlease enter the permission to check whether it is valid or not: \n")
					permissions := GetInput("check")
					userType := mustString(findUser, "type", "")
					currentUser := make(map[string]interface{})
					if userType == "indirect" {
						currentUser = PrepareSubUser(permissions, cities, findUser)
					} else if userType == "direct" {
						currentUser = checkUserPerm(permissions, cities, findUser)
					}
					if currentUser["err"] == nil {
						fmt.Printf("Valid permission !!\n")
					} else {
						fmt.Printf("%v", currentUser["err"])
					}
				} else {
					fmt.Printf("\nDistributor not exist :( \n")
					status = "fail"
				}
			}
		} else {
			fmt.Printf("\nPlease enter a valid input :( ")
		}

	}
	return
}

func mustString(data map[string]interface{}, property string, defaultValue string) string {
	if data != nil {
		userObj := data[property]
		switch t := userObj.(type) {
		case string:
			return t
		}
	}
	return defaultValue
}

func printDistributor(distributorMap map[string]interface{}, distType string) string {
	var distString string
	for k := range distributorMap {
		dist := distributorMap[k].(map[string]interface{})
		userType := mustString(dist, "type", "")
		if userType == distType {
			distString = fmt.Sprintf("%v\n%v", distString, k)
		}
	}

	return distString
}

func checkUserPerm(permission []string, cities []Cities, findUser map[string]interface{}) map[string]interface{} {
	currentUser := make(map[string]interface{})
	if strings.Contains(permission[0], ":") {
		custom := strings.Split(permission[0], ":")
		custom1 := custom[1]
		if strings.Contains(custom1, "_") {
			detect := strings.Split(custom1, "_")
			countries, _ := StringArray(findUser, "countries")
			if len(detect) == 2 {
				allowedStates := extendStateAccess(cities, countries)
				stateName := ExistInArray(allowedStates, custom1)
				if stateName == "" {
					currentUser["err"] = fmt.Sprintf("[%v] Not permitted\n", permission[0])
				}
			} else if len(detect) == 3 {
				excludedStates, _ := StringArray(findUser, "excluded_states")
				allowedAreas := extendAreaAccess(cities, countries, excludedStates)
				areaName := ExistInArray(allowedAreas, custom1)
				if areaName == "" {
					currentUser["err"] = fmt.Sprintf("[%v] Not permitted\n", permission[0])
				}
			}
		} else {
			allCountries := getCountryNames(cities)
			countryName := ExistInArray(allCountries, custom1)
			if countryName == "" {
				currentUser["err"] = fmt.Sprintf("[%v] Not permitted\n", permission[0])
			}
		}
	} else {
		currentUser["err"] = fmt.Sprintf("[%v] Not permitted - mention the INCLUDE/EXCLUDE operation\n", permission[0])
	}
	return currentUser
}
