package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"realImage/models"
	"sort"

	"fyne.io/fyne/v2/widget"
)

var distributors = make(map[string]*models.Distributor)

// To Load csv file
func LoadCSVData(filename string) (map[string]map[string]map[string]bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	data := make(map[string]map[string]map[string]bool)
	for _, record := range records[1:] {
		country := record[5]
		state := record[4]
		city := record[3]

		if _, exists := data[country]; !exists {
			data[country] = make(map[string]map[string]bool)
		}
		if _, exists := data[country][state]; !exists {
			data[country][state] = make(map[string]bool)
		}

		data[country][state][city] = true
	}
	return data, nil
}

// If parent distrubutor selected, given parent specific region permision
func FilterPermissionsBasedOnParent(parentDistributorSelect *widget.Select, locationData map[string]map[string]map[string]bool, includeCountrySelect *widget.Select, includeStateSelect *widget.Select, includeCitySelect *widget.Select, countries []string) {
	parentDistributor := distributors[parentDistributorSelect.Selected]
	if parentDistributor == nil {
		return
	}

	if parentDistributor.Include.Country != "" {
		includeCountrySelect.Options = []string{parentDistributor.Include.Country}
	} else {
		includeCountrySelect.Options = countries
	}
	includeCountrySelect.Refresh()

	if parentDistributor.Include.State != "" {
		includeStateSelect.Options = []string{parentDistributor.Include.State}
	} else if parentDistributor.Include.Country != "" {
		states := []string{"Select State"}
		for state := range locationData[parentDistributor.Include.Country] {
			states = append(states, state)
		}
		sort.Strings(states)
		includeStateSelect.Options = states
	} else {
		includeStateSelect.Options = []string{"Select State"}
	}
	includeStateSelect.Refresh()

	if parentDistributor.Include.City != "" {
		includeCitySelect.Options = []string{parentDistributor.Include.City}
	} else if parentDistributor.Include.State != "" {
		cities := []string{"Select City"}
		for city := range locationData[parentDistributor.Include.Country][parentDistributor.Include.State] {
			cities = append(cities, city)
		}
		sort.Strings(cities)
		includeCitySelect.Options = cities
	} else {
		includeCitySelect.Options = []string{"Select City"}
	}
	includeCitySelect.Refresh()
}

// Naming distributor
func SetupDistributorNameEntry() *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter the name of the distributor")
	return entry
}

// fetching selector values
func SetupCountryStateCitySelectors() (*widget.Select, *widget.Select, *widget.Select) {
	countrySelect := widget.NewSelect(nil, nil)
	stateSelect := widget.NewSelect(nil, nil)
	citySelect := widget.NewSelect(nil, nil)
	return countrySelect, stateSelect, citySelect
}

// filtering states and city drop down
func UpdateDropdowns(countrySelect, stateSelect, citySelect *widget.Select, locationData map[string]map[string]map[string]bool) {
	if country := countrySelect.Selected; country != "Select Country" {
		states := []string{"Select State"}
		for state := range locationData[country] {
			states = append(states, state)
		}
		sort.Strings(states)
		stateSelect.Options = states
		stateSelect.Refresh()
	}

	if state := stateSelect.Selected; state != "Select State" && countrySelect.Selected != "Select Country" {
		cities := []string{"Select City"}
		for city := range locationData[countrySelect.Selected][state] {
			cities = append(cities, city)
		}
		sort.Strings(cities)
		citySelect.Options = cities
		citySelect.Refresh()
	}
}

// selecting a parent distributor
func UpdateDistributorList(parentDistributorSelect *widget.Select) {
	distributorNames := []string{"Select Parent"}
	for name := range distributors {
		distributorNames = append(distributorNames, name)
	}
	sort.Strings(distributorNames)
	parentDistributorSelect.Options = distributorNames
	parentDistributorSelect.Refresh()
}

// selecting a distributor for check permision drop-down
func UpdateCheckDistributorList(checkDistributorSelect *widget.Select) {
	distributorNames := []string{"Select Distributor"}
	for name := range distributors {
		distributorNames = append(distributorNames, name)
	}
	sort.Strings(distributorNames)
	checkDistributorSelect.Options = distributorNames
	checkDistributorSelect.Refresh()
}

