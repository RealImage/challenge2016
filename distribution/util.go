package distribution

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)
var AreaToStateMap = make(map[string][]string)
var StateToCountryMap = make(map[string][]string)

func RetrieveAreas() []Area {
	csvFile, _ := os.Open("./cities.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var areas []Area
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		areas = append(areas, Area{
			// to prevent whitespaces in permission hash
			AreaCode:    strings.ToLower(strings.Replace(line[0], " ", "", -1)),
			StateCode:   strings.ToLower(strings.Replace(line[1], " ", "", -1)),
			CountryCode: strings.ToLower(strings.Replace(line[2], " ", "", -1)),
			Area:        strings.ToLower(strings.Replace(line[3], " ", "", -1)),
			State:       strings.ToLower(strings.Replace(line[4], " ", "", -1)),
			Country:     strings.ToLower(strings.Replace(line[5], " ", "", -1)),
		})
	}
	return areas
}

func GetInput() {

	var userSelectStatus string
	var status string
	var user User
	var err error

	fmt.Println("Here is a list of users with their respective permissions the")
	for _, v := range UserDataMap {
		fmt.Printf("UserID => %s \n", v.ID)
		fmt.Print("	Permissions: \n")
		for _, j := range v.Permissions {
			fmt.Printf("		%s\n", j)
		}
		fmt.Print("-----------------------\n")
	}
	fmt.Printf("Please type in any User id and press enter \n")

	for userSelectStatus != "selected" {
		scanner.Scan()
		user_id := strings.TrimSpace(scanner.Text())

		user, err = getUser(user_id)
		if err == nil {
			userSelectStatus = "selected"
		} else {
			fmt.Printf("error %+v", err)
		}

	}

	userPermissions, err := user.ParsePermission()
	if err != nil {
		fmt.Printf("Encountered error in user permissions: %+v \n", err)
		status = "pass"
	}

	for status != "pass" {
		fmt.Println("Please type in the location in format `area-state-country`")
		scanner.Scan()
		text := strings.ToLower(strings.TrimSpace(scanner.Text()))

		if text == "" {
			status = "pass"
		} else {
			if text != "" {
				val, _ := userPermissions[text]
				if val {
					fmt.Printf("user is allowed for distribution in %+v \n \n", text)
				} else {
					fmt.Printf("user has no permission in %+v \n \n", text)
				}
			}
		}
	}
}
