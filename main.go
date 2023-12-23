package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

var inputScanner = bufio.NewScanner(os.Stdin)

type distributor struct {
	name            string
	include         []string
	exclude         []string
	subDistributors []distributor
}

type Location struct {
	city    string
	state   string
	country string
}

var locationMap map[string]Location
var distributorMap map[string][]distributor

func init() {
	distributorMap = make(map[string][]distributor, 0)
	mainDistributor := distributor{
		name:    strings.ToLower("MAIN_DISTRIBUTOR"),
		include: []string{"INDIA", "UNITED_STATES"},
		exclude: []string{"KARNATAKA-INDIA", "CHENNAI-TAMILNADU-INDIA"},
	}
	subDistributor := distributor{
		name:    strings.ToLower("SUB_DISTRIBUTOR"),
		include: []string{"INDIA"},
		exclude: []string{"TAMILNADU-INDIA"},
	}
	mainDistributor.subDistributors = append(mainDistributor.subDistributors, subDistributor)
	distributorMap[mainDistributor.name] = append(distributorMap[mainDistributor.name], mainDistributor)

	file, err := os.Open("./cities.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	locationMap = make(map[string]Location)

	for _, row := range data {
		location := Location{
			city:    strings.ReplaceAll(strings.ToLower(row[3]), " ", "_"),
			state:   strings.ReplaceAll(strings.ToLower(row[4]), " ", "_"),
			country: strings.ReplaceAll(strings.ToLower(row[5]), " ", "_"),
		}
		locationMap[location.city] = location
		locationMap[location.state] = location
		locationMap[location.country] = location
	}
}

func main() {
	for {
		fmt.Println("------------Main Menu----------------")
		fmt.Print(" 1.Add Distributor \n 2.Add Permissions \n 3.Search Permission By Location \n 4. Add SubDistributor \n 5.Exit ")
		var choice string
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			distributorName := ""
			fmt.Print("\n Enter the name of the distributor \n \n")
			fmt.Scanln(&distributorName)
			if strings.TrimSpace(distributorName) == "" {
				fmt.Print("Distributor Name cannot be empty \n")
				continue
			}
			distributorName = strings.ReplaceAll(distributorName, " ", "_")
			distributor := distributor{
				name:            distributorName,
				include:         []string{},
				exclude:         []string{},
				subDistributors: []distributor{},
			}
			addPermission(&distributor, "nil", false)
			distributorMap[distributorName] = append(distributorMap[distributorName], distributor)
		case "2":
			distributorName := ""
			fmt.Println("Enter the name of the distributor or Sub Distributor ")
			fmt.Scanln(&distributorName)
			distributor, rel := checkDistributor(distributorName)
			var isSub bool
			if rel == "child" {
				isSub = true
			}
			if rel != "nil" {
				addPermission(&distributor, "nil", isSub)
				distributorMap[distributorName] = append(distributorMap[distributorName], distributor)
			} else {
				fmt.Println(" Distributor or Sub Distributor does not exist")
			}
		case "3":
			distributorName := ""
			fmt.Println("Enter the name of the distributor or Sub Distributor ")
			fmt.Scanln(&distributorName)
			distributor, rel := checkDistributor(distributorName)
			if rel == "nil" {
				fmt.Println("Distributor or Sub Distributor does not exist")
				continue
			}
			searchByLocation(distributor)
		case "4":
			addSubDistributor()
		case "5":
			os.Exit(1)
		default:
			continue
		}
	}
}

func addPermission(distributor *distributor, parentDistributorName string, isSubDistributor bool) {
	fmt.Println("-----------Permission Menu-------------")

	includeThings := []string{}
	excludeThings := []string{}
permissionOpt:
	fmt.Print("Adding Permissions..\n  1. Include \n 2. Exclude \n  3.Exit \n")
	permissions := 0
	fmt.Scanln(&permissions)
	switch permissions {
	case 1, 2:
		fmt.Println("Enter Country name(enter nil if not needed) : ")
		country, state, city := "", "", ""
		inputScanner.Scan()
		country = inputScanner.Text()
		if country != "nil" && country != "" {
			fmt.Println("Enter State name(enter nil if not needed) :  ")
			inputScanner.Scan()
			state = inputScanner.Text()
			if state != "nil" && state != "" {
				fmt.Println("Enter City name(enter nil if not needed) : ")
				inputScanner.Scan()
				city = inputScanner.Text()
			}
		}

		includePerm := ""
		if (country != "nil" && country != "") || (state != "nil" && state != "") || (city != "nil" && city != "") {
			if country != "nil" && country != "" {
				includePerm += strings.ReplaceAll(strings.ToLower(country), " ", "_")
			}
			if state != "nil" && state != "" {
				includePerm = "_" + includePerm
				includePerm = strings.ReplaceAll(strings.ToLower(state), " ", "_") + includePerm
			}
			if city != "nil" && city != "" {
				includePerm = "_" + includePerm
				includePerm = strings.ReplaceAll(strings.ToLower(city), " ", "_") + includePerm
			}

		}

		if permissions == 1 {
			includeThings = append(includeThings, includePerm)
		} else {
			excludeThings = append(excludeThings, includePerm)
		}
		if isSubDistributor {
			parent := findParentDistributor(parentDistributorName)
			authStatus := false
			if country != "nil" && country != "" {
				if parent.isAuthorized(country, "country") {
					authStatus = true
				}
			}
			if state != "nil" && state != "" {
				if parent.isAuthorized(state, "state") {
					authStatus = true
				}
			}
			if city != "nil" && city != "" {
				if parent.isAuthorized(city, "city") {
					authStatus = true
				}
			}
			if authStatus {
				distributor.include = includeThings
				distributor.exclude = excludeThings
			} else {
				fmt.Println("You are not authorized to access this location")
			}

		} else {
			distributor.include = includeThings
			distributor.exclude = excludeThings

		}
	case 3:
		fmt.Println("Exiting permission setup.")
		return
	default:
		fmt.Println("Invalid option. Please choose 1, 2, or 3.")
		goto permissionOpt
	}
	goto permissionOpt
}

