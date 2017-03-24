package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
)

type DataExtractController struct {
	BaseController
}

func (c *DataExtractController) ExtractCities() {
	fmt.Println("hihihi")
	allCities, err := ioutil.ReadFile("./cities.csv")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("CIties: ", string(allCities))
}