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
			status = "pass"
		} else {
			inputLines = append(inputLines, text)
		}
		if len(inputLines) == 1 && action == "add" {
			fmt.Printf("Please list the permissions,\n")
		}

	}

	return inputLines

}

func GetDistType() string {
	var status string
	var text string

	for status != "pass" {
		fmt.Printf("\nEnter the distributor type (Direct/Indirect): ")
		scanner.Scan()
		text = strings.ToLower(scanner.Text())
		if text == "direct" {
			status = "pass"
		} else if text == "indirect" {
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
					currentUser := PrepareSubUser(permissions, cities, findUser)
					if currentUser["err"] == nil {
						fmt.Printf("Valid permissions !!\n")
					} else {
						fmt.Printf("%v", currentUser["err"])
					}
				}
				fmt.Printf("\nDistributor not exist :( ")
				status = "fail"
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
