package main

import (
	"log"
	"strconv"
	"testing"

	"github.com/atyagi9006/challenge2016/distributer"

	"github.com/atyagi9006/challenge2016/csvreader"
	"github.com/atyagi9006/challenge2016/models"
	"github.com/atyagi9006/challenge2016/utilites"
)

func TestStaticInput(t *testing.T) {
	csvFileName := "cities.csv"
	distributerMap := make(models.DistributerMap)
	countryStateMap := make(models.CountryMap)
	csvreader.MakeDataStore(csvFileName, countryStateMap)

	input := models.InputModel{
		Name:       utilites.UpperCaseNoSpace("distributer"),
		Permission: "India",
		AuthType:   models.Include,
	}

	err := distributer.AddDistributer(input, countryStateMap, distributerMap)

	if err != nil {
		log.Printf("Error : %v \n", err)
	}

	input1 := models.InputModel{
		Name:       utilites.UpperCaseNoSpace("distributer"),
		Permission: "Tamil Nadu-India",
		AuthType:   models.Exclude,
	}

	err = distributer.AddDistributer(input1, countryStateMap, distributerMap)

	if err != nil {
		log.Printf("Error : %v \n", err)
	}

	input2 := models.InputModel{
		Name:       utilites.UpperCaseNoSpace("distributer1 < distributer"),
		Permission: "Keelakarai-Tamil Nadu-India",
		AuthType:   models.Include,
	}

	err = distributer.AddDistributer(input2, countryStateMap, distributerMap)

	if err != nil {
		log.Printf("Error : %v \n", err)
	}
}

func TestAddDistributer(t *testing.T) {

	testScenrio := []struct {
		input models.InputModel
	}{
		{
			input: models.InputModel{
				Name:       utilites.UpperCaseNoSpace("distributer"),
				Permission: "India",
				AuthType:   models.Include,
			},
		}, {
			input: models.InputModel{
				Name:       utilites.UpperCaseNoSpace("distributer"),
				Permission: "Tamil Nadu-India",
				AuthType:   models.Exclude,
			},
		}, {
			input: models.InputModel{
				Name:       utilites.UpperCaseNoSpace("distributer1 < distributer"),
				Permission: "Keelakarai-Tamil Nadu-India",
				AuthType:   models.Include,
			},
		},
	}
	csvFileName := "cities.csv"
	distributerMap := make(models.DistributerMap)
	countryStateMap := make(models.CountryMap)

	csvreader.MakeDataStore(csvFileName, countryStateMap)

	for i, scenrio := range testScenrio {
		t.Run(strconv.Itoa(i), func(t *testing.T) {

			errmsg := distributer.AddDistributer(scenrio.input, countryStateMap, distributerMap)

			if errmsg != nil {
				log.Println("Error is : ", errmsg)
				return //expected error; done here
			}

		})
	}
}
