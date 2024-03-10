package main

import (
	"bufio"
	"fmt"
	"github.com/RealImage/challenge2016/internal"
	"os"
	"strings"
)

func init() {
	// Initialize internal data stores
	internal.NewRegionDB()
	internal.NewFilmDB()
	internal.NewDistributorDB()
	internal.NewAuthorizationDB()
}

func main() {

	fmt.Println("AK's Welcome to Film Distribution Management System")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nOptions:")
		fmt.Println("1. Add Region Data")
		fmt.Println("2. Add Film")
		fmt.Println("3. Add Distributor")
		fmt.Println("4. Authorize distributor for film in a region")
		fmt.Println("5. Check if distributor has permission to distribute film in a region")
		fmt.Println("6. Exit")
		fmt.Print("Enter option number: ")

		optionStr, _ := reader.ReadString('\n')
		option := strings.TrimSpace(optionStr)

		switch option {
		case "1":
			err := readRegionDataInput(reader)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case "2":
			fmt.Print("Enter film ID(ex. AVTR): ")
			filmID, _ := reader.ReadString('\n')
			filmID = strings.TrimSpace(filmID)
			fmt.Print("Enter film Name(ex. AVATAR): ")
			filmName, _ := reader.ReadString('\n')
			filmName = strings.TrimSpace(filmName)
			err := internal.AddFilm(filmID, filmName)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Film '%s' added successfully.\n", filmName)
		case "3":
			fmt.Print("Enter distributor ID(ex. YRF): ")
			distributorID, _ := reader.ReadString('\n')
			distributorID = strings.TrimSpace(distributorID)
			fmt.Print("Enter distributor name(ex. Yash Raj Films): ")
			distributorName, _ := reader.ReadString('\n')
			distributorName = strings.TrimSpace(distributorName)
			err := internal.AddDistributor(distributorID, distributorName)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Distributor '%s' added successfully.\n", distributorName)
		case "4":
			err := readFilmAuthorizationInput(reader)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Authorization successfully granted")
		case "5":
			err := readPermissionCheckInput(reader)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case "6":
			fmt.Println("Thanks for using AK's film distribution management")
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please enter a valid option number.")
		}
	}
}

// readRegionDataInput reads input for region data
func readRegionDataInput(reader *bufio.Reader) error {
	fmt.Println("\nOptions:")
	fmt.Println("1. Provide URL to read region data from remote CSV file")
	fmt.Println("2. Provide path to read region data from local CSV file")
	fmt.Println("3. Exit")
	fmt.Print("Enter option number: ")

	optionStr, _ := reader.ReadString('\n')
	option := strings.TrimSpace(optionStr)
	switch option {
	case "1":
		fmt.Print("Enter region data URL: ")
		url, _ := reader.ReadString('\n')
		url = strings.TrimSpace(url)
		err := internal.ReadRegionDataFromRemoteCSV(url)
		if err != nil {
			return fmt.Errorf("error reading region data from remote file: %v", err)
		} else {
			fmt.Println("Region data read successfully from remote file.")
		}
	case "2":
		fmt.Print("Enter region data file path: ")
		path, _ := reader.ReadString('\n')
		path = strings.TrimSpace(path)
		err := internal.ReadRegionDataFromLocalCSV(path)
		if err != nil {
			return fmt.Errorf("error reading region data from local file: %v", err)
		} else {
			fmt.Println("Region data read successfully from local file.")
		}
	case "3":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid option. Please enter a valid option number.")
	}
	return nil
}

