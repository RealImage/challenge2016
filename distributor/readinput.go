package distribution

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func GetInput() []string {
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
			fmt.Printf("\nEnter the Root - Distributor name: ")
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

func CallClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ActionIdentifier(directUserList []string, indirectUserList []string, distributorMap map[string]interface{}, cities []Cities) {
	var status string
	var text string

	for status != "pass" {
		fmt.Printf("1. Continue adding distributors?\n2. View the existing distributors?\nPlease select your choice")
		scanner.Scan()
		text = scanner.Text()
		if text == "1" || text == "2" {
			if text == "1" {
				status = "pass"
			} else {
				CallClear()
				if len(directUserList) > 0 {
					fmt.Printf("Direct Distributors\n")
					for _, user := range directUserList {
						fmt.Printf("%v\n", user)
					}
				}

				fmt.Printf("\n\n")

				if len(indirectUserList) > 0 {
					fmt.Printf("\n\n")
					fmt.Printf("Sub - Distributors\n")
					for _, user := range indirectUserList {
						fmt.Printf("%v\n", user)
					}
				}

				fmt.Printf("\nPlease enter any Distributor name to check permission: ")
				scanner.Scan()
				givenName := scanner.Text()
				findUser := distributorMap[givenName].(map[string]interface{})
				if findUser != nil {
					fmt.Printf("\nPlease enter the permission to check whether it is valid or not: \n")
					permissions := GetInput()
					userType := mustString(findUser, "type", "")
					if userType == "direct" {
						currentUser := PrepareRoorUser(permissions, cities)
						if currentUser["err"] == nil {
							fmt.Printf("Valid permissions !!\n")
						} else {
							fmt.Printf("%v", currentUser["err"])
						}
					} else if userType == "indirect" {
						currentUser := PrepareSubUser(permissions, cities, findUser)
						if currentUser["err"] == nil {
							fmt.Printf("Valid permissions !!\n")
						} else {
							fmt.Printf("%v", currentUser["err"])
						}
					}
				}
				status = "pass"
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
