package input

import (
	"bufio"
	"challenge2016/dto" // Importing DTO package for data transfer objects
	"fmt"
	"os"
	"strings"
)

// AskNextQuestion prints the menu and returns the user's choice
func AskNextQuestion() string {
	fmt.Println("Select one of the below choices:")
	fmt.Println("1. Create a new distributor")
	fmt.Println("2. Create a sub distributor")
	fmt.Println("3. Check permission for a distributor")
	fmt.Println("4. View Distributors information")
	fmt.Println("5. Exit the program")

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	return choice
}

// GetDistributorData retrieves data for a new distributor from user input
func GetDistributorData() dto.Distributor {
	var distributor dto.Distributor
	fmt.Print("Enter distributor name: ")
	fmt.Scanln(&distributor.Name)
	fmt.Print("Enter the regions you want to include for this distributor: ")
	includeRegions := GetInputRegions()
	distributor.Include = includeRegions
	fmt.Print("Enter the regions you want to exclude for this distributor: ")
	excludeRegions := GetInputRegions()
	distributor.Exclude = excludeRegions
	return distributor
}

// GetSubDistributorData retrieves data for a new sub-distributor from user input
func GetSubDistributorData() dto.Distributor {
	var distributor dto.Distributor
	fmt.Print("Enter distributor name: ")
	fmt.Scanln(&distributor.Name)
	fmt.Print("Enter the regions you want to include for this distributor: ")
	includeRegions := GetInputRegions()
	distributor.Include = includeRegions
	fmt.Print("Enter the regions you want to exclude for this distributor: ")
	excludeRegions := GetInputRegions()
	distributor.Exclude = excludeRegions
	fmt.Print("Enter the name of the parent distributor: ")
	fmt.Scanln(&distributor.Parent)
	return distributor
}

// GetInputRegions retrieves regions from user input
func GetInputRegions() []string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Split by comma first
	regions := strings.Split(input, ",")

	cleanedRegions := make([]string, 0)

	for _, region := range regions {
		// Check if the region contains a hyphen
		if strings.Contains(region, "-") {
			cleanedRegions = append(cleanedRegions, strings.TrimSpace(strings.ToUpper(region)))
		} else {
			// If no hyphen, split by comma
			subRegions := strings.Split(strings.TrimSpace(region), "-")
			for _, subRegion := range subRegions {
				cleanedRegion := strings.TrimSpace(subRegion)
				if cleanedRegion != "" {
					cleanedRegions = append(cleanedRegions, strings.ToUpper(cleanedRegion))
				}
			}
		}
	}
	return cleanedRegions
}

// GetCheckPermissionData retrieves data for permission checking from user input
func GetCheckPermissionData() dto.CheckPermissionData {
	var data dto.CheckPermissionData
	fmt.Print("Enter distributor name that needs to be checked: ")
	fmt.Scanln(&data.DistributorName)
	fmt.Print("Enter regions that need to be checked: ")
	data.Regions = GetInputRegions()
	return data
}

// CreateNewDistributor creates a new distributor object
func CreateNewDistributor(data dto.Distributor) dto.Distributor {
	return dto.Distributor{
		Name:    strings.ToUpper(data.Name),
		Include: data.Include,
		Exclude: data.Exclude,
		Parent:  strings.ToUpper(data.Parent),
	}
}

// DisplayDistributorInformation displays information about distributors
func DisplayDistributorInformation(distributorInformation []dto.Distributor) {
	fmt.Println("Distributor Information:")
	for _, distributor := range distributorInformation {
		fmt.Printf("Name: %s, Include: %v, Exclude: %v, Parent: %s\n", distributor.Name, distributor.Include, distributor.Exclude, distributor.Parent)
	}
}
