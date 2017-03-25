package controllers

import (
	"challenge2016/helpers"
	"log"
	"challenge2016/viewmodels"
)

type DistributorController struct {
	BaseController
}

func (c *DistributorController) NewDistributor() {
	r := c.Ctx.Request
	if(r.Method == "POST"){

	} else {
		view := viewmodels.NewDistributorVM{}
		view.PageTitle = "New Distributor"
		allCities, err := helpers.DataFromFile("./cities.csv")
		if err != nil {
			log.Println(err)
		}
		view.AllCities = allCities

		//var uniqueCountries, uniqueProvinces, uniqueCities []string
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

