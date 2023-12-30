package distributor

import (
	"fmt"
	"qubecinema/parse"
	"testing"
)

func validateRegionMap(t *testing.T, distributor *Distributor, country, province, city string, expected bool) {
	t.Helper()
	query := fmt.Sprintf("%s-%s-%s", city, province, country)
	result := distributor.CheckCityProvinceCountry(query)
	if result {
		fmt.Printf("Yes Distributor %d can distribute in %s\n", distributor.DistributorID, query)
	} else {
		fmt.Printf("No Distributor %d  cannot distribute in %s\n", distributor.DistributorID, query)
	}
	if result != expected {
		t.Errorf("For query %s, expected %t, but got %t", query, expected, result)
	}
}
//for debugging
func printRegionMap(regionMap map[string]map[string]map[string]bool) {
	for country, provinces := range regionMap {
		fmt.Println(country)
		for province, cities := range provinces {
			fmt.Printf("  %s\n", province)
			for city := range cities {
				fmt.Printf("    %s\n", city)
			}
		}
	}
}

func TestDistributors(t *testing.T) {
	// Original region m
	filePath := "../cities.csv"

	originalRegionMap, err := parse.ParseCSV(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Test case 1
	distributor1, err :=  NewDistributor(originalRegionMap, []string{"IN"}, []string{}, 1)
	if err != nil {
		t.Errorf("Error creating distributor1: %v", err)
	}else{
	validateRegionMap(t, distributor1, "IN", "KA", "YELUR", true) // Updated from "BANGALORE" to "YELUR"
	validateRegionMap(t, distributor1, "IN", "TN", "CENAI", true)
	validateRegionMap(t, distributor1, "US", "CA", "AGOUR", false) // Updated from "LOSANGELES" to "AGOUR"
	}
	// Test case 2
	distributor2, err := NewDistributor(distributor1.RegionMap, []string{"TN-IN"}, []string{"KA-IN"}, 2)
	if err != nil {
		t.Errorf("Error creating distributor2: %v", err)
	}else {
	validateRegionMap(t, distributor2, "IN", "KA", "YELUR", false) // Updated from "BANGALORE" to "YELUR"
	validateRegionMap(t, distributor2, "IN", "TN", "CENAI", true)
	validateRegionMap(t, distributor2, "US", "IL", "CHIAO", false)

	// Test case 3
	distributor3, err :=  NewDistributor(distributor2.RegionMap, []string{"CENAI-TN-IN"}, []string{}, 3)
	if err != nil {
		t.Errorf("Error creating distributor3: %v", err)
	}else{
	validateRegionMap(t, distributor3, "IN", "TN", "CENAI", true)
	validateRegionMap(t, distributor3, "US", "IL", "CHIAO", false)}}
}
