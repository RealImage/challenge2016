package main

import (
	model "challenge2016/models"
	database "challenge2016/utils"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"challenge2016/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Qube cinema")
	database.GetInstancemysql()

	csvfile, err := os.Open("cities.csv")
	if err != nil {
		fmt.Println("Couldn't open the csv file", err)
	}
	record := csv.NewReader(csvfile)
	var cities []model.Cities
	for {
		// Read each record from csv
		singlerecord, err := record.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		data := model.Cities{
			City_Code:     singlerecord[0],
			Province_Code: strings.Replace(singlerecord[1], " ", "", -1),
			Country_Code:  strings.Replace(singlerecord[2], " ", "", -1),
			City_Name:     strings.Replace(singlerecord[3], " ", "", -1),
			Province_Name: strings.Replace(singlerecord[4], " ", "", -1),
			Country_Name:  strings.Replace(singlerecord[5], " ", "", -1),
		}
		cities = append(cities, data)
	}
	r := gin.Default()
	routers.Endpoints(r)
}
