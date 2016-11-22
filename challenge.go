package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Location struct {
	CityCode      string
	ProvinceCode  string
	CountryCode   string
	CityName      string
	ProvinceName  string
	CountryName   string
	CanonicalName string
}

type Distributer struct {
	Name            string
	IncLocs         []Location
	ExcLocs         []Location
	ParentDistNames []string
}

// Single depth map to quickly locate the availability of the distributer
// for a given location.
type Distributers map[string]Distributer

var DistributerMap Distributers

/**
* Check given location is comes under this location object.
*
* TODO: Use better substring matching, Tries or so.
*
 */
func (_l *Location) IsSublocation(loc *Location) bool {
	if len(_l.CountryCode) > 0 && loc.CountryCode != _l.CountryCode {
		return false
	}
	if len(_l.ProvinceCode) > 0 && loc.ProvinceCode != _l.ProvinceCode {
		return false
	}
	if len(_l.CityCode) > 0 && loc.CityCode != _l.CityCode {
		return false
	}
	return true
}

/**
* Check the loc objects comes under any of given location.
 */
func (_l *Location) IsUnderAny(locs []Location) bool {
	for _, loc := range locs {
		if loc.IsSublocation(_l) {
			return true
		}
	}
	return false
}

/**
* Get all Include and Exclude Location recursively for given distributer.
 */
func (_d *Distributer) GetAllLocs() ([]Location, []Location) {
	var incLocs, excLocs []Location
	incLocs = append(incLocs, _d.IncLocs...)
	excLocs = append(excLocs, _d.ExcLocs...)
	for _, d_parent := range _d.ParentDistNames {
		parent_obj := DistributerMap[d_parent]
		parentIncLocs, parentExcLocs := parent_obj.GetAllLocs()
		incLocs = append(incLocs, parentIncLocs...)
		excLocs = append(excLocs, parentExcLocs...)
	}
	return incLocs, excLocs
}

/**
* Check the given distributer has permission to distribute movies under the
* given location.
*
* @param location - Formated after CITY-PROVINCE-COUNTRY format.
*
 */
func (_d *Distributer) HasPermission(location string) bool {
	srLoc := getLocations(location)[0]

	// Check in Include list, if found exact match return
	incLocs, excLocs := _d.GetAllLocs()
	if len(incLocs) > 0 {
		for _, loc := range incLocs {
			if loc.IsSublocation(&srLoc) {
				return !srLoc.IsUnderAny(excLocs)
			}
		}
	} else {
		return !srLoc.IsUnderAny(excLocs)
	}
	// Doesn't match any Include locations.
	return false
}

func PrintDistributerMap(distributers Distributers) {
	for k, v := range distributers {
		fmt.Printf("Distributer: %s\n", k)
		fmt.Printf("\tName: %s\n", v.Name)
		fmt.Printf("\tIncLocs: %s\n", v.IncLocs)
		fmt.Printf("\tExcLocs: %s\n", v.ExcLocs)
		fmt.Printf("\tParentDistName: %s\n\n", v.ParentDistNames)
	}
}

/*
* Split the locations from the input CSV columns.
*
* Format in each location csv would be:
*	Location1:Location2:etc..
*
*	Where Location1 follows CITY-PROVINCE-COUNTRY format.
 */
func getLocations(location string) []Location {
	location = strings.TrimSpace(location)
	locs := strings.Split(location, ":")
	var locObjs []Location
	for _, l := range locs {
		locObj := new(Location)
		// Get sub locations.
		locObj.CanonicalName = l
		subLocs := strings.Split(l, "-")
		if len(subLocs) == 3 {
			locObj.CityCode = subLocs[0]
			locObj.ProvinceCode = subLocs[1]
			locObj.CountryCode = subLocs[2]
		} else if len(subLocs) == 2 {
			locObj.ProvinceCode = subLocs[0]
			locObj.CountryCode = subLocs[1]
		} else if len(subLocs) == 1 {
			locObj.CountryCode = subLocs[0]
		}

		if locObj.CountryCode != "" || locObj.ProvinceCode != "" ||
			locObj.CityCode != "" {
			locObjs = append(locObjs, *locObj)
		}
	}
	return locObjs
}

/*
* @param name: formated "D1 < D2" or "D1"
*
* @return: D1, [D2, ...]
 */
func getDistributerName(name string) (string, []string) {
	var dName []string
	for _, name := range strings.Split(name, "<") {
		dName = append(dName, strings.TrimSpace(name))
	}
	if len(dName) == 1 {
		return dName[0], []string{}
	} else if len(dName) > 1 {
		return dName[0], dName[1:]
	} else {
		return "", []string{}
	}
}

func HasAuthorized(d string, l string) bool {
	d_obj := DistributerMap[d]
	return d_obj.HasPermission(l)
}

func inputFromStdin() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\n1. Check Distributer Permission: ")
		fmt.Print("\n2. Quit.")
		fmt.Print("\nEnter Your choice: ")
		input, _ := reader.ReadString('\n')
		fmt.Printf("Your Choice: %s", input)

		switch strings.TrimSpace(input) {
		case "1":
			fmt.Print("Enter the Input eg: D1, CITY-PROVINCE-COUNTRY format: ")
			input, _ := reader.ReadString('\n')
			text := strings.Split(input, ",")
			d, l := strings.TrimSpace(text[0]), strings.TrimSpace(text[1])

			HasPermission := ""
			if HasAuthorized(d, l) {
				HasPermission = "ON"
			} else {
				HasPermission = "OFF"
			}
			fmt.Printf(" Distributer: %s, Location %s - Has permission ?: %s",
				d, l, HasPermission)
		case "2":
			os.Exit(0)
		default:
			fmt.Println("Please pick correct option...")
		}
	}
}

func LoadRuleCsv() {
	// CSV Format
	// Distributers, Included location, Excluded locations
	// D1, CITY-ST-COUN:..., CITY-ST-COUNT:...
	// D2 < D1, .....
	//
	// NOTE: CSV format is in such a way that we use key:value pairs to
	// designate the distributer hierarchy. Ie; We have Top level distributers
	// who are not inheriting from other parent distributers.
	//
	// Right now we have to ensure the order in which we input the distributer
	// relation, we have to pass the parent distributers first, so that when the
	// child distributers comes they can refer already known parent distributer
	// details.
	//
	// TODO: Use two-pass search to avoid this constrain.
	dist_permission_csv, _ := os.Open("./dist_permission.csv")

	DistributerMap = make(Distributers)

	csv_reader := csv.NewReader(bufio.NewReader(dist_permission_csv))
	records, _ := csv_reader.ReadAll()

	// Skip the headers.
	for _, record := range records[1:] {
		distributer := new(Distributer)
		dName, dParents := getDistributerName(record[0])
		distributer.Name = strings.TrimSpace(dName)
		distributer.ParentDistNames = dParents
		distributer.IncLocs = getLocations(record[1])
		distributer.ExcLocs = getLocations(record[2])
		DistributerMap[distributer.Name] = *distributer
	}
}

func main() {

	/*************************************************************************
	* Load Data
	**************************************************************************
	*
	 */
	LoadRuleCsv()

	//PrintDistributerMap(DistributerMap)

	/*************************************************************************
	* Search On the populated data.
	**************************************************************************
	*
	 */
	inputFromStdin()
}
