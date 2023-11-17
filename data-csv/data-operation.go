package datacsv

import (
	"encoding/csv"
	"fmt"
	"os"
	"qube-cinemas-challenge/models"
)

type Locations []*models.City

//Distributor
var Distributor []*models.Distributor

//Locations
var Cities []*models.City

//Rules
var Rules []*models.Rule

type CSV struct{
	FileName string `json:"csv"`
}

func(c *CSV) Read() ([]*models.City, error){
	file, err := os.Open(c.FileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader :=csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	var locations Locations

	for _, record:= range records[1:] {
		location:=&models.City{
			Name: record[3],
			Code: record[0],
			Province: &models.Province{
				Name: record[4],
				Code: record[1],
				Country: &models.Country{
					Name: record[5],
					Code: record[2],
				},
			},
		}
		locations = append(locations, location)
	}
	return locations, nil
}

func DataFetch() {
	//File name asking from the user
	var file CSV
	file.FileName = "cities.csv"
	var err error
	Cities, err = file.Read()
	if err != nil {
		fmt.Println("Error: "+err.Error())
	}
}