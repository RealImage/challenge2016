package controllers

import (
	"challenge2016/helpers"
	"log"
	"challenge2016/viewmodels"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"fmt"
)

type DistributorController struct {
	BaseController
}


/*Func to display the list page*/
func (c *DistributorController) ListDistributor() {


	view := viewmodels.ListDistributorVM{}
	allFiles, err := ioutil.ReadDir("./datafiles/distributors/")
	if err != nil {
		log.Fatal(err)
	}

	var fileNameSlice []string

	for _, file := range allFiles {
		fileName := file.Name()
		if fileName != "readme.txt" {
			last := len(fileName) - len(".csv")
			fileName = fileName[:last]
			fileNameSlice = append(fileNameSlice, fileName)
		}
	}

	view.List = fileNameSlice
	view.PageTitle = "List"
	c.Data["vm"] = view

	c.TplName = "templates/list-distributor.html"




}


/*Func to display the new page*/
func (c *DistributorController) NewDistributor() {

	r := c.Ctx.Request
	if(r.Method == "POST"){ // Add new Distributor
		selectedCities := c.GetStrings("selectedCities")
		name := c.GetString("name")
		mode := c.GetString("mode")

		var buffer bytes.Buffer
		buffer.WriteString(mode)

		for i := 0; i < len(selectedCities); i++ {
			if i < len(selectedCities) {
				buffer.WriteString("\n")
			}
			buffer.WriteString(selectedCities[i])
		}

		// Write the details to file
		var fileLocation bytes.Buffer
		fileLocation.WriteString("./datafiles/distributors/")
		fileLocation.WriteString(name)
		fileLocation.WriteString(".csv")

		helpers.DataToFile(fileLocation.String(), buffer.String())

		c.Ctx.ResponseWriter.Write([]byte("true"))


	} else { // Display the Add page

		// Collecting all places from file
		allCities, err := helpers.DataFromFile("./datafiles/data/cities.csv")
		if err != nil {
			log.Println(err)
		}
		view := viewmodels.NewDistributorVM{}
		view.PageTitle = "New Distributor"
		view.AllCities = allCities



		// Collecting Details for add page

		var uniqueCountries, uniqueProvinces [][]string

		for i := 0; i < len(allCities); i++ {
			if !helpers.StringInSlice(allCities[i][2], uniqueCountries) {
				countrySlice := []string{allCities[i][2],allCities[i][5]}
				uniqueCountries = append(uniqueCountries, countrySlice)
			}

			if !helpers.StringInSlice(allCities[i][1], uniqueProvinces) {
				countrySlice := []string{allCities[i][2],allCities[i][5], allCities[i][1],allCities[i][4]}
				uniqueProvinces = append(uniqueProvinces, countrySlice)
			}

		}


		// Reading existing distributor details

		allFiles, err := ioutil.ReadDir("./datafiles/distributors/")
		if err != nil {
			log.Fatal(err)
		}

		var fileNameSlice []string

		for _, file := range allFiles {
			fileName := file.Name()
			if fileName != "readme.txt" {
				last := len(fileName) - len(".csv")
				fileName = fileName[:last]
				fileNameSlice = append(fileNameSlice, fileName)
			}
		}
		view.DistributorNames = fileNameSlice

		distributorCitiesMap := make(map[string][][]string)
		distributorCountriesMap := make(map[string][][]string)
		for k := 0; k < len(fileNameSlice); k++ {
			var fileLocation bytes.Buffer
			fileLocation.WriteString("./datafiles/distributors/")
			fileLocation.WriteString(fileNameSlice[k])
			fileLocation.WriteString(".csv")
			distributorCities, err := helpers.DataFromFile(fileLocation.String())
			if err != nil {
				log.Println(err)
			}

			var tempDistributorCities [][]string

			var distributorCountries [][]string

			for i := 0; i < len(allCities); i++ {
				for j := 0; j < len(distributorCities); j++ {
					if len(distributorCities[j]) == 3 {

							if distributorCities[j][0] == allCities[i][0] && distributorCities[j][1] == allCities[i][1] && distributorCities[j][2] == allCities[i][2] {
								tempDistributorCities = append(tempDistributorCities, allCities[i])
								var tempDistributorCountry []string
								if len(distributorCountries) == 0 {
									tempDistributorCountry = append(tempDistributorCountry, allCities[i][2])
									tempDistributorCountry = append(tempDistributorCountry, allCities[i][5])
									distributorCountries = append(distributorCountries, tempDistributorCountry)
								} else {
									var flag = 0
									for k := 0; k < len(distributorCountries); k++ {
										if distributorCountries[k][0] == allCities[i][2] {
											flag = 1
										}
									}
									if flag == 0 {
										tempDistributorCountry = append(tempDistributorCountry, allCities[i][2])
										tempDistributorCountry = append(tempDistributorCountry, allCities[i][5])
										distributorCountries = append(distributorCountries, tempDistributorCountry)
									}
								}
							}

					} else if  len(distributorCities[j]) == 2 {

							if distributorCities[j][0] == allCities[i][1] && distributorCities[j][1] == allCities[i][2] {
								tempDistributorCities = append(tempDistributorCities, allCities[i])
								var tempDistributorCountry []string
								if len(distributorCountries) == 0 {
									tempDistributorCountry = append(tempDistributorCountry, allCities[i][2])
									tempDistributorCountry = append(tempDistributorCountry, allCities[i][5])
									distributorCountries = append(distributorCountries, tempDistributorCountry)
								} else {
									var flag = 0
									for k := 0; k < len(distributorCountries); k++ {
										if distributorCountries[k][0] == allCities[i][2] {
											flag = 1
										}
									}
									if flag == 0 {
										tempDistributorCountry = append(tempDistributorCountry, allCities[i][2])
										tempDistributorCountry = append(tempDistributorCountry, allCities[i][5])
										distributorCountries = append(distributorCountries, tempDistributorCountry)
									}
								}
							}

					} else {

							if distributorCities[j][0] == allCities[i][2] {
								tempDistributorCities = append(tempDistributorCities, allCities[i])
								var tempDistributorCountry []string
								if len(distributorCountries) == 0 {
									tempDistributorCountry = append(tempDistributorCountry, allCities[i][2])
									tempDistributorCountry = append(tempDistributorCountry, allCities[i][5])
									distributorCountries = append(distributorCountries, tempDistributorCountry)
								} else {
									var flag = 0
									for k := 0; k < len(distributorCountries); k++ {
										if distributorCountries[k][0] == allCities[i][2] {
											flag = 1
										}
									}
									if flag == 0 {
										tempDistributorCountry = append(tempDistributorCountry, allCities[i][2])
										tempDistributorCountry = append(tempDistributorCountry, allCities[i][5])
										distributorCountries = append(distributorCountries, tempDistributorCountry)
									}
								}
							}

					}
				}

			}

			distributorCitiesMap[fileNameSlice[k]] = tempDistributorCities
			distributorCountriesMap[fileNameSlice[k]] = distributorCountries


		}

		view.DistributorCities = distributorCitiesMap
		view.DistributorCountries = distributorCountriesMap



		view.UniqueCountries = uniqueCountries
		view.UniqueProvinces = uniqueProvinces

		c.Data["vm"] = view

		c.TplName = "templates/new-distributor.html"
	}



}

