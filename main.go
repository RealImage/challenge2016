package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type triplet struct {
	city    string
	state   string
	country string
	typ     int
}

type city struct {
	name  string
	code  string
	state *state
}

type state struct {
	name    string
	code    string
	country *country
	citySet map[string]*city
}

type country struct {
	name     string
	code     string
	stateSet map[string]*state
	citySet  map[string]*city
}

type distributor struct {
	name           string
	master         *distributor
	subDistributor map[string]*distributor
	inclSet        map[string]*triplet
	exclSet        map[string]*triplet
}

const (
	csvFilePath = "cities.csv"
)

var (
	cityMap        map[string]*city
	stateMap       map[string]*state
	countryMap     map[string]*country
	distributorMap map[string]*distributor
)

func main() {
	records, err := readCsvFile()
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}
	initMaps()
	createRecordsMap(records)
	showMainMenu()
}

func readCsvFile() ([][]string, error) {
	file, err := os.Open(csvFilePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV: %v", err)
	}
	return records, nil
}

func showMainMenu() {
	for {
		var (
			input    int
			distName string
		)
		fmt.Println("\nMain Menu -> 1.Add Distributor  2.Show All Entities  3.Check Distributor Authorization  4.End")
		fmt.Scanf("%d %s", &input, &distName)
		if distName == "" && (input == 1 || input == 3) {
			fmt.Println("\nPlease Enter Distributor Name with option 1 and 3")
			fmt.Println("Eg : 1 <Distributor Name> / 2 / 3 <Distributor Name> / 4")
			continue
		}
		switch input {
		case 1:
			addDistributor(distName, nil)
		case 2:
			showAllMenu()
		case 3:
			checkAuthorization(distName)
		default:
			return
		}
	}
}

func addDistributor(distName string, master *distributor) {
	dist := getDistributor(distName, master)
	showPermissionMenu(dist)
}

func getDistributor(distName string, master *distributor) *distributor {
	dist, ok := distributorMap[distName]
	if !ok {
		dist = &distributor{
			name:           distName,
			master:         master,
			subDistributor: make(map[string]*distributor),
			inclSet:        make(map[string]*triplet),
			exclSet:        make(map[string]*triplet),
		}
		if master != nil {
			master.subDistributor[distName] = dist
		}
		distributorMap[distName] = dist
	}
	return dist
}

func showPermissionMenu(dist *distributor) {
	for {
		var (
			input   int
			perName string
		)
		fmt.Printf("\n%s Distributor Menu -> 1.Include Place  2.Exclude Place  3.Add Sub-Distributor  4.Show Info  5.Back\n", dist.name)
		fmt.Scanf("%d %s", &input, &perName)
		if input > 4 {
			return
		}
		if input == 4 {
			showDistributor(dist)
		} else {
			if perName == "" {
				fmt.Println("\nPlease enter Place Code or Distributor Name with 1, 2, 3 option")
				fmt.Println("Eg : 1 <Place Code> / 2 <Place Code> / 3 <Distributor Name> / 4 / 5 ")
				continue
			}
			switch input {
			case 1:
				permissionAction(dist, input, perName)
			case 2:
				permissionAction(dist, input, perName)
			case 3:
				addDistributor(perName, dist)
			default:
				fmt.Println("Wrong option selected")
			}
		}
	}
}

func showAllMenu() {
	for {
		var (
			input int
		)
		fmt.Println("\nEntity Menu -> 1.Show All Distributor  2.Show All Countries  3.Show All Provinces  4.Show All Cities  5.Back")
		fmt.Scanf("%d", &input)
		switch input {
		case 1:
			showAllDistributors()
		case 2:
			showAllCountry()
		case 3:
			showAllState()
		case 4:
			showAllCity()
		default:
			return
		}
	}
}

func checkAuthorization(distName string) {
	if distName == "" {
		fmt.Println("Distributor Name is empty")
		return
	}
	dist, ok := distributorMap[distName]
	if !ok {
		fmt.Println("Distributor not present. Try using add Distributor option")
		return
	}
	showAuthorizationMenu(dist)
}

func showAuthorizationMenu(dist *distributor) {
	for {
		var (
			input   int
			perName string
		)
		fmt.Printf("Disributor %s Authorization Menu -> 1. Is Place Authorized  2.Back\n", dist.name)
		fmt.Scanf("%d %s", &input, &perName)
		switch input {
		case 1:
			if perName == "" {
				fmt.Println("\nPlease Enter Place Code with option 1")
				fmt.Println("Eg : 1 <Place Code> / 2 ")
				continue
			}
			if !isPlaceCodeValid(perName) {
				fmt.Printf("Place Code %s is not valid\n", perName)
				fmt.Printf("Please Enter Place Code in format <Country Code-Province Code-City Code> \n")
				fmt.Printf("Eg : For CHENNAI Enter this Code -> IN-TN-CENAI\n")
				continue
			}
			auth := checkAuthorized(dist, perName)
			fmt.Printf("%s Authorization : %v ", perName, auth)
		default:
			return
		}
	}
}

