package main

import (
	"example.com/realimage_2016/constants"
	"example.com/realimage_2016/distributor"
	"example.com/realimage_2016/parsing"

	"fmt"
	"os"
)

func main() {

	// Parsing the cities CSV file and storing the region data in a map
	if !parsing.ParseCsvFile(constants.CsvFilePath) {
		fmt.Println("Error while reading from csv file")
		os.Exit(1)
	}

	// Feeding input data of distributors
	includeRegions := []string{"TN-US", "IN", "JZRAH-RK-AE"}
	excludeRegions := []string{"EBRIN-TN-US", "TN-IN"}
	dis1 := distributor.AddDistributor("dis1", includeRegions, excludeRegions, nil)

	includeRegions = []string{"COLUB-TN-US", "AN-IN", "TN-IN"}
	excludeRegions = []string{"ACHMP-AP-IN"}
	dis2 := distributor.AddDistributor("dis2", includeRegions, excludeRegions, &dis1)
	
	includeRegions = []string{"NIADS-AN-IN"}
	excludeRegions = []string{}
	dis3 := distributor.AddDistributor("dis3", includeRegions, excludeRegions, &dis2)	

	// Generating ouput by validating the permission of three distributors for different regions
	regionList := []string{"EBRIN-TN-US", "AL-US", "FORTH-VA-US", "ALGPU-TN-IN", "COLIR-TN-US", "COLUB-TN-US", "NIADS-AN-IN"}
	for _, region := range regionList {
		// Distributor 1
		if distributor.IsAuthorized(dis1, region) {
			fmt.Println("Yes, distributor '", dis1.Id, "' has permission in region: ", region )
		} else {
			fmt.Println("No, distributor '", dis1.Id, "' don't have permission in region: ", region )
		}
		// Distributor 2
		if distributor.IsAuthorized(dis2, region) {
			fmt.Println("Yes, distributor '", dis2.Id, "' has permission in region: ", region )
		} else {
			fmt.Println("No, distributor '", dis2.Id, "' don't have permission in region: ", region )
		}
		// Distributor 3
		if distributor.IsAuthorized(dis3, region) {
			fmt.Println("Yes, distributor '", dis3.Id, "' has permission in region: ", region )
		} else {
			fmt.Println("No, distributor '", dis3.Id, "' don't have permission in region: ", region )
		}
		fmt.Println("---------------------------------------")
	}
	

}