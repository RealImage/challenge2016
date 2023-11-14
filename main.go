package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

type Distributor struct {
	Name     string
	Includes []int
	Excludes []int
	Parent   []int
}

type Location struct {
	CityCode     string
	ProvinceCode string
	CountryCode  string
	CityName     string
	ProvinceName string
	CountryName  string
}

var distributors []Distributor
var locations []Location

func main() {
	readLocationFromCSV("cities.csv")
	for {
		_, action := showPrompt("Select an action", []string{"view", "add", "exit"})

		switch action {
		case "view":
			fmt.Println()
			searchDistributor()
		case "add":
			fmt.Println()
			promptMenu := promptui.Prompt{
				Label: "Enter distributor name",
			}

			distributorName, err := promptMenu.Run()
			if err != nil {
				fmt.Printf("Unable to store distributor %v\n", err)
				os.Exit(1)
			}
			updateDistributorConfig(distributorName)
		case "exit":
			fmt.Println()
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func readLocationFromCSV(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Unable to read cities data %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Unable to read cities data %v\n", err)
		os.Exit(1)
	}

	for _, record := range records[1:] {
		if len(record) > 0 {
			location := Location{
				CityCode:     record[0],
				ProvinceCode: record[1],
				CountryCode:  record[2],
				CityName:     record[3],
				ProvinceName: record[4],
				CountryName:  record[5],
			}
			locations = append(locations, location)
		}
	}
}

func showPrompt(label string, items []string) (int, string) {
	searcher := func(input string, index int) bool {
		item := items[index]
		return index != -1 && (input == "" || containsIgnoreCase(item, input))
	}
	prompt := promptui.Select{
		Label:    label,
		Items:    items,
		Size:     50,
		Searcher: searcher,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . | bold }}",
			Active:   "> {{ . | cyan }}",
			Inactive: "  {{ . | white }}",
		},
	}
	index, result, err := prompt.Run()
	fmt.Println()
	if err != nil {
		fmt.Println("Prompt failed:", err)
		return -1, ""
	}
	return index, result
}

func updateDistributorConfig(distributorName string) {
	var includes []int
	var excludes []int
	var parents []int
	for {
		_, savePrompt := showPrompt("Select distributor config to update", []string{"includes", "excludes", "parents", "confirm"})
		switch savePrompt {
		case "includes":
			fmt.Println()
			fmt.Println("Select locations to be included")
			index := showLocationPrompt()
			includes = append(includes, index)
		case "excludes":
			fmt.Println()
			fmt.Println("Select locations to be excluded")
			index := showLocationPrompt()
			excludes = append(excludes, index)
		case "parents":
			fmt.Println()
			fmt.Println("Select distrubutor for parents")
			index := showDistributorPrompt()
			parents = append(parents, index)
		case "confirm":
			fmt.Println()
			fmt.Println("Confirmed changes")
			parents, includes, excludes = updateParentData(parents, includes, excludes)
			distributor := Distributor{
				Name:     distributorName,
				Includes: includes,
				Excludes: excludes,
				Parent:   parents,
			}
			distributors = append(distributors, distributor)
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func updateParentData(parents []int, includes []int, excludes []int) ([]int, []int, []int) {
	if len(parents) > 0 {
		for _, p := range parents {
			parentDist := &distributors[p]
			includes = append(includes, parentDist.Includes...)
			excludes = append(excludes, parentDist.Excludes...)
			if len(parentDist.Parent) > 0 {
				parents, includes, excludes = updateParentData(parentDist.Parent, includes, excludes)
			}
		}
	}
	return parents, includes, excludes
}

func showLocationPrompt() int {
	var locationNames []string
	for _, loc := range locations {
		locationNames = append(locationNames, loc.CityName+", "+loc.ProvinceName+", "+loc.CountryName)
	}
	index, _ := showPrompt("Select location", locationNames)
	return index
}

func showDistributorPrompt() int {
	var distributorNames []string
	for _, distro := range distributors {
		distributorNames = append(distributorNames, distro.Name)
	}
	index, _ := showPrompt("Select a distributor", distributorNames)
	return index
}

func containsIgnoreCase(s, substr string) bool {
	return len(s) >= len(substr) && (s[:len(substr)] == substr || containsIgnoreCase(s[1:], substr))
}

func searchDistributor() {
	fmt.Println("Check if a location is serviceable.")
	distributorIndx := showDistributorPrompt()
	location := showLocationPrompt()
	distributor := distributors[distributorIndx]
	isPresent := checkIfPresent(location, distributor.Includes)
	if isPresent {
		fmt.Println("Location is Serviceable")
	} else {
		fmt.Println("Non Serviceable")
	}
}

func checkIfPresent(value int, includes []int) bool {
	for _, v := range includes {
		if v == value {
			return true
		}
	}
	return false
}
