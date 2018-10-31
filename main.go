package main

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bsyed6/challenge2016/controller"
	"github.com/bsyed6/challenge2016/model"
)

func main() {
	countries := make(map[string][]string)
	cities := make(map[string][]string)
	dataChannel := make(chan model.Region, 150)
	distributors := make(map[string]model.Permission)
	initializeData(dataChannel, countries, cities)

	assign := &controller.AssignDistributor{Countries: countries, Cities: cities, Distributors: distributors}

	http.Handle("/assign", assign)
	http.ListenAndServe(":8000", nil)

}

func initializeData(dataChannel chan model.Region, data map[string][]string, cities map[string][]string) {
	go model.DataStore(dataChannel, data, cities)
	csvIn, err := os.Open("./cities.csv")

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(csvIn)
	for {
		record, err := r.Read()
		if err == io.EOF {
			close(dataChannel)
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		data := model.Region{City: record[0], Province: record[1], Country: record[2]}
		dataChannel <- data
	}
}
