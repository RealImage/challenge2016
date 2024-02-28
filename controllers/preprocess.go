package controllers

import (
	"bufio"
	"encoding/csv"
	"os"
	"strings"

	"realImage.com/m/model"
)

var LocMap map[string]model.Location
var DistributorMap map[string][]model.Distributor // to track all parents and their children, for easy retrieval
var scanner *bufio.Scanner

// Preprocess the provided data, create a map out of the same for easy access.
func Preprocess() {
	scanner = bufio.NewScanner(os.Stdin)
	DistributorMap = make(map[string][]model.Distributor, 0)

	file, err := os.Open("cities.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	LocMap = make(map[string]model.Location)
	for _, ele := range data {
		loc := model.Location{
			City:    strings.ToLower(ele[3]),
			State:   strings.ToLower(ele[4]),
			Country: strings.ToLower(ele[5]),
		}
		LocMap[loc.Country] = loc
		LocMap[loc.State] = loc
		LocMap[loc.City] = loc
	}
}
