package controllers

import (
	"challenge2016/helpers"
	"log"
	"challenge2016/viewmodels"
	"strings"
	"bytes"
	"io/ioutil"
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

	// Collecting all places from file
	allCities, err := helpers.DataFromFile("./datafiles/data/cities.csv")
	if err != nil {
		log.Println(err)
	}

	r := c.Ctx.Request
	if(r.Method == "POST"){ // Add new Distributor
		selectedCities := c.GetStrings("selectedCities")
		name := c.GetString("name")
		mode := c.GetString("mode")

		fmt.Println(selectedCities)

		var buffer bytes.Buffer
		buffer.WriteString("City Code,Province Code,Country Code,City Name,Province Name,Country Name")

		for i := 0; i < len(selectedCities); i++ {
			for j := 0; j < len(allCities); j++ {
				if mode == "0" {
					if strings.Compare(selectedCities[i], allCities[j][5]) == 0 {
						if i < len(selectedCities) {
							buffer.WriteString("\n")
						}
						for k := 0; k < len(allCities[j]); k++ {
							buffer.WriteString(allCities[j][k])
							if k != len(allCities[j]) - 1 {
								buffer.WriteString(",")
							}
						}
					}
				} else if mode == "1" {
					if strings.Compare(selectedCities[i], allCities[j][4]) == 0 {
						if i < len(selectedCities) {
							buffer.WriteString("\n")
						}
						for k := 0; k < len(allCities[j]); k++ {
							buffer.WriteString(allCities[j][k])
							if k != len(allCities[j]) - 1 {
								buffer.WriteString(",")
							}
						}
					}
				} else {
					if strings.Compare(selectedCities[i], allCities[j][3]) == 0 {
						if i < len(selectedCities) {
							buffer.WriteString("\n")
						}
						for k := 0; k < len(allCities[j]); k++ {
							buffer.WriteString(allCities[j][k])
							if k != len(allCities[j]) - 1 {
								buffer.WriteString(",")
							}
						}
					}
				}
			}
		}

		// Write the details to file
		var fileLocation bytes.Buffer
		fileLocation.WriteString("./datafiles/distributors/")
		fileLocation.WriteString(name)
		fileLocation.WriteString(".csv")

		helpers.DataToFile(fileLocation.String(), buffer.String())

		c.Ctx.ResponseWriter.Write([]byte("true"))

	} else { // Display the Add page
		view := viewmodels.NewDistributorVM{}
		view.PageTitle = "New Distributor"
		view.AllCities = allCities

		allFiles, err := ioutil.ReadDir("./datafiles/distributors/")
		if err != nil {
			log.Fatal(err)
		}

		// Reading existing distributor details
		var fileNameSlice []string

		for _, file := range allFiles {
			fileName := file.Name()
			if fileName != "readme.txt" {
				last := len(fileName) - len(".csv")
				fileName = fileName[:last]
				fileNameSlice = append(fileNameSlice, fileName)
			}
		}

		distributorCitiesMap := make(map[string][][]string)
		for i := 0; i < len(fileNameSlice); i++ {
			var fileLocation bytes.Buffer
			fileLocation.WriteString("./datafiles/distributors/")
			fileLocation.WriteString(fileNameSlice[i])
			fileLocation.WriteString(".csv")
			distributorCities, err := helpers.DataFromFile(fileLocation.String())
			if err != nil {
				log.Println(err)
			}

			distributorCitiesMap[fileNameSlice[i]] = distributorCities
		}

		view.DistributorCities = distributorCitiesMap

		// Collecting Details for add page
		var uniqueCountries, tempUniqueProvinces []string
		var uniqueProvinces [][]string
		for i := 0; i < len(allCities); i++ {
			if !helpers.StringInSlice(allCities[i][5], uniqueCountries) {
				uniqueCountries = append(uniqueCountries, allCities[i][5])
			}
			if !helpers.StringInSlice(allCities[i][4], tempUniqueProvinces) {
				provinceSlice := []string{allCities[i][4],allCities[i][5]}
				uniqueProvinces = append(uniqueProvinces, provinceSlice)
				tempUniqueProvinces = append(tempUniqueProvinces, allCities[i][4])
			}
		}

		view.UniqueCountries = uniqueCountries
		view.UniqueProvinces = uniqueProvinces

		c.Data["vm"] = view

		c.TplName = "templates/new-distributor.html"
	}



}

func (c *DistributorController) ViewDistributor() {
	name := c.GetString("name")
	allCities, err := helpers.DataFromFile("./datafiles/distributors/"+ name +".csv")
	if err != nil {
		log.Println(err)
	}
	view := viewmodels.ViewDistributorVM{}
	view.PageTitle = "New Distributor"
	view.AllCities = allCities
	c.Data["vm"] = view

	c.TplName = "templates/view-distributor.html"
}

