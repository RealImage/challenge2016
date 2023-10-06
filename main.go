package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Permission struct {
	Country string
	State   string
	City    string
}

type Distributor struct {
	Name    string
	Include Permission
	Exclude Permission
	Parent  *Distributor
}

var distributors = make(map[string]*Distributor)

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

func filterPermissionsBasedOnParent(parentDistributorSelect *widget.Select, locationData map[string]map[string]map[string]bool, includeCountrySelect *widget.Select, includeStateSelect *widget.Select, includeCitySelect *widget.Select, countries []string) {
	parentDistributor := distributors[parentDistributorSelect.Selected]
	if parentDistributor == nil {
		return
	}

	// For Country
	if parentDistributor.Include.Country != "" {
		includeCountrySelect.Options = []string{parentDistributor.Include.Country}
	} else {
		includeCountrySelect.Options = countries
	}
	includeCountrySelect.Refresh()

	// For State
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

	// For City
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

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Distributor Manager")

	locationData, err := LoadCSVData("cities.csv")
	if err != nil {
		panic(err)
	}

	distributorNameEntry := widget.NewEntry()
	distributorNameEntry.SetPlaceHolder("Enter the name of the distributor")

	includeCountrySelect := widget.NewSelect(nil, nil)
	includeStateSelect := widget.NewSelect(nil, nil)
	includeCitySelect := widget.NewSelect(nil, nil)
	excludeCountrySelect := widget.NewSelect(nil, nil)
	excludeStateSelect := widget.NewSelect(nil, nil)
	excludeCitySelect := widget.NewSelect(nil, nil)

	countries := []string{"Select Country"}
	for country := range locationData {
		countries = append(countries, country)
	}
	sort.Strings(countries)

	updateDropdowns := func(countrySelect, stateSelect, citySelect *widget.Select) {
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

	includeCountrySelect.OnChanged = func(string) { updateDropdowns(includeCountrySelect, includeStateSelect, includeCitySelect) }
	includeStateSelect.OnChanged = func(string) { updateDropdowns(includeCountrySelect, includeStateSelect, includeCitySelect) }
	excludeCountrySelect.OnChanged = func(string) { updateDropdowns(excludeCountrySelect, excludeStateSelect, excludeCitySelect) }
	excludeStateSelect.OnChanged = func(string) { updateDropdowns(excludeCountrySelect, excludeStateSelect, excludeCitySelect) }

	includeCountrySelect.Options = countries
	excludeCountrySelect.Options = countries

	parentDistributorSelect := widget.NewSelect(nil, nil)

	updateDistributorList := func() {
		distributorNames := []string{"Select Parent"}
		for name := range distributors {
			distributorNames = append(distributorNames, name)
		}
		sort.Strings(distributorNames)
		parentDistributorSelect.Options = distributorNames
		parentDistributorSelect.Refresh()
	}

	checkDistributorSelect := widget.NewSelect([]string{"Select Distributor"}, nil)

	parentDistributorSelect.OnChanged = func(string) {
		includeCountrySelect.Selected = "Select Country"
		includeStateSelect.Selected = "Select State"
		includeCitySelect.Selected = "Select City"

		includeCountrySelect.Refresh()
		includeStateSelect.Refresh()
		includeCitySelect.Refresh()

		filterPermissionsBasedOnParent(parentDistributorSelect, locationData, includeCountrySelect, includeStateSelect, includeCitySelect, countries)

	}

	updateCheckDistributorList := func() {
		distributorNames := []string{"Select Distributor"}
		for name := range distributors {
			distributorNames = append(distributorNames, name)
		}
		sort.Strings(distributorNames)
		checkDistributorSelect.Options = distributorNames
		checkDistributorSelect.Refresh()
	}

	saveButton := widget.NewButton("Save Distributor", func() {
		distributor := &Distributor{
			Name: distributorNameEntry.Text,
			Include: Permission{
				Country: "",
				State:   "",
				City:    "",
			},
			Exclude: Permission{
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
		updateDistributorList()
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

		updateDistributorList()
		updateCheckDistributorList()
	})

	checkPermissionLabel := widget.NewLabel("Check Distributor Permission")
	checkCountrySelect := widget.NewSelect(countries, nil)
	checkStateSelect := widget.NewSelect([]string{"Select State"}, nil)
	checkCitySelect := widget.NewSelect([]string{"Select City"}, nil)
	permissionResultLabel := widget.NewLabel("")

	checkCountrySelect.OnChanged = func(string) { updateDropdowns(checkCountrySelect, checkStateSelect, checkCitySelect) }
	checkStateSelect.OnChanged = func(string) { updateDropdowns(checkCountrySelect, checkStateSelect, checkCitySelect) }

	checkPermissionButton := widget.NewButton("Check Permission", func() {
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
	})

	content := container.NewVBox(
		widget.NewLabel("Parent Distributor:"),
		parentDistributorSelect,
		distributorNameEntry,
		widget.NewLabel("Include:"),
		container.NewHBox(includeCountrySelect, includeStateSelect, includeCitySelect),
		widget.NewLabel("Exclude:"),
		container.NewHBox(excludeCountrySelect, excludeStateSelect, excludeCitySelect),
		saveButton,
		checkPermissionLabel,
		checkDistributorSelect,
		container.NewHBox(checkCountrySelect, checkStateSelect, checkCitySelect),
		checkPermissionButton,
		permissionResultLabel,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
