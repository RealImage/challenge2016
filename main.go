package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

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

var LocationMap map[string]Location
var Distributors map[string][]distributor

func init() {
	Distributors = make(map[string][]distributor, 0)
	distributer1 := distributor{
		name:    strings.ToLower("DISTRIBUTOR1"),
		include: []string{"INDIA", "UNITEDstateS"},
		exclude: []string{"KARNATAKA-INDIA", "CHENNAI-TAMILNADU-INDIA"},
	}
	subDistributer := distributor{
		name:    strings.ToLower("DISTRIBUTOR2"),
		include: []string{"INDIA"},
		exclude: []string{"TAMILNADU-INDIA"},
	}
	distributer1.subDistributors = append(distributer1.subDistributors, subDistributer)
	Distributors[distributer1.name] = append(Distributors[distributer1.name], distributer1)

	file, err := os.Open("cities.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	LocationMap = make(map[string]Location)

	for _, row := range data {
		location := Location{
			city:    strings.ReplaceAll(strings.ToLower(row[3]), " ", "_"),
			state:   strings.ReplaceAll(strings.ToLower(row[4]), " ", "_"),
			country: strings.ReplaceAll(strings.ToLower(row[5]), " ", "_"),
		}
		LocationMap[location.city] = location
		LocationMap[location.state] = location
		LocationMap[location.country] = location

	}
}

func main() {

	for {
		fmt.Println("----- Choose one option ----- ")
		fmt.Println()
		fmt.Println("1. Add Distributor")
		fmt.Println("2. Add Permissions")
		fmt.Println("3. Search Permission By Location")
		fmt.Println("4. Add SubDistributor")
		fmt.Println("5. Exit")
		fmt.Println()

		var choice string
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			distributorName := ""
			fmt.Println("\n Enter the name of the distributor")
			fmt.Println()

			fmt.Scanln(&distributorName)
			if strings.TrimSpace(distributorName) == "" {
				fmt.Println("Distributor Name cannot be empty")
				fmt.Println()
				continue
			}
			distributorName = strings.ReplaceAll(distributorName, " ", "_")
			distributor := distributor{
				name:            distributorName,
				include:         []string{},
				exclude:         []string{},
				subDistributors: []distributor{},
			}
			AddPermission(&distributor, "nil", false)
			Distributors[distributorName] = append(Distributors[distributorName], distributor)
		case "2":
			distributorName := ""
			fmt.Println("Enter the name of the distributor or Subdistributor ")
			fmt.Scanln(&distributorName)
			distributor, rel := CheckDistributor(distributorName)
			var isSub bool
			if rel == "child" {
				isSub = true
			}
			if rel != "nil" {
				AddPermission(&distributor, "nil", isSub)
				Distributors[distributorName] = append(Distributors[distributorName], distributor)
			} else {
				fmt.Println(" Distributor or Sub Distributor does not exist")
			}
		case "3":
			distributorName := ""
			fmt.Println("Enter the name of the distributor or Subdistributor ")
			fmt.Scanln(&distributorName)
			distributor, rel := CheckDistributor(distributorName)
			if rel == "nil" {
				fmt.Println("Distributor or Sub Distributor does not exist")
				continue
			}
			SearchByLocation(distributor)
		case "4":
			AddSubDistributor()
		case "5":
			os.Exit(1)
		default:
			continue
		}
	}
}

func AddPermission(distributor *distributor, parentDistributorName string, isSubDistributor bool) {
	fmt.Println("---- Permissions ----")

	include_things := []string{}
	exclude_things := []string{}
permissionOpt:
	fmt.Println("Add Permissions")
	fmt.Println()
	fmt.Println("1. INCLUDED")
	fmt.Println("2. EXCLUDED")
	fmt.Println("3. LEAVE")
	fmt.Println()
	permissions := 0
	fmt.Scanln(&permissions)
	switch permissions {
	case 1, 2:
		fmt.Println("Enter Country Name: ")
		country, state, city := "", "", ""

		scanner.Scan()
		country = scanner.Text()

		if country != "nil" && country != "" {
			fmt.Println("Enter State name: ")

			scanner.Scan()
			state = scanner.Text()

			if state != "nil" && state != "" {
				fmt.Println("Enter City name: ")
				scanner.Scan()
				city = scanner.Text()
			}
		}

		include_perm := ""
		if (country != "nil" && country != "") || (state != "nil" && state != "") || (city != "nil" && city != "") {
			if country != "nil" && country != "" {
				include_perm += strings.ReplaceAll(strings.ToLower(country), " ", "_")
			}
			if state != "nil" && state != "" {
				include_perm = "_" + include_perm
				include_perm = strings.ReplaceAll(strings.ToLower(state), " ", "_") + include_perm
			}
			if city != "nil" && city != "" {
				include_perm = "_" + include_perm
				include_perm = strings.ReplaceAll(strings.ToLower(city), " ", "_") + include_perm
			}

		}

		if permissions == 1 {
			include_things = append(include_things, include_perm)
		} else {
			exclude_things = append(exclude_things, include_perm)
		}
		if isSubDistributor {
			Parent := FindParentDistributor(parentDistributorName)
			authStatus := false
			if country != "nil" && country != "" {
				if Parent.IsAuthorized(country, "country") {
					authStatus = true
				}
			}
			if state != "nil" && state != "" {
				if Parent.IsAuthorized(state, "state") {
					authStatus = true

				}
			}
			if city != "nil" && city != "" {
				if Parent.IsAuthorized(city, "city") {
					authStatus = true
				}
			}
			if authStatus {
				distributor.include = include_things
				distributor.exclude = exclude_things
			} else {
				fmt.Println("NOT AUTHORISED")
			}

		} else {
			distributor.include = include_things
			distributor.exclude = exclude_things

		}
	case 3:
		fmt.Println("Exited permission setup.")
		return
	default:
		fmt.Println("Invalid option. Please choose 1, 2, or 3.")
		goto permissionOpt

	}
	goto permissionOpt

}