func showCountryMenu() {
	for {
		var (
			input       int
			countryName string
		)
		fmt.Println("Country Menu -> 1.Show All Province In Country  2.Show All City In Country  3.Back")
		fmt.Scanf("%d %s", &input, &countryName)
		if countryName == "" && (input == 1 || input == 2) {
			fmt.Println("\nPlease Enter Country Code with option 1 and 2")
			fmt.Println("Eg : 1 <Country Code> / 2 <Country Code> / 3")
			continue
		}
		countryName = strings.ToUpper(countryName)
		switch input {
		case 1:
			printAllStateInCountry(countryName)
		case 2:
			printAllCityInCountry(countryName)
		default:
			return
		}
	}

}

func showStateMenu() {
	for {
		var (
			input     int
			stateName string
		)
		fmt.Println("Province Menu -> 1.Show All City In Province  2.Back")
		fmt.Scanf("%d %s", &input, &stateName)
		if stateName == "" && input == 1 {
			fmt.Println("\nEnter Province Code with option 1")
			fmt.Println("Eg : 1 <Province Code> / 2 ")
		}
		stateName = strings.ToUpper(stateName)
		switch input {
		case 1:
			printAllCityInState(stateName)
		default:
			return
		}
	}

}

func showAllCity() {
	showAllCityInMap()
}

func showAllState() {
	showAllStateInMap()
	showStateMenu()
}

func showAllCountry() {
	showAllCountryInMap()
	showCountryMenu()
}

func showAllCityInMap() {
	fmt.Println()
	fmt.Printf("All Cities : %d\n", len(cityMap))
	for _, city := range cityMap {
		fmt.Printf("%s -> %s\n", city.name, city.code)
	}
	fmt.Printf("All Cities : %d\n", len(cityMap))
	fmt.Println()
}

func showAllStateInMap() {
	fmt.Println()
	fmt.Printf("All Provinces : %d\n", len(stateMap))
	for _, state := range stateMap {
		fmt.Printf("%s -> %s\n", state.name, state.code)
	}
	fmt.Printf("All Provinces : %d\n", len(stateMap))
	fmt.Println()
}

func showAllCountryInMap() {
	fmt.Println()
	fmt.Printf("All Countries : %d\n", len(countryMap))
	for _, country := range countryMap {
		fmt.Printf("%s -> %s\n", country.name, country.code)
	}
	fmt.Printf("All Countries : %d\n", len(countryMap))
	fmt.Println()
}

func printAllStateInCountry(countryName string) {
	if countryName == "" {
		fmt.Println("Country Code is empty")
		return
	}
	if country, ok := countryMap[countryName]; ok {
		fmt.Printf("%s Provinces : %d\n", country.name, len(country.stateSet))
		for _, state := range country.stateSet {
			fmt.Printf("%s -> %s\n", state.name, state.code)
		}
		fmt.Printf("%s Provinces : %d\n", country.name, len(country.stateSet))
		fmt.Println()
		showStateMenu()
	} else {
		fmt.Printf("Country Code %s is invalid\n", countryName)
	}
}

func printAllCityInState(stateName string) {
	if stateName == "" {
		fmt.Println("Province Code is empty")
		return
	}
	if state, ok := stateMap[stateName]; ok {
		fmt.Printf("%s Cities : %d\n", state.name, len(state.citySet))
		for _, city := range state.citySet {
			fmt.Printf("%s -> %s\n", city.name, city.code)
		}
		fmt.Printf("%s Cities : %d\n", state.name, len(state.citySet))
		fmt.Println()
	} else {
		fmt.Printf("Province Code %s is not Valid\n", stateName)
		fmt.Printf("Please Enter Province Code in format <Country Code-Province Code> \n")
		fmt.Printf("Eg : For Tamil Nadu Enter this Code -> IN-TN\n")
	}
}

func printAllCityInCountry(countryName string) {
	if countryName == "" {
		fmt.Println("Country Code is empty")
		return
	}
	if country, ok := countryMap[countryName]; ok {
		fmt.Printf("%s Cities : %d\n", country.name, len(country.citySet))
		for _, city := range country.citySet {
			fmt.Printf("%s -> %s\n", city.name, city.code)
		}
		fmt.Printf("%s Cities : %d\n", country.name, len(country.citySet))
		fmt.Println()
	} else {
		fmt.Printf("Country Code %s is not Valid\n", countryName)
	}
}

