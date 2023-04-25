package main

import (
	"fmt"
	"time"

	csv2 "github.com/ishhyoboytarun/challenge2016/csv"
	"github.com/ishhyoboytarun/challenge2016/distributor"
	"github.com/ishhyoboytarun/challenge2016/handler"
)

func main() {

	csv := &csv2.Csv{
		FileName: "cities.csv",
	}
	regions, err := csv.Read()
	if err != nil {
		fmt.Println("Error reading and parsing file")
		return
	}

	countriesToRegionsMap := handler.FetchCountriesToRegionMap(regions)

	distributor := &handler.Distributors{
		CountryLevelDistributorsMap:  make(map[*distributor.Distributor]bool),
		ProvinceLevelDistributorsMap: make(map[*distributor.Distributor]bool),
		CountryWiseDistributorsCount: 5,
	}

	distributor.AllotRegionsToDistributorsCountryWise(countriesToRegionsMap)

	fmt.Println("Country level distributors ->")
	time.Sleep(time.Second * 2)
	fmt.Println(distributor.CountryLevelDistributorsMap)

	time.Sleep(time.Second * 2)

	fmt.Println("Province level distributors ->")
	time.Sleep(time.Second * 2)
	distributor.AddProvinceLevelDistributors()
	fmt.Println(distributor.ProvinceLevelDistributorsMap)
}
