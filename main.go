package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/bsyed6/challenge2016/model"
)

func main() {
	dataChannel := make(chan model.Region, 150)
	initializeData(dataChannel)

}

func initializeData(dataChannel chan model.Region) {
	go model.DataStore(dataChannel)
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