func SearchByLocation(d distributor) {
	// search:
	var city, state, country string

	fmt.Println("Search by Location")
	fmt.Println()
	fmt.Println("1. Search By City")
	fmt.Println("2. Search By state")
	fmt.Println("3. Search By country")
	fmt.Println()
	var choice string
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		fmt.Println("Enter City Name")
		fmt.Scanln(&city)
		if d.IsAuthorized(city, "city") {
			fmt.Printf("YES\n")
		} else {
			fmt.Printf("NO\n")
		}
	case "2":
		fmt.Println("Enter State Name")
		fmt.Scanln(&state)
		if d.IsAuthorized(state, "state") {
			fmt.Printf("YES\n")
		} else {
			fmt.Printf("NO\n")
		}
	case "3":
		fmt.Println("Enter Country Name")
		fmt.Scanln(&country)
		if d.IsAuthorized(country, "country") {
			fmt.Printf("YES\n")
		} else {
			fmt.Printf("NO\n")
		}
	default:
		fmt.Println("Invalid option.")
	}
}

func (d distributor) IsAuthorized(region string, type_req string) bool {
	location, found := LocationMap[region]
	new_reg1 := ""
	new_reg2 := ""
	new_reg3 := ""

	if !found {
		fmt.Println("Location not found")
	} else {
		if type_req == "city" {
			new_reg1 = region + "_" + location.state + "_" + location.country
			new_reg2 = location.state + "_" + location.country
			new_reg3 = location.country
		} else if type_req == "state" {
			new_reg1 = location.state + "_" + location.country
			new_reg2 = location.country
			new_reg3 = location.country
		} else if type_req == "country" {
			new_reg1 = location.country
			new_reg2 = location.country
			new_reg3 = location.country
		}
	}
	for _, e := range d.exclude {
		if strings.HasPrefix(new_reg1, e) || strings.HasPrefix(new_reg2, e) || strings.HasPrefix(new_reg3, e) {
			return false
		}
	}

	for _, i := range d.include {
		if strings.HasPrefix(new_reg1, i) || strings.HasPrefix(new_reg2, i) || strings.HasPrefix(new_reg3, i) {
			return true
		}
	}

	return false
}

func FindParentDistributor(Distributor string) distributor {
	for key, value := range Distributors {
		if key == Distributor {
			return value[0]
		} else {
			for _, childDistributor := range value {
				if childDistributor.name == Distributor {
					return value[0]
				}
			}
		}
	}
	return distributor{}
}
func CheckDistributor(parentDistributor string) (distributor, string) {

	for key, value := range Distributors {
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

func AddSubDistributor() {
	parentDistributorName := ""
	fmt.Println("Enter the name of the Parent Distributor: ")
	fmt.Scanln(&parentDistributorName)

	parentDistributor, relation := CheckDistributor(parentDistributorName)
	if relation == "nil" {
		fmt.Println("Distributor or Sub Distributor does not exist")
		return
	}

	if relation == "child" {
		Distributors[parentDistributor.name] = append(Distributors[parentDistributor.name], parentDistributor)
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

	AddPermission(&subDistributor, parentDistributorName, true)
	Distributors[parentDistributor.name] = append(Distributors[parentDistributor.name], subDistributor)
	fmt.Printf("Sub Distributor '%s' added to Parent Distributor '%s'\n", subDistributorName, parentDistributorName)

}