func checkAuthorized(dist *distributor, perName string) bool {
	trip := createTriplet(perName)
	return checkDistributorAuthorization(dist, trip)
}

func checkDistributorAuthorization(dist *distributor, trip triplet) bool {
	if checkCityExcl(dist, trip) || checkStateExcl(dist, trip) || checkCountryExcl(dist, trip) {
		return false
	}
	if checkCityIncl(dist, trip) || checkStateIncl(dist, trip) || checkCountryIncl(dist, trip) {
		return true
	}
	return false
}

func checkCityIncl(dist *distributor, trip triplet) bool {
	_, ok := dist.inclSet[trip.city]
	return ok
}

func checkCityExcl(dist *distributor, trip triplet) bool {
	_, ok := dist.exclSet[trip.city]
	return ok
}

func checkStateIncl(dist *distributor, trip triplet) bool {
	_, ok := dist.inclSet[trip.state]
	return ok
}

func checkStateExcl(dist *distributor, trip triplet) bool {
	_, ok := dist.exclSet[trip.state]
	return ok
}

func checkCountryIncl(dist *distributor, trip triplet) bool {
	_, ok := dist.inclSet[trip.country]
	return ok
}

func checkCountryExcl(dist *distributor, trip triplet) bool {
	_, ok := dist.exclSet[trip.country]
	return ok
}

func createTriplet(perName string) triplet {
	var auth triplet
	dashCount := strings.Count(perName, "-")
	switch dashCount {
	case 0:
		if country, ok := countryMap[perName]; ok {
			auth.country = country.code
			auth.typ = 3
		}
	case 1:
		if state, ok := stateMap[perName]; ok {
			auth.state = state.code
			auth.country = state.country.code
			auth.typ = 2
		}
	case 2:
		if city, ok := cityMap[perName]; ok {
			auth.city = city.code
			auth.state = city.state.code
			auth.country = city.state.country.code
			auth.typ = 1
		}
	}
	return auth
}

func printMasterDist(dist *distributor) {
	if dist == nil {
		return
	}
	fmt.Printf("< %s ", dist.name)
	printMasterDist(dist.master)
}

func showAllDistributors() {
	fmt.Printf("\nDistributors : %d\n", len(distributorMap))
	for _, dist := range distributorMap {
		showDistributor(dist)
	}
	fmt.Printf("\nDistributors : %d\n", len(distributorMap))
	fmt.Println()
}

func showDistributor(dist *distributor) {
	fmt.Printf("\nName : %s ", dist.name)
	printMasterDist(dist.master)
	fmt.Printf("\nIncluded : ")
	for name := range dist.inclSet {
		fmt.Printf("%s, ", name)
	}
	fmt.Printf("\nExcluded : ")
	for name := range dist.exclSet {
		fmt.Printf("%s, ", name)
	}
	fmt.Println()
}

func permissionAction(dist *distributor, action int, perName string) {
	perName = strings.ToUpper(perName)
	if !isPlaceCodeValid(perName) {
		fmt.Printf("Place Code %s is not valid\n", perName)
		fmt.Printf("Please Enter Place Code in format <Country Code-Province Code-City Code> \n")
		fmt.Printf("Eg : For CHENNAI Enter this Code -> IN-TN-CENAI\n")
		return
	}
	switch action {
	case 1:
		addPerInDistIncl(dist, perName)
	case 2:
		addPerInDistExcl(dist, perName)
	default:
		fmt.Println("Wrong option selected")
	}
}

func addPerInDistIncl(dist *distributor, perName string) {
	if _, ok := dist.inclSet[perName]; ok {
		fmt.Printf("Distributor %s already included %s\n", dist.name, perName)
		return
	}
	trip := createTriplet(perName)
	if len(dist.subDistributor) == 0 && isDistributorAuthorisedToInclu(dist, trip) {
		dist.inclSet[perName] = &trip
		addExclusionFromMaster(dist, trip)
	} else {
		fmt.Printf("Distributor %s is not authorized to include %s\n", dist.name, perName)
	}
}

func isDistributorAuthorisedToInclu(dist *distributor, trip triplet) bool {
	if dist.master != nil {
		if !isMasterAuthorisedToInclu(dist.master, trip) {
			fmt.Printf("Master %s is not authorized to include\n", dist.master.name)
			return false
		} else {
			return true
		}
	}
	switch trip.typ {
	case 1:
		if !checkCityExcl(dist, trip) && !checkStateExcl(dist, trip) {
			return true
		} else {
			return false
		}
	case 2:
		if !checkStateExcl(dist, trip) {
			return true
		} else {
			return false
		}
	default:
		return true
	}
}

