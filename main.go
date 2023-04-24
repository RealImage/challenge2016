package main

import (
	"fmt"
	"qube-cinemas/utils"
)

func main() {
	// Define permissions
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
			// {Country: "INDIA", State: "TAMILNADU", City: ""},
			{Country: "INDIA", State: "KARNATAKA", City: ""},
		},
	}

	distributor3 := utils.Permission{
		Included: []utils.Region{
			{Country: "INDIA", State: "TAMILNADU", City: "TIRUCHIRAPALLI"},
		},
		Excluded: []utils.Region{},
	}

	distributerAuthorized(distributor1, distributor2, distributor3)
}

func distributerAuthorized(distributor1 utils.Permission, distributor2 utils.Permission, distributor3 utils.Permission) {
	//    If data to be loaded from csv file ************
	// f, err := os.Open("cities.csv")
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	return
	// }
	// defer f.Close()

	// reader := csv.NewReader(f)
	// reader.TrimLeadingSpace = true

	// records, err := reader.ReadAll()
	// if err != nil {
	// 	fmt.Println("Error reading CSV file:", err)
	// 	return
	// }

	// Map cities to regions
	// regions := make(map[string]utils.Region)
	// for _, record := range records {
	// 	city := record[3]
	// 	region := utils.Region{Country: record[5], State: record[4], City: city}
	// 	regions[city] = region
	// }

	distributor2.Excluded = append(distributor2.Excluded, distributor1.Excluded...)
	distributor3.Excluded = append(distributor3.Excluded, distributor2.Excluded...)

	// Check permissions
	if checkPermission("CHICAGO", "ILLINOIS", "UNITEDSTATES", distributor1) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	if checkPermission("CHENNAI", "TAMILNADU", "INDIA", distributor1) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	if checkPermission("BANGALORE", "KARNATAKA", "INDIA", distributor1) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	if checkPermission("CHICAGO", "ILLINOIS", "UNITEDSTATES", distributor2) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	if checkPermission("CHENNAI", "TAMILNADU", "INDIA", distributor2) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	if checkPermission("BANGALORE", "KARNATAKA", "INDIA", distributor2) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	if checkPermission("MADURAI", "TAMILNADU", "INDIA", distributor3) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	if checkPermission("TIRUCHIRAPALLI", "TAMILNADU", "INDIA", distributor3) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func checkPermission(city string, state string, country string, permission utils.Permission) bool {
	region := utils.Region{Country: country, State: state, City: city}

	for _, excluded := range permission.Excluded {
		if isSubregion(excluded, region) {
			return false
		}
	}

	for _, included := range permission.Included {
		if isSubregion(included, region) {
			return true
		}
	}
	return false
}

func isSubregion(parent utils.Region, child utils.Region) bool {
	if parent.Country != "" && parent.Country != child.Country {
		return false
	}
	if parent.State != "" && parent.State != child.State {
		return false
	}
	if parent.City != "" && parent.City != child.City {
		return false
	}
	return true
}

// package main

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"os"
// 	"qube-cinemas/config"
// 	"qube-cinemas/utils"
// )

// var rows [][]string
// var matchFound [][]string
// var excluded [][]string
// var subInclude [][]string
// var subExclude [][]string

// func main() {
// 	distributer := utils.NewDistributer{}
// 	distributer.Name = config.MainDistributer
// 	distributer.Include = config.Include
// 	distributer.Exclude = config.Exclude
// 	distributer.Check = config.Check

// 	subDistributer := utils.NewDistributer{}
// 	subDistributer.Name = config.SubDistributer1
// 	subDistributer.Include = config.SubInclude1
// 	subDistributer.Exclude = config.SubExclude1
// 	subDistributer.Check = config.SubCheck1

// 	subDistributerList := []utils.NewDistributer{}
// 	subDistributerList = append(subDistributerList, subDistributer)
// 	fmt.Println(subDistributerList)

// 	csvFile, err := os.Open("cities.csv")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer csvFile.Close()
// 	csvLines, err := csv.NewReader(csvFile).ReadAll()
// 	// var matches []string
// 	for _, val := range distributer.Include {
// 		matchFound = find(csvLines, val, true)
// 	}

// 	for _, val := range distributer.Exclude {
// 		excluded = find(matchFound, val, false)
// 	}
// 	allowed := check(excluded, distributer.Check)
// 	if distributer.Check != "" {
// 		fmt.Println("Primary distributer allowed in this region = ", allowed)
// 	}
// 	for _, val := range subDistributerList {
// 		for _, inc := range val.Include {
// 			subInclude = find(excluded, inc, true)
// 		}
// 	}
// 	for _, val := range subDistributerList {
// 		for _, exc := range val.Exclude {
// 			subInclude = find(excluded, exc, true)
// 		}
// 		allowedSub := check(excluded, val.Check)
// 		if val.Check != "" {
// 			fmt.Printf("Distributer = %+v; allowed in this region = %+v;", val.Name, allowedSub)
// 		}
// 	}

// }

// func find(records [][]string, val string, task bool) [][]string {
// 	if task == true {
// 		for _, row := range records {
// 			if row[5] == val {
// 				rows = append(rows, row)
// 			}
// 			if row[4] == val {
// 				rows = append(rows, row)
// 			}
// 			if row[3] == val {
// 				rows = append(rows, row)
// 			}
// 			// fmt.Println(row[3], val)
// 		}
// 	} else {
// 		for i, row := range records {
// 			if row[5] == val {
// 				rows = remove(rows, i)
// 			}
// 			if row[4] == val {
// 				rows = remove(rows, i)
// 			}
// 			if row[3] == val {
// 				rows = remove(rows, i)
// 			}
// 		}
// 	}
// 	return rows
// }

// func check(records [][]string, val string) (allowed bool) {
// 	for _, row := range records {
// 		if row[5] == val {
// 			return true
// 		}
// 		if row[4] == val {
// 			return true
// 		}
// 		if row[3] == val {
// 			return true
// 		}
// 	}
// 	return false
// }

// func remove(s [][]string, i int) [][]string {
// 	return append(s[:i], s[i+1:]...)
// }
