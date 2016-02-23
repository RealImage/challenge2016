package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Location struct {
	City_code     string
	Province_code string
	Country_code  string
	City_name     string
	Province_name string
	Country_name  string
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
	if len(_l.Country_code) > 0 && loc.Country_code != _l.Country_code {
		return false
	}
	if len(_l.Province_code) > 0 && loc.Province_code != _l.Province_code {
		return false
	}
	if len(_l.City_code) > 0 && loc.City_code != _l.City_code {
		return false
	}
	return true
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
	sr_loc := getLocations(location)[0]

	// Check in Include list, if found exact match return
	incLocs, excLocs := _d.GetAllLocs()
	for _, loc := range incLocs {
		if loc.IsSublocation(&sr_loc) {
			for _, loc := range excLocs {
				if loc.IsSublocation(&sr_loc) {
					return false
				}
			}
			// Matched Include location and not there in Exclude location.
			return true
		}
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
	var loc_objs []Location
	for _, l := range locs {
		loc_obj := new(Location)
		// Get sub locations.
		sub_locs := strings.Split(l, "-")
		if len(sub_locs) == 3 {
			loc_obj.City_code = sub_locs[0]
			loc_obj.Province_code = sub_locs[1]
			loc_obj.Country_code = sub_locs[2]
		} else if len(sub_locs) == 2 {
			loc_obj.Province_code = sub_locs[0]
			loc_obj.Country_code = sub_locs[1]
		} else if len(sub_locs) == 1 {
			loc_obj.Country_code = sub_locs[0]
		}

		if loc_obj.Country_code != "" || loc_obj.Province_code != "" ||
			loc_obj.City_code != "" {
			loc_objs = append(loc_objs, *loc_obj)
		}
	}
	return loc_objs
}

/*
* @param name: formated "D1 < D2" or "D1"
*
* @return: D1, [D2, ...]
 */
func getDistributerName(name string) (string, []string) {
	var d_name []string
	for _, name := range strings.Split(name, "<") {
		d_name = append(d_name, strings.TrimSpace(name))
	}
	if len(d_name) == 1 {
		return d_name[0], []string{}
	} else if len(d_name) > 1 {
		return d_name[0], d_name[1:]
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

func Load_rule_csv() {
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
		d_name, d_parents := getDistributerName(record[0])
		distributer.Name = strings.TrimSpace(d_name)
		distributer.ParentDistNames = d_parents
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
	Load_rule_csv()

	//PrintDistributerMap(DistributerMap)

	/*************************************************************************
	* Search On the populated data.
	**************************************************************************
	*
	 */
	inputFromStdin()
}
