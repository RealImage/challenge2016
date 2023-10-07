package main

import (
	"realImage/utils"
	"sort"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Distributor Manager")

	locationData, err := utils.LoadCSVData("cities.csv")
	if err != nil {
		panic(err)
	}

	distributorNameEntry := utils.SetupDistributorNameEntry()
	includeCountrySelect, includeStateSelect, includeCitySelect := utils.SetupCountryStateCitySelectors()
	excludeCountrySelect, excludeStateSelect, excludeCitySelect := utils.SetupCountryStateCitySelectors()

	countries := []string{"Select Country"}
	for country := range locationData {
		countries = append(countries, country)
	}
	sort.Strings(countries)

	includeCountrySelect.OnChanged = func(string) {
		utils.UpdateDropdowns(includeCountrySelect, includeStateSelect, includeCitySelect, locationData)
	}
	includeStateSelect.OnChanged = func(string) {
		utils.UpdateDropdowns(includeCountrySelect, includeStateSelect, includeCitySelect, locationData)
	}
	excludeCountrySelect.OnChanged = func(string) {
		utils.UpdateDropdowns(excludeCountrySelect, excludeStateSelect, excludeCitySelect, locationData)
	}
	excludeStateSelect.OnChanged = func(string) {
		utils.UpdateDropdowns(excludeCountrySelect, excludeStateSelect, excludeCitySelect, locationData)
	}

	includeCountrySelect.Options = countries
	excludeCountrySelect.Options = countries

	parentDistributorSelect := widget.NewSelect(nil, nil)

	utils.UpdateDistributorList(parentDistributorSelect)
	parentDistributorSelect.OnChanged = func(string) {
		includeCountrySelect.Selected = "Select Country"
		includeStateSelect.Selected = "Select State"
		includeCitySelect.Selected = "Select City"

		includeCountrySelect.Refresh()
		includeStateSelect.Refresh()
		includeCitySelect.Refresh()

		utils.FilterPermissionsBasedOnParent(parentDistributorSelect, locationData, includeCountrySelect, includeStateSelect, includeCitySelect, countries)
	}

	checkDistributorSelect := widget.NewSelect([]string{"Select Distributor"}, nil)
	utils.UpdateCheckDistributorList(checkDistributorSelect)

	saveButton := widget.NewButton("Save Distributor", func() {
		utils.OnSaveDistributor(distributorNameEntry, includeCountrySelect, includeStateSelect, includeCitySelect, excludeCountrySelect, excludeStateSelect, excludeCitySelect, parentDistributorSelect, checkDistributorSelect)
	})

	checkPermissionLabel := widget.NewLabel("Check Distributor Permission")
	checkCountrySelect := widget.NewSelect(countries, nil)
	checkStateSelect := widget.NewSelect([]string{"Select State"}, nil)
	checkCitySelect := widget.NewSelect([]string{"Select City"}, nil)
	permissionResultLabel := widget.NewLabel("")

	checkCountrySelect.OnChanged = func(string) {
		utils.UpdateDropdowns(checkCountrySelect, checkStateSelect, checkCitySelect, locationData)
	}
	checkStateSelect.OnChanged = func(string) {
		utils.UpdateDropdowns(checkCountrySelect, checkStateSelect, checkCitySelect, locationData)
	}

	checkPermissionButton := widget.NewButton("Check Permission", func() {
		utils.OnCheckPermission(checkDistributorSelect, checkCountrySelect, checkStateSelect, checkCitySelect, permissionResultLabel)
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
