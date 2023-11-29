package main

import (
	"Qcube/data"
	"Qcube/models"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkDistributorName(distributorName string) bool {
	if distributorName == "" {
		fmt.Println("Distributor Name cannot be empty \n")
		return false
	}
	_, ok := DistributorMap[distributorName]
	if ok {
		fmt.Println("Distributor Name already exists \n")
		return false
	}
	return true
}

func verifyLocation(location models.DistributorLocation) bool {
	country := strings.TrimSpace(location.Country)
	if country == "" {
		fmt.Println("Country cannot be empty")
		return false
	}
	provinces, ok := data.Locations[location.Country]
	if !ok {
		fmt.Println("Country does not exist")
		return false
	}
	statefound := false
	for _, province := range provinces {
		cities, stateExist := province[location.State]
		if stateExist {
			statefound = true
			for _, city := range cities {
				if city == location.City {
					return true
				}
			}
		}
	}
	if !statefound && len(location.State) > 0 {
		fmt.Println("State does not exist")
		return false
	}
	if len(location.City) > 0 {
		fmt.Println("City does not exist")
		return false
	}
	return true
}

func checkDuplicate(distributor *models.Distributor, location models.DistributorLocation) bool {
	for _, include := range distributor.Include {
		if include.City == location.City && include.State == location.State && include.Country == location.Country {
			return true
		}
	}
	for _, exclude := range distributor.Exclude {
		if exclude.City == location.City && exclude.State == location.State && exclude.Country == location.Country {
			return true
		}
	}
	return false
}

func getLocationFromUser() (models.DistributorLocation, bool) {
	var country, state, city string
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	input = strings.TrimSpace(input)

	if input == "" {
		fmt.Println("Invalid input")
		return models.DistributorLocation{}, false
	}
	distributorLocation := models.DistributorLocation{}
	inputSlice := strings.Split(input, "-")
	len := len(inputSlice)
	log.Println("Input", inputSlice, " input length", len)
	if len == 1 {
		country = inputSlice[0]
		distributorLocation.Country = strings.ToLower(strings.TrimSpace(country))
	} else if len == 2 {
		country = inputSlice[1]
		state = inputSlice[0]
		distributorLocation.Country = strings.ToLower(strings.TrimSpace(country))
		distributorLocation.State = strings.ToLower(strings.TrimSpace(state))
	} else if len == 3 {
		country = inputSlice[2]
		state = inputSlice[1]
		city = inputSlice[0]
		distributorLocation.Country = strings.ToLower(strings.TrimSpace(country))
		distributorLocation.State = strings.ToLower(strings.TrimSpace(state))
		distributorLocation.City = strings.ToLower(strings.TrimSpace(city))
	}
	log.Println("Location", distributorLocation)
	if !verifyLocation(distributorLocation) {
		fmt.Println("Invalid location")
		return models.DistributorLocation{}, false
	}
	return distributorLocation, true
}

func addIncludePermission(distributor *models.Distributor) bool {
	fmt.Println("Enter Include Location:\n")
	distributorLocation, ok := getLocationFromUser()
	if !ok {
		return false
	}
	if checkDuplicate(distributor, distributorLocation) {
		fmt.Println("Duplicate location")
		return false
	}
	if distributor.Parent == nil {
		distributor.Include = append(distributor.Include, distributorLocation)
		fmt.Println("Include location added successfully")
		return true
	}
	parent := distributor.Parent
	isAuthorized := false
	for _, parentLocation := range parent.Include {
		if parentLocation.Country == distributorLocation.Country {
			if parentLocation.State == distributorLocation.State || parentLocation.State == "" {
				if parentLocation.City == distributorLocation.City || parentLocation.City == "" {
					isAuthorized = true
					break
				}
			}
		}
	}
	if isAuthorized {
		var tempExcludeLocation []models.DistributorLocation
		for _, excludeLocation := range parent.Exclude {
			if excludeLocation.Country == distributorLocation.Country {
				if excludeLocation.State == distributorLocation.State || excludeLocation.State == "" {
					if excludeLocation.City == distributorLocation.City || excludeLocation.City == "" {
						fmt.Println("Cannot include location which is excluded by parent")
						return false
					}
				}
				// handlng scenario where parent has excluded state and child has included country
				if distributorLocation.State == "" && excludeLocation.State != "" {
					tempExcludeLocation = append(tempExcludeLocation, excludeLocation)
				} else if (distributorLocation.State == excludeLocation.State) && (distributorLocation.City == "" && excludeLocation.City != "") {
					tempExcludeLocation = append(tempExcludeLocation, excludeLocation)
				}
			}
		}
		distributor.Exclude = tempExcludeLocation
		distributor.Include = append(distributor.Include, distributorLocation)
		fmt.Println("Include location added successfully")
		return true
	}
	fmt.Println("Cannot include location as parent distributor does not have permission for this location")
	return false
}

func addExcludePermission(distributor *models.Distributor) bool {
	fmt.Println("Enter Exclude Location")
	distributorLocation, ok := getLocationFromUser()
	if !ok {
		return false
	}
	if checkDuplicate(distributor, distributorLocation) {
		fmt.Println("Duplicate location")
		return false
	}
	countryIncluded := false
	stateIncluded := false
	isStateNull := false
	for _, include := range distributor.Include {
		if include.Country == distributorLocation.Country {
			countryIncluded = true
			if include.State == "" {
				isStateNull = true
			}
			if include.State == distributorLocation.State {
				stateIncluded = true
				if include.City == distributorLocation.City {
					fmt.Println("Cannot exclude location which is already included")
					return false
				}
			}
		}

	}
	if !countryIncluded {
		fmt.Println("Cannot exclude country which is not included")
		return false
	}
	if isStateNull {
		distributor.Exclude = append(distributor.Exclude, distributorLocation)
		fmt.Print("Exclude location added successfully\n")
		return true
	}
	if !stateIncluded {
		fmt.Println("Cannot exclude state which is not included")
		return false
	}
	if distributor.Parent == nil {
		distributor.Exclude = append(distributor.Exclude, distributorLocation)
		fmt.Print("Exclude location added successfully\n")
		return true
	}

	return false
}

func AddDistributor() {
	distributorName := ""
	fmt.Println("\n Enter the name of the distributor")
	fmt.Scanln(&distributorName)
	distributorName = strings.TrimSpace(distributorName)
	if !checkDistributorName(distributorName) {
		log.Println("Invalid distributor name")
		return
	}
	distributor := models.Distributor{Name: distributorName, Parent: nil}
	choice := "y"
	fmt.Println("--Adding Permissions Include---")

	for choice == "y" {
		addIncludePermission(&distributor)
		fmt.Print("Add more Include(y/n)? ")
		fmt.Scanln(&choice)
	}
	choice = "y"
	for choice == "y" {
		fmt.Print("Add Exclude (y/n)? ")
		fmt.Scanln(&choice)
		if choice == "y" {
			addExcludePermission(&distributor)
		}
	}

	DistributorMap[distributorName] = distributor
	fmt.Println("Distributor added successfully")
}

func AddSubDistributor() {
	subDistributorName := ""
	fmt.Println("\n Enter the name of the Subdistributor")
	fmt.Scanln(&subDistributorName)
	subDistributorName = strings.TrimSpace(subDistributorName)
	if !checkDistributorName(subDistributorName) {
		log.Println("Invalid distributor name")
		return
	}
	distributorname := ""
	fmt.Println("\n Enter the name of the parent Distributor")
	fmt.Scanln(&distributorname)
	distributorname = strings.TrimSpace(distributorname)
	parentdistributor, ok := DistributorMap[distributorname]
	if !ok {
		fmt.Println("Distributor does not exist")
		return
	}
	distributor := models.Distributor{Name: subDistributorName, Parent: &parentdistributor}
	choice := "y"
	fmt.Println("--Adding Permissions Include---")

	for choice == "y" {
		addIncludePermission(&distributor)
		fmt.Print("Add more Include(y/n)? ")
		fmt.Scanln(&choice)
	}
	choice = "y"
	for choice == "y" {
		fmt.Print("Add Exclude (y/n)? ")
		fmt.Scanln(&choice)
		if choice == "y" {
			addExcludePermission(&distributor)
		}
	}

	DistributorMap[subDistributorName] = distributor
	fmt.Println("Sub-Distributor added successfully")

}

func ViewDistributors() {
	fmt.Println("Distributors:")
	for _, distributor := range DistributorMap {
		fmt.Println(distributor.Name)
		fmt.Println("Include:", distributor.Include)
		fmt.Println("Exclude:", distributor.Exclude)
		isSubDistributor := (distributor.Parent != nil)
		fmt.Println("isSubDistributor:", isSubDistributor)
		if isSubDistributor {
			fmt.Println("Parent:", distributor.Parent.Name)
		}
		fmt.Println("--------------------------------------------------\n")
	}
}

func CheckDistributorByLocation() {
	fmt.Println("Enter Distributor Name")
	var distributorName string
	fmt.Scanln(&distributorName)
	distributorName = strings.TrimSpace(distributorName)
	distributor, ok := DistributorMap[distributorName]
	if !ok {
		fmt.Println("Distributor does not exist")
		return
	}
	fmt.Println("Enter Location")
	distributorLocation, ok := getLocationFromUser()
	if !ok {
		return
	}
	isAuthorized := false
	for _, include := range distributor.Include {
		if include.Country == distributorLocation.Country {
			if include.State == distributorLocation.State || include.State == "" {
				if include.City == distributorLocation.City || include.City == "" {
					isAuthorized = true
					break
				}
			}
		}
	}
	if isAuthorized {
		for _, exclude := range distributor.Exclude {
			if exclude.Country == distributorLocation.Country {
				if exclude.State == distributorLocation.State || exclude.State == "" {
					if exclude.City == distributorLocation.City || exclude.City == "" {
						isAuthorized = false
						break
					}
				}
			}
		}
	}
	fmt.Println("Is ", distributorName, "Distributor Authorized in this location:", isAuthorized)
}