/*Filter distributor provinces*/
func (c *DistributorController) GetDistributorProvinces() {
	allCities, err := helpers.DataFromFile("./datafiles/data/cities.csv")
	if err != nil {
		log.Println(err)
	}
	selectedCountries := c.GetStrings("selectedCountries")

	superDistributor := c.GetString("superDistributor")

	distributorCities, err := helpers.DataFromFile("./datafiles/distributors/" + superDistributor + ".csv")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(distributorCities)

	var tempUniqueProvinces []string
	var uniqueProvinces [][]string
	for i := 0; i < len(selectedCountries); i++ {
		for j := 0; j < len(allCities); j++ {
		// in, india. kl, kerala
			if selectedCountries[i] == allCities[j][2] && len(distributorCities[0]) == 1 {
				tempProvince := allCities[j][2] + "," + allCities[j][1]
				if !helpers.StringInSingleSlice(tempProvince, tempUniqueProvinces) {
					provinceSlice := []string{allCities[j][2], allCities[j][5], allCities[j][1], allCities[j][4]}
					uniqueProvinces = append(uniqueProvinces, provinceSlice)
					tempUniqueProvinces = append(tempUniqueProvinces, tempProvince)
				}
			} else if selectedCountries[i] == allCities[j][2] && len(distributorCities[0]) == 2 {
				for k := 0; k < len(distributorCities); k++ {
					if distributorCities[k][0] == allCities[j][1] && distributorCities[k][1] == allCities[j][2] {
						tempProvince := allCities[j][2] + "," + allCities[j][1]
						if !helpers.StringInSingleSlice(tempProvince, tempUniqueProvinces) {
							provinceSlice := []string{allCities[j][2], allCities[j][5], allCities[j][1], allCities[j][4]}
							uniqueProvinces = append(uniqueProvinces, provinceSlice)
							tempUniqueProvinces = append(tempUniqueProvinces, tempProvince)
						}
					}
				}
			} else if selectedCountries[i] == allCities[j][2] && len(distributorCities[0]) == 3 {
				for k := 0; k < len(distributorCities); k++ {
					if distributorCities[k][0] == allCities[j][0] && distributorCities[k][1] == allCities[j][1]  && distributorCities[k][2] == allCities[j][2] {
						tempProvince := allCities[j][2] + "," + allCities[j][1]
						if !helpers.StringInSingleSlice(tempProvince, tempUniqueProvinces) {
							provinceSlice := []string{allCities[j][2], allCities[j][5], allCities[j][1], allCities[j][4]}
							uniqueProvinces = append(uniqueProvinces, provinceSlice)
							tempUniqueProvinces = append(tempUniqueProvinces, tempProvince)
						}
					}
				}
			}

		}
	}

	toClient, _ := json.Marshal(uniqueProvinces)
	c.Ctx.ResponseWriter.Write(toClient)
}