// readFilmAuthorizationInput reads input for film authorization
func readFilmAuthorizationInput(reader *bufio.Reader) error {
	var filmID, authorityID, agentDistributorID string
	var includedRegionsList, excludedRegionsList []string
	for {
		fmt.Println("Provide Film ID (ex. AVTR)")
		fmt.Println("Enter 0 to go to main menu")
		filmID, _ = reader.ReadString('\n')
		filmID = strings.TrimSpace(filmID)
		if filmID == "0" {
			fmt.Println("returning to main menu...")
			return nil
		}
		if !internal.IsValidFilm(filmID) {
			fmt.Println("Error: Invalid Film ID")
			continue
		}
		break
	}
	for {
		fmt.Println("Provide Authority Distributor ID (ex. YRF or leave empty if directly authorized by film authority)")
		fmt.Println("Enter 0 to go to main menu")
		authorityID, _ = reader.ReadString('\n')
		authorityID = strings.TrimSpace(authorityID)
		if authorityID == "0" {
			fmt.Println("returning to main menu...")
			return nil
		}
		if authorityID != "" && !internal.IsValidDistributor(authorityID) {
			fmt.Println("Error: Invalid Authority Distributor ID")
			continue
		}
		break
	}
	for {
		fmt.Println("Provide Agent Distributor ID(ex. DHRM)")
		fmt.Println("Enter 0 to go to main menu")
		agentDistributorID, _ = reader.ReadString('\n')
		agentDistributorID = strings.TrimSpace(agentDistributorID)
		if agentDistributorID == "0" {
			fmt.Println("returning to main menu...")
			return nil
		}
		if !internal.IsValidDistributor(agentDistributorID) {
			fmt.Println("Error: Invalid Agent Distributor ID")
			continue
		}
		break
	}
	for {
		fmt.Println("Provide Included Regions List(ex. IN,US,PUNCH-JK-IN,TN-IN)")
		fmt.Println("Enter 0 to go to main menu")
		includedRegions, _ := reader.ReadString('\n')
		includedRegions = strings.TrimSpace(includedRegions)
		if includedRegions == "0" {
			fmt.Println("returning to main menu...")
			return nil
		}
		includedRegionsList = strings.Split(includedRegions, ",")
		ok := true
		for _, region := range includedRegionsList {
			if !internal.IsValidRegion(region) {
				fmt.Println("Error: Invalid Included Region: ", region)
				ok = false
				break
			}
		}
		if !ok {
			continue
		}
		break
	}
	for {
		fmt.Println("Provide Excluded Regions List(ex. CN,UK,PUNCH-JK-IN,JK-IN)")
		fmt.Println("Enter 0 to go to main menu")
		excludedRegions, _ := reader.ReadString('\n')
		excludedRegions = strings.TrimSpace(excludedRegions)
		if excludedRegions == "0" {
			fmt.Println("returning to main menu...")
			return nil
		}
		excludedRegionsList = strings.Split(excludedRegions, ",")
		ok := true
		for _, region := range excludedRegionsList {
			if !internal.IsValidRegion(region) {
				fmt.Println("Error: Invalid Excluded Region: ", region)
				ok = false
				break
			}
		}
		if !ok {
			continue
		}
		break
	}
	err := internal.AuthorizeDistributor(filmID, authorityID, agentDistributorID, includedRegionsList, excludedRegionsList)
	if err != nil {
		return fmt.Errorf("failed to authorize distributor '%s': %v", agentDistributorID, err)
	}
	return nil
}

// readPermissionCheckInput reads the input to check if a distributor has permission
func readPermissionCheckInput(reader *bufio.Reader) error {
	var filmID, distributorID, region string
	for {
		fmt.Println("Provide Film ID (ex. AVTR)")
		fmt.Println("Enter 0 to go to main menu")
		filmID, _ = reader.ReadString('\n')
		filmID = strings.TrimSpace(filmID)
		if filmID == "0" {
			fmt.Println("returning to main menu...")
			return nil
		}
		if !internal.IsValidFilm(filmID) {
			fmt.Println("Error: Invalid Film ID")
			continue
		}
		break
	}
	for {
		fmt.Println("Provide Distributor ID (ex. YRF)")
		fmt.Println("Enter 0 to go to main menu")
		distributorID, _ = reader.ReadString('\n')
		distributorID = strings.TrimSpace(distributorID)
		if distributorID == "0" {
			fmt.Println("returning to main menu...")
			return nil
		}
		if !internal.IsValidDistributor(distributorID) {
			fmt.Println("Error: Invalid Distributor ID")
			continue
		}
		break
	}
	for {
		fmt.Println("Provide Region (ex. CN or PUNCH-JK-IN or JK-IN)")
		fmt.Println("Enter 0 to go to main menu")
		region, _ = reader.ReadString('\n')
		region = strings.TrimSpace(region)
		if region == "0" {
			fmt.Println("returning to main menu...")
			return nil
		}
		if !internal.IsValidRegion(region) {
			fmt.Println("Error: Invalid region")
			continue
		}
		if internal.HasPermission(filmID, distributorID, region) {
			fmt.Printf("Distributor '%s' has permission to distribute film '%s' in region '%s'\n", distributorID,
				filmID, region)
			fmt.Println("If you want to check more...")
		} else {
			fmt.Printf("Distributor '%s' doesn't have permission to distribute film '%s' in region '%s'\n", distributorID,
				filmID, region)
			fmt.Println("If you want to check more...")
		}
	}
}
