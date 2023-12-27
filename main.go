package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var distributorPermissions = make(map[string]Permission)
var distributors = make(map[string]Distributor)
var regs []Region

func init() {
	// initialize or sync 'regs' with latest CSV file
	regs = ReadCSV()
}

func ReadCSV() []Region {
	file, err := os.Open("cities.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return nil
	}

	var regions []Region
	for _, record := range records {
		region := Region{
			CityCode:     record[0],
			ProvinceCode: record[1],
			CountryCode:  record[2],
			CityName:     record[3],
			ProvinceName: record[4],
			CountryName:  record[5],
		}
		regions = append(regions, region)
	}

	return regions
}

func main() {
	fmt.Println("$$$ ------- Welcome To Cinema Distribution Centre ------- $$$")
	for {
		fmt.Println("------------   $ Menu $  ----------------")
		fmt.Println("Please choose a service, Press key \n 1.Add a New Distributor \n 2.View Existing Distributors \n 3.Exit ")
		var key int
		fmt.Scanln(&key)
		switch key {
		case 1:
			AddDistributor()
		case 2:
			ViewDistributors()
		case 3:
			os.Exit(0)
		default:
			fmt.Println("Invalid Option...Try Again!!")
		}
	}

}