func searchByLocation(d distributor) {
	var city, state, country string
	fmt.Println("------Search By Location Menu -------- ")
	fmt.Print("Select 1.Search By City \n 2. Search By state \n 3. Search By country \n")
	var option string
	fmt.Scanln(&option)

	switch option {
	case "1":
		fmt.Println("Enter City Name")
		fmt.Scanln(&city)
		if d.isAuthorized(city, "city") {
			fmt.Printf("YES\n")
		} else {
			fmt.Printf("NO\n")
		}
	case "2":
		fmt.Println("Enter State Name")
		fmt.Scanln(&state)
		if d.isAuthorized(state, "state") {
			fmt.Printf("YES\n")
		} else {
			fmt.Printf("NO\n")
		}
	case "3":
		fmt.Println("Enter Country Name")
		fmt.Scanln(&country)
		if d.isAuthorized(country, "country") {
			fmt.Printf("YES\n")
		} else {
			fmt.Printf("NO\n")
		}
	default:
		fmt.Println("Invalid option.")
	}
}

func (d distributor) isAuthorized(region string, typeReq string) bool {
	location, found := locationMap[region]
	newReg1 := ""
	newReg2 := ""
	newReg3 := ""

	if !found {
		fmt.Println("Location not found")
	} else {
		if typeReq == "city" {
			newReg1 = region + "_" + location.state + "_" + location.country
			newReg2 = location.state + "_" + location.country
			newReg3 = location.country
		} else if typeReq == "state" {
			newReg1 = location.state + "_" + location.country
			newReg2 = location.country
			newReg3 = location.country
		} else if typeReq == "country" {
			newReg1 = location.country
			newReg2 = location.country
			newReg3 = location.country
		}
	}
	for _, e := range d.exclude {
		if strings.HasPrefix(newReg1, e) || strings.HasPrefix(newReg2, e) || strings.HasPrefix(newReg3, e) {
			return false
		}
	}

	for _, i := range d.include {
		if strings.HasPrefix(newReg1, i) || strings.HasPrefix(newReg2, i) || strings.HasPrefix(newReg3, i) {
			return true
		}
	}

	return false
}

func findParentDistributor(distributorName string) distributor {
	for key, value := range distributorMap {
		if key == distributorName {
			return value[0]
		} else {
			for _, childDistributor := range value {
				if childDistributor.name == distributorName {
					return value[0]
				}
			}
		}
	}
	return distributor{}
}

func checkDistributor(parentDistributor string) (distributor, string) {
	for key, value := range distributorMap {
		if key == parentDistributor {
			return value[0], "parent"
		} else {
			for _, childDistributor := range value {
				if childDistributor.name == parentDistributor {
					return childDistributor, "child"
				}
			}
		}
	}
	return distributor{}, "nil"
}

func addSubDistributor() {
	parentDistributorName := ""
	fmt.Println("Enter the name of the Parent Distributor: ")
	fmt.Scanln(&parentDistributorName)

	parentDistributor, relation := checkDistributor(parentDistributorName)
	if relation == "nil" {
		fmt.Println("Distributor or Sub Distributor does not exist")
		return
	}

	if relation == "child" {
		distributorMap[parentDistributor.name] = append(distributorMap[parentDistributor.name], parentDistributor)
	}

	subDistributorName := ""
	fmt.Println("Enter the name of the Sub Distributor: ")
	fmt.Scanln(&subDistributorName)

	subDistributor := distributor{
		name:            subDistributorName,
		include:         []string{},
		exclude:         []string{},
		subDistributors: []distributor{},
	}

	addPermission(&subDistributor, parentDistributorName, true)

	distributorMap[parentDistributor.name] = append(distributorMap[parentDistributor.name], subDistributor)
	fmt.Printf("Sub Distributor '%s' added to Parent Distributor '%s'\n", subDistributorName, parentDistributorName)
}
