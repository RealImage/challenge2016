package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Region struct {
    Country string
    State   string
    City    string
}

type Permission struct {
    Included []Region
    Excluded []Region
}

func main() {
    // Load data from CSV file
    f, err := os.Open("cities.csv")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer f.Close()

    reader := csv.NewReader(f)
    reader.TrimLeadingSpace = true

    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Error reading CSV file:", err)
        return
    }

    // Map cities to regions
    regions := make(map[string]Region)
    for _, record := range records {
        city := record[2]
        region := Region{record[0], record[1], city}
        regions[city] = region
    }

    // Define permissions
    distributor1 := Permission{
        Included: []Region{
            {"INDIA", "", ""},
            {"UNITEDSTATES", "", ""},
        },
        Excluded: []Region{
            {"INDIA", "KARNATAKA", ""},
            {"INDIA", "TAMILNADU", "CHENNAI"},
        },
    }

    distributor2 := Permission{
        Included: []Region{
            {"INDIA", "", ""},
        },
        Excluded: []Region{
            {"INDIA", "TAMILNADU", ""},
            {"INDIA", "KARNATAKA", ""},
        },
    }

    distributor3 := Permission{
        Included: []Region{
            {"INDIA", "KARNATAKA", "HUBLI"},
        },
        Excluded: []Region{},
    }

    // Check permissions
    fmt.Println(checkPermission("CHICAGO", "ILLINOIS", "UNITEDSTATES", distributor1)) // should print true
    fmt.Println(checkPermission("CHENNAI", "TAMILNADU", "INDIA", distributor1))      // should print false
    fmt.Println(checkPermission("BANGALORE", "KARNATAKA", "INDIA", distributor1))    // should print false

    fmt.Println(checkPermission("CHICAGO", "ILLINOIS", "UNITEDSTATES", distributor2)) // should print true
    fmt.Println(checkPermission("CHENNAI", "TAMILNADU", "INDIA", distributor2))      // should print false
    fmt.Println(checkPermission("BANGALORE", "KARNATAKA", "INDIA", distributor2))    // should print false

    fmt.Println(checkPermission("HUBLI", "KARNATAKA", "INDIA", distributor3)) // should print true
    fmt.Println(checkPermission("BANGALORE", "KARNATAKA", "INDIA", distributor3)) // should print false
}

func checkPermission(city string, state string, country string, permission Permission) bool {
    region := Region{country, state, city}

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

func isSubregion(parent Region, child Region) bool {
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
