package main

import (
	"encoding/csv"
	"os"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type LocationData struct {
	Country string
	State   string
	City    string
}

type Permissions struct {
	Include []string
	Exclude []string
}

type Distributor struct {
	Name        string
	Permissions Permissions
	Parent      *Distributor
}

var (
	distributors []*Distributor
	countries    []string
)

func LoadCSVData(filename string) ([]LocationData, error) {
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

	var data []LocationData
	for _, record := range records[1:] {
		data = append(data, LocationData{
			City:    record[3],
			State:   record[4],
			Country: record[5],
		})
	}
	return data, nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func unique(strings []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strings {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func resetFields(
	distributorNameEntry *widget.Entry,
	includeCountrySelect, includeStateSelect, includeCitySelect,
	excludeCountrySelect, excludeStateSelect, excludeCitySelect,
	editDistributorSelect, parentDistributorSelect *widget.Select) {

	distributorNameEntry.SetText("")

	includeCountrySelect.Options = append([]string{"(Select one)"}, unique(countries)...)
	includeCountrySelect.SetSelected("(Select one)")
	includeStateSelect.Options = []string{"(Select one)"}
	includeStateSelect.SetSelected("(Select one)")
	includeCitySelect.Options = []string{"(Select one)"}
	includeCitySelect.SetSelected("(Select one)")

	excludeCountrySelect.Options = append([]string{"(Select one)"}, unique(countries)...)
	excludeCountrySelect.SetSelected("(Select one)")
	excludeStateSelect.Options = []string{"(Select one)"}
	excludeStateSelect.SetSelected("(Select one)")
	excludeCitySelect.Options = []string{"(Select one)"}
	excludeCitySelect.SetSelected("(Select one)")

	editDistributorSelect.Options = append([]string{"Select Distributor to Edit"}, getDistributorNames()...)
	editDistributorSelect.SetSelected("Select Distributor to Edit")

	parentDistributorSelect.Options = append([]string{"(Select one)"}, getDistributorNames()...)
	parentDistributorSelect.SetSelected("(Select one)")

	includeCountrySelect.Refresh()
	includeStateSelect.Refresh()
	includeCitySelect.Refresh()

	excludeCountrySelect.Refresh()
	excludeStateSelect.Refresh()
	excludeCitySelect.Refresh()

	editDistributorSelect.Refresh()
	parentDistributorSelect.Refresh()
}

func getDistributorNames() []string {
	distributorNames := []string{}
	for _, dist := range distributors {
		distributorNames = append(distributorNames, dist.Name)
	}
	return distributorNames
}

func getDistributorByName(name string) *Distributor {
	for _, dist := range distributors {
		if dist.Name == name {
			return dist
		}
	}
	return nil
}

func hasPermission(dist *Distributor, location string) bool {
	parts := strings.Split(location, "-")

	// Constructing the location hierarchy for checking
	var locationsToCheck []string
	for i := 1; i <= len(parts); i++ {
		locationsToCheck = append(locationsToCheck, strings.Join(parts[:i], "-"))
	}

	for _, loc := range locationsToCheck {
		// Checking if the location exists in the distributor's permissions
		if contains(dist.Permissions.Include, loc) {
			return true
		}

		if contains(dist.Permissions.Exclude, loc) {
			return false
		}
	}

	// If permission is not determined and distributor has a parent, check parent's permissions
	if dist.Parent != nil {
		return hasPermission(dist.Parent, location)
	}

	return false
}

func main() {
	data, err := LoadCSVData("cities.csv")
	if err != nil {
		panic(err)
	}

	countries = []string{}
	stateByCountry := make(map[string][]string)
	cityByStateCountry := make(map[string][]string)
	for _, location := range data {
		if !contains(countries, location.Country) {
			countries = append(countries, location.Country)
		}
		stateKey := location.Country
		cityKey := location.State + "-" + location.Country

		stateByCountry[stateKey] = append(stateByCountry[stateKey], location.State)
		cityByStateCountry[cityKey] = append(cityByStateCountry[cityKey], location.City)
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("Distributor Manager")

	distributorNameEntry := widget.NewEntry()
	distributorNameEntry.SetPlaceHolder("Distributor Name")

	includeCountrySelect := widget.NewSelect(unique(countries), nil)
	includeStateSelect := widget.NewSelect(nil, nil)
	includeCitySelect := widget.NewSelect(nil, nil)
	excludeCountrySelect := widget.NewSelect(unique(countries), nil)
	excludeStateSelect := widget.NewSelect(nil, nil)
	excludeCitySelect := widget.NewSelect(nil, nil)

	includeCountrySelect.OnChanged = func(country string) {
		includeStateSelect.Options = unique(stateByCountry[country])
		includeStateSelect.Refresh()
		includeCitySelect.Options = nil
		includeCitySelect.Refresh()
	}

	includeStateSelect.OnChanged = func(state string) {
		key := state + "-" + includeCountrySelect.Selected
		includeCitySelect.Options = unique(cityByStateCountry[key])
		includeCitySelect.Refresh()
	}

	excludeCountrySelect.OnChanged = func(country string) {
		excludeStateSelect.Options = unique(stateByCountry[country])
		excludeStateSelect.Refresh()
		excludeCitySelect.Options = nil
		excludeCitySelect.Refresh()
	}

	excludeStateSelect.OnChanged = func(state string) {
		key := state + "-" + excludeCountrySelect.Selected
		excludeCitySelect.Options = unique(cityByStateCountry[key])
		excludeCitySelect.Refresh()
	}

	parentDistributorSelect := widget.NewSelect(nil, nil)

	editMode := false
	var currentDistributor *Distributor

	editDistributorSelect := widget.NewSelect(nil, func(selected string) {
		for _, dist := range distributors {
			if dist.Name == selected {
				currentDistributor = dist
				distributorNameEntry.SetText(dist.Name)
				includeParts := strings.Split(dist.Permissions.Include[0], "-")
				includeCountrySelect.SetSelected(includeParts[2])
				includeStateSelect.Options = unique(stateByCountry[includeParts[2]])
				includeStateSelect.Refresh()
				includeStateSelect.SetSelected(includeParts[1])
				includeCitySelect.Options = unique(cityByStateCountry[includeParts[1]+"-"+includeParts[2]])
				includeCitySelect.Refresh()
				includeCitySelect.SetSelected(includeParts[0])

				excludeParts := strings.Split(dist.Permissions.Exclude[0], "-")
				excludeCountrySelect.SetSelected(excludeParts[2])
				excludeStateSelect.Options = unique(stateByCountry[excludeParts[2]])
				excludeStateSelect.Refresh()
				excludeStateSelect.SetSelected(excludeParts[1])
				excludeCitySelect.Options = unique(cityByStateCountry[excludeParts[1]+"-"+excludeParts[2]])
				excludeCitySelect.Refresh()
				excludeCitySelect.SetSelected(excludeParts[0])

				editMode = true
				break
			}
		}
	})
	editDistributorSelect.PlaceHolder = "Select Distributor to Edit"

	logDisplay := widget.NewLabel("")

	checkDistributorSelect := widget.NewSelect(nil, nil)
	checkCountrySelect := widget.NewSelect(unique(countries), nil)
	checkStateSelect := widget.NewSelect(nil, nil)
	checkCitySelect := widget.NewSelect(nil, nil)
	resultLabel := widget.NewLabel("")

	checkCountrySelect.OnChanged = func(country string) {
		checkStateSelect.Options = unique(stateByCountry[country])
		checkStateSelect.Refresh()
		checkCitySelect.Options = nil
		checkCitySelect.Refresh()
	}

	checkStateSelect.OnChanged = func(state string) {
		key := state + "-" + checkCountrySelect.Selected
		checkCitySelect.Options = unique(cityByStateCountry[key])
		checkCitySelect.Refresh()
	}

	saveButton := widget.NewButton("Save Distributor", func() {
		if editMode && currentDistributor != nil {
			currentDistributor.Name = distributorNameEntry.Text
			currentDistributor.Permissions.Include = []string{includeCitySelect.Selected + "-" + includeStateSelect.Selected + "-" + includeCountrySelect.Selected}
			currentDistributor.Permissions.Exclude = []string{excludeCitySelect.Selected + "-" + excludeStateSelect.Selected + "-" + excludeCountrySelect.Selected}
		} else {
			newDistributor := &Distributor{
				Name: distributorNameEntry.Text,
				Permissions: Permissions{
					Include: []string{includeCitySelect.Selected + "-" + includeStateSelect.Selected + "-" + includeCountrySelect.Selected},
					Exclude: []string{excludeCitySelect.Selected + "-" + excludeStateSelect.Selected + "-" + excludeCountrySelect.Selected},
				},
			}
			if parentDistributorSelect.Selected != "" {
				for _, dist := range distributors {
					if dist.Name == parentDistributorSelect.Selected {
						newDistributor.Parent = dist
						break
					}
				}
			}
			distributors = append(distributors, newDistributor)
			currentDistributor = newDistributor
		}

		editDistributorSelect.Options = getDistributorNames()
		editDistributorSelect.Refresh()

		parentDistributorSelect.Options = getDistributorNames()
		parentDistributorSelect.Refresh()

		checkDistributorSelect.Options = getDistributorNames()
		checkDistributorSelect.Refresh()

		// Prepareinf the log message
		distributorChain := currentDistributor.Name
		parentDistributor := currentDistributor.Parent
		for parentDistributor != nil {
			distributorChain = parentDistributor.Name + " -> " + distributorChain
			parentDistributor = parentDistributor.Parent
		}

		includeInfo := "Include of " + currentDistributor.Name + ": " + strings.Join(currentDistributor.Permissions.Include, ", ")
		excludeInfo := "Exclude of " + currentDistributor.Name + ": " + strings.Join(currentDistributor.Permissions.Exclude, ", ")

		// Appending the log message to the label
		currentLog := logDisplay.Text
		newLog := distributorChain + "\n" + includeInfo + "\n" + excludeInfo
		if currentLog != "" {
			newLog = currentLog + "\n\n" + newLog
		}
		logDisplay.SetText(newLog)

		// Reseting fields after saving
		resetFields(distributorNameEntry, includeCountrySelect, includeStateSelect, includeCitySelect, excludeCountrySelect, excludeStateSelect, excludeCitySelect, editDistributorSelect, parentDistributorSelect)

		editMode = false
		currentDistributor = nil

		// Clearing selections for parent and edit dropdowns
		parentDistributorSelect.SetSelected("")
		editDistributorSelect.SetSelected("")
	})

	checkButton := widget.NewButton("Check Permissions", func() {
		selectedDistributor := getDistributorByName(checkDistributorSelect.Selected)
		location := checkCitySelect.Selected + "-" + checkStateSelect.Selected + "-" + checkCountrySelect.Selected
		if selectedDistributor != nil {
			if hasPermission(selectedDistributor, location) {
				resultLabel.SetText("YES")
			} else {
				resultLabel.SetText("NO")
			}
		} else {
			resultLabel.SetText("Invalid Distributor")
		}
	})

	content := container.NewVBox(
		editDistributorSelect,
		distributorNameEntry,
		container.NewHBox(widget.NewLabel("Include: "), includeCountrySelect, includeStateSelect, includeCitySelect),
		container.NewHBox(widget.NewLabel("Exclude: "), excludeCountrySelect, excludeStateSelect, excludeCitySelect),
		widget.NewLabel("Parent Distributor:"), parentDistributorSelect,
		saveButton,
		logDisplay,
		widget.NewLabel("Check permissions for Distributor:"),
		checkDistributorSelect,
		container.NewHBox(widget.NewLabel("Location: "), checkCountrySelect, checkStateSelect, checkCitySelect),
		checkButton,
		resultLabel,
	)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func splitLocation(location string) []string {
	return strings.Split(location, "-")
}