/*Filter distributor cities*/
func (c *DistributorController) GetDistributorCities() {
	allCities, err := helpers.DataFromFile("./datafiles/data/cities.csv")
	if err != nil {
		log.Println(err)
	}

	codes := c.GetStrings("selectedProvinces")
	selectedProvinces := helpers.DelimitorSplitor(codes)

	superDistributor := c.GetString("superDistributor")
	distributorCities, err := helpers.DataFromFile("./datafiles/distributors/" + superDistributor + ".csv")
	if err != nil {
		log.Println(err)
	}




	var tempUniqueCities []string
	var uniqueCities [][]string

	for i := 0; i < len(selectedProvinces); i++ {
		for j := 0; j < len(allCities); j++ {
			// in, india. kl, kerala
			if selectedProvinces[i][0] == allCities[j][2] && selectedProvinces[i][1] == allCities[j][1] && len(distributorCities[0]) == 1 {
				tempProvince := allCities[j][2] + "," + allCities[j][1] + "," + allCities[j][0]
				if !helpers.StringInSingleSlice(tempProvince, tempUniqueCities) {
					//provinceSlice := []string{allCities[j][2], allCities[j][5], allCities[j][1], allCities[j][4]}
					uniqueCities = append(uniqueCities, allCities[j])
					tempUniqueCities = append(tempUniqueCities, tempProvince)
				}
			} else if selectedProvinces[i][0] == allCities[j][2] && selectedProvinces[i][1] == allCities[j][1] && len(distributorCities[0]) == 2 {
				for k := 0; k < len(distributorCities); k++ {
					if distributorCities[k][0] == allCities[j][1] && distributorCities[k][1] == allCities[j][2] {
						tempProvince := allCities[j][2] + "," + allCities[j][1] + "," + allCities[j][0]
						if !helpers.StringInSingleSlice(tempProvince, tempUniqueCities) {
							//provinceSlice := []string{allCities[j][2], allCities[j][5], allCities[j][1], allCities[j][4]}
							uniqueCities = append(uniqueCities, allCities[j])
							tempUniqueCities = append(tempUniqueCities, tempProvince)
						}
					}
				}
			} else if selectedProvinces[i][0] == allCities[j][2] && selectedProvinces[i][1] == allCities[j][1] && len(distributorCities[0]) == 3 {
				for k := 0; k < len(distributorCities); k++ {
					if distributorCities[k][0] == allCities[j][0] && distributorCities[k][1] == allCities[j][1]  && distributorCities[k][2] == allCities[j][2] {
						tempProvince := allCities[j][2] + "," + allCities[j][1] + "," + allCities[j][0]
						if !helpers.StringInSingleSlice(tempProvince, tempUniqueCities) {
							//provinceSlice := []string{allCities[j][2], allCities[j][5], allCities[j][1], allCities[j][4]}
							uniqueCities = append(uniqueCities, allCities[j])
							tempUniqueCities = append(tempUniqueCities, tempProvince)
						}
					}
				}
			}

		}
	}

	toClient, _ := json.Marshal(uniqueCities)
	c.Ctx.ResponseWriter.Write(toClient)
}

/*Function for the view page*/
func (c *DistributorController) ViewDistributor() {

	allCities, err := helpers.DataFromFile("./datafiles/data/cities.csv")
	if err != nil {
		log.Println(err)
	}

	superDistributor := c.GetString("name")

	distributorCities, err := helpers.DataFromFile("./datafiles/distributors/" + superDistributor + ".csv")
	if err != nil {
		log.Println(err)
	}

	var distributorPlaces [][]string
	for i := 0; i < len(distributorCities); i++ {
		for j := 0; j < len(allCities); j++ {
			// in, india. kl, kerala
			if len(distributorCities[0]) == 1 && distributorCities[i][0] == allCities[j][2]{
				distributorPlaces = append(distributorPlaces, allCities[j])
			} else if len(distributorCities[0]) == 2 && distributorCities[i][0] == allCities[j][1] && distributorCities[i][1] == allCities[j][2]{
				distributorPlaces = append(distributorPlaces, allCities[j])
			} else if len(distributorCities[0]) == 3 && distributorCities[i][0] == allCities[j][0] && distributorCities[i][1] == allCities[j][1] && distributorCities[i][2] == allCities[j][2] {
				distributorPlaces = append(distributorPlaces, allCities[j])
			}

		}
	}


	view := viewmodels.ViewDistributorVM{}
	view.PageTitle = "Distributor"
	view.AllCities = distributorPlaces
	c.Data["vm"] = view

	c.TplName = "templates/view-distributor.html"
}

