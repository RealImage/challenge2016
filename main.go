package main

import (
	"fmt"
	"qube/utils"
)

func main() {
	// Define three distributors with included and excluded regions
	distributor1 := utils.Permission{
		Included: []utils.Region{
			{Country: "INDIA", State: "", City: ""},
			{Country: "UNITEDSTATES", State: "", City: ""},
		},
		Excluded: []utils.Region{
			{Country: "INDIA", State: "KARNATAKA", City: ""},
			{Country: "INDIA", State: "TAMILNADU", City: "CHENNAI"},
		},
	}

	distributor2 := utils.Permission{
		Included: []utils.Region{
			{Country: "INDIA", State: "", City: ""},
		},
		Excluded: []utils.Region{
			{Country: "INDIA", State: "KARNATAKA", City: ""},
		},
	}

	distributor3 := utils.Permission{
		Included: []utils.Region{
			{Country: "INDIA", State: "TAMILNADU", City: "TIRUCHIRAPALLI"},
		},
		Excluded: []utils.Region{},
	}

	// Call checkDistributor with the three distributors as arguments
	checkDistributor(distributor1, distributor2, distributor3)
}

func checkDistributor(distributors ...utils.Permission) {
	excludedRegions := make(map[utils.Region]bool)
	for _, distributor := range distributors {
		for _, excluded := range distributor.Excluded {
			excludedRegions[excluded] = true
		}
	}

	checkPermission := func(city, state, country string) {
		region := utils.Region{Country: country, State: state, City: city}
		if excludedRegions[region] {
			fmt.Println("no")
		} else {
			fmt.Println("yes")
		}
	}

	regions := []struct {
		city, state, country string
	}{
		{"CHICAGO", "ILLINOIS", "UNITEDSTATES"},
		{"CHENNAI", "TAMILNADU", "INDIA"},
		{"BANGALORE", "KARNATAKA", "INDIA"},
		{"MADURAI", "TAMILNADU", "INDIA"},
		{"TIRUCHIRAPALLI", "TAMILNADU", "INDIA"},
	}
	for _, r := range regions {
		checkPermission(r.city, r.state, r.country)
	}
}