// save distributor logic
func OnSaveDistributor(distributorNameEntry *widget.Entry, includeCountrySelect, includeStateSelect, includeCitySelect, excludeCountrySelect, excludeStateSelect, excludeCitySelect, parentDistributorSelect *widget.Select, checkDistributorSelect *widget.Select) {
	distributor := &models.Distributor{
		Name: distributorNameEntry.Text,
		Include: models.Permission{
			Country: "",
			State:   "",
			City:    "",
		},
		Exclude: models.Permission{
			Country: "",
			State:   "",
			City:    "",
		},
	}

	if includeCountrySelect.Selected != "Select Country" {
		distributor.Include.Country = includeCountrySelect.Selected
	}
	if includeStateSelect.Selected != "Select State" {
		distributor.Include.State = includeStateSelect.Selected
	}
	if includeCitySelect.Selected != "Select City" {
		distributor.Include.City = includeCitySelect.Selected
	}

	if excludeCountrySelect.Selected != "Select Country" {
		distributor.Exclude.Country = excludeCountrySelect.Selected
	}
	if excludeStateSelect.Selected != "Select State" {
		distributor.Exclude.State = excludeStateSelect.Selected
	}
	if excludeCitySelect.Selected != "Select City" {
		distributor.Exclude.City = excludeCitySelect.Selected
	}

	if parent := parentDistributorSelect.Selected; parent != "Select Parent" && distributors[parent] != nil {
		distributor.Parent = distributors[parent]
	}

	distributors[distributor.Name] = distributor
	UpdateDistributorList(parentDistributorSelect)
	fmt.Println("Saved Distributor:", distributor)

	distributorNameEntry.SetText("")
	includeCountrySelect.Selected = "Select Country"
	includeStateSelect.Selected = "Select State"
	includeCitySelect.Selected = "Select City"
	excludeCountrySelect.Selected = "Select Country"
	excludeStateSelect.Selected = "Select State"
	excludeCitySelect.Selected = "Select City"
	parentDistributorSelect.Selected = "Select Parent"

	includeCountrySelect.Refresh()
	includeStateSelect.Refresh()
	includeCitySelect.Refresh()
	excludeCountrySelect.Refresh()
	excludeStateSelect.Refresh()
	excludeCitySelect.Refresh()
	parentDistributorSelect.Refresh()

	UpdateDistributorList(parentDistributorSelect)
	UpdateCheckDistributorList(checkDistributorSelect)
}

// checck permission button logic
func OnCheckPermission(checkDistributorSelect, checkCountrySelect, checkStateSelect, checkCitySelect *widget.Select, permissionResultLabel *widget.Label) {
	selectedDistributor := distributors[checkDistributorSelect.Selected]
	if selectedDistributor == nil {
		permissionResultLabel.SetText("Invalid Distributor")
		return
	}

	selectedCountry := checkCountrySelect.Selected
	selectedState := checkStateSelect.Selected
	selectedCity := checkCitySelect.Selected

	hasIncludePermission := false
	hasExcludePermission := false

	if selectedDistributor.Include.Country == selectedCountry || selectedDistributor.Include.Country == "" {
		if selectedDistributor.Include.State == selectedState || selectedDistributor.Include.State == "" {
			if selectedDistributor.Include.City == selectedCity || selectedDistributor.Include.City == "" {
				hasIncludePermission = true
			}
		}
	}

	if selectedDistributor.Exclude.Country == selectedCountry {
		if selectedDistributor.Exclude.State == "" || selectedDistributor.Exclude.State == selectedState {
			if selectedDistributor.Exclude.City == "" || selectedDistributor.Exclude.City == selectedCity {
				hasExcludePermission = true
			}
		}
	} else if selectedDistributor.Exclude.Country == "" && selectedDistributor.Exclude.State == selectedState {
		if selectedDistributor.Exclude.City == "" || selectedDistributor.Exclude.City == selectedCity {
			hasExcludePermission = true
		}
	}

	if hasIncludePermission && !hasExcludePermission {
		permissionResultLabel.SetText("Has Access")
	} else {
		permissionResultLabel.SetText("Does Not Have Access")
	}

	checkCountrySelect.Selected = "Select Country"
	checkStateSelect.Selected = "Select State"
	checkCitySelect.Selected = "Select City"

	checkCountrySelect.Refresh()
	checkStateSelect.Refresh()
	checkCitySelect.Refresh()
}