func addExclusionFromMaster(dist *distributor, trip triplet) {
	if dist.master == nil {
		return
	}
	for name, t := range dist.master.exclSet {
		switch trip.typ {
		case 1:
		case 2:
			if t.typ == 1 {
				if isCountryStateSame(trip, t) {
					dist.exclSet[name] = t
				}
			}

		case 3:
			if t.typ == 2 || t.typ == 1 {
				if isCountrySame(trip, t) {
					dist.exclSet[name] = t
				}
			}
		}
	}
}

func isCountryStateSame(trip triplet, t *triplet) bool {
	return trip.country == t.country && trip.state == t.state
}

func isCountrySame(trip triplet, t *triplet) bool {
	return trip.country == t.country
}

func isMasterAuthorisedToInclu(dist *distributor, trip triplet) bool {
	switch trip.typ {
	case 1:
		if !checkCityExcl(dist, trip) && !checkStateExcl(dist, trip) && (checkStateIncl(dist, trip) || checkCountryIncl(dist, trip)) {
			return true
		} else {
			return false
		}
	case 2:
		if !checkStateExcl(dist, trip) && checkCountryIncl(dist, trip) {
			return true
		} else {
			return false
		}
	default:
		if checkCountryIncl(dist, trip) {
			return true
		} else {
			return false
		}
	}
}

func addPerInDistExcl(dist *distributor, perName string) {
	if _, ok := dist.exclSet[perName]; ok {
		fmt.Printf("Distributor %s already excluded %s\n", dist.name, perName)
		return
	}
	trip := createTriplet(perName)
	if len(dist.subDistributor) == 0 && isDistributorAuthorisedToExclu(dist, trip) {
		dist.exclSet[perName] = &trip
	} else {
		fmt.Printf("Distributor %s is not authorized to exclude %s\n", dist.name, perName)
	}
}

func isDistributorAuthorisedToExclu(dist *distributor, trip triplet) bool {
	switch trip.typ {
	case 1:
		if !checkCityIncl(dist, trip) && (checkStateIncl(dist, trip) || checkCountryIncl(dist, trip)) {
			return true
		} else {
			return false
		}
	case 2:
		if !checkStateIncl(dist, trip) && checkCountryIncl(dist, trip) {
			return true
		} else {
			return false
		}
	default:
		return false
	}
}

func isPlaceCodeValid(perName string) bool {
	return checkCity(perName) || checkState(perName) || checkCountry(perName)
}

func checkCity(perName string) bool {
	_, ok := cityMap[perName]
	return ok
}

func checkState(perName string) bool {
	_, ok := stateMap[perName]
	return ok
}

func checkCountry(perName string) bool {
	_, ok := countryMap[perName]
	return ok
}

func initMaps() {
	cityMap = make(map[string]*city)
	stateMap = make(map[string]*state)
	countryMap = make(map[string]*country)
	distributorMap = make(map[string]*distributor)
}

func createRecordsMap(records [][]string) {
	fmt.Printf("Total Records : %d\n", len(records))
	for i, record := range records {
		if i == 0 {
			continue
		}
		var (
			ctyCode string
			ctyName string
			staCode string
			staName string
			conCode string
			conName string
			cty     city
			sta     state
			con     country
		)

		for i, field := range record {
			field = strings.ToUpper(field)
			switch i {
			case 0:
				ctyCode = field
			case 1:
				staCode = field
			case 2:
				conCode = field
			case 3:
				ctyName = field
			case 4:
				staName = field
			case 5:
				conName = field
			}
		}

		cty.code = conCode + "-" + staCode + "-" + ctyCode
		sta.code = conCode + "-" + staCode
		con.code = conCode
		cty.name = conName + "-" + staName + "-" + ctyName
		sta.name = conName + "-" + staName
		con.name = conName

		_, isCon := countryMap[con.code]
		_, isSta := stateMap[sta.code]
		_, isCty := cityMap[cty.code]

		if !isCon {
			countryMap[con.code] = &con
		}
		if !isSta {
			stateMap[sta.code] = &sta
		}
		if !isCty {
			cityMap[cty.code] = &cty
		}

		coun := countryMap[con.code]
		stat := stateMap[sta.code]
		cety := cityMap[cty.code]

		cety.state = stat
		stat.country = coun

		if coun.stateSet == nil {
			coun.stateSet = make(map[string]*state)
		}
		coun.stateSet[sta.code] = stat
		if coun.citySet == nil {
			coun.citySet = make(map[string]*city)
		}
		coun.citySet[cty.code] = cety
		if stat.citySet == nil {
			stat.citySet = make(map[string]*city)
		}
		stat.citySet[cty.code] = cety
	}
}
