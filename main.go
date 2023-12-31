package main

import (
	"fmt"
	"qubecinema/distributor"
	"qubecinema/parse"
	"encoding/json"
	"io/ioutil"


)

type DistributorData struct {
	ID       int               `json:"id"`
	Includes []string          `json:"includes"`
	Excludes []string          `json:"excludes"`
	Children []*DistributorData `json:"children"`
}
type QueryData struct {
	ID      int      `json:"id"`
	Queries []string `json:"queries"`
}
func main() {

	filePath := "cities.csv"

	regionMap, err := parse.ParseCSV(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Read distributor information from JSON file
	distributorData, err := ioutil.ReadFile("inputs/input.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Unmarshal JSON data
	var distributorsData []DistributorData
	err = json.Unmarshal(distributorData, &distributorsData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Create distributors
	var distributors []*distributor.Distributor
	for _, data := range distributorsData {
		// Assuming createRegionMap is a function that creates the initial region map
		

		parentdistributor, err := distributor.NewDistributor(regionMap, data.Includes, data.Excludes,data.ID)
		if err != nil {
			fmt.Println("Error creating distributor:", err)
			return
		}

		// Recursively create child distributors
		for _, childData := range data.Children {
			childRegionMap := parentdistributor.RegionMap // Child inherits parent's region map
			childDistributor, err := distributor.NewDistributor(childRegionMap, childData.Includes, childData.Excludes,childData.ID)
			if err != nil {
				fmt.Println("Error creating child distributor:", err)
				return
			}
			distributors = append(distributors, childDistributor)

		}

		distributors = append(distributors, parentdistributor)
	}

	
	queryData, err := ioutil.ReadFile("inputs/query.json")
	if err != nil {
		fmt.Println("Error reading query JSON file:", err)
		return
	}

	// Unmarshal JSON data
	var queriesData []QueryData
	err = json.Unmarshal(queryData, &queriesData)
	if err != nil {
		fmt.Println("Error unmarshalling query JSON:", err)
		return
	}

	
	// Loop through queries and check if they exist in each distributor's region map
	for _, qData := range queriesData {
		for _, distributor := range distributors {
			if distributor.DistributorID == qData.ID {
				fmt.Printf("Queries for Distributor %d:\n", qData.ID)
				for _, query := range qData.Queries {
					if distributor.CheckCityProvinceCountry(query) {
						fmt.Printf("  Yes the region %s: is a valid region to distribute for distributor %d \n", query,distributor.DistributorID)
					} else {
						fmt.Printf("  No the region %s: is not a valid region to distribute for distributor %d \n", query,distributor.DistributorID)
					}
				}
				fmt.Println()
				break
			}
		}
	}



}