package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	// "io"
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
	Name               string
	IncLocs            []Location
	ExcLocs            []Location
	SubDistributerKeys []string // Sub distributer keys, we keep map of distrbuter -> its permission logic indepndendly.
}

// Single depth map to quickly locate the availability of the distributer
// for a given location.
type Distributers map[string]Distributer

/**
* Check the given distributer has permission to distribute movies under the
* given location.
*
* @param location - Formated after CITY-PROVINCE-COUNTRY format.
*
 */
func (*Distributer) has_permission(location string) bool {
	return false
}

func (d *Distributer) set_parents(name string, distributers *Distributers) {
	var parents []string
	parents = append(parents, name)
}

func PrintDistributerMap(distributers Distributers) {
	for k, v := range distributers {
		fmt.Printf("Distributer: %s\n", k)
		fmt.Printf("\tName: %s\n", v.Name)
		fmt.Printf("\tIncLocs: %s\n", v.IncLocs)
		fmt.Printf("\tExcLocs: %s\n", v.ExcLocs)
		fmt.Printf("\tSubDistributerKeys: %s\n\n", v.SubDistributerKeys)
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
func get_locations(location string) []Location {
	location = strings.TrimSpace(location)
	locs := strings.Split(location, ":")
	loc_objs := make([]Location, len(locs))
	for i, l := range locs {
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
		loc_objs[i] = *loc_obj
	}
	return loc_objs
}

/*
* @param name: formated "D1 < D2" or "D1"
*
* @return: D1, [D2, ...]
 */
func get_distributer_name(name string) (string, []string) {
	d_name := strings.Split(name, "<")
	if len(d_name) == 1 {
		return d_name[0], []string{}
	} else if len(d_name) > 1 {
		return d_name[0], d_name[1:]
	} else {
		return "", []string{}
	}
}

// Split the distributer hiearachy by spliting with delimiter as "<"
func get_sub_distributers(distributer string) []string {
	sub_distributers := make([]string, 0)
	return sub_distributers
}

func input_from_stdin() {
	for {
		fmt.Print("Enter Your choice: ")
	}
}

func main() {

	/*************************************************************************
	* Load Data
	**************************************************************************
	*
	 */

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

	distributers := make(Distributers)

	csv_reader := csv.NewReader(bufio.NewReader(dist_permission_csv))
	records, _ := csv_reader.ReadAll()

	// Skip the headers.
	for _, record := range records[1:] {
		distributer := new(Distributer)
		d_name, d_parents := get_distributer_name(record[0])
		distributer.Name = d_name
		distributer.SubDistributerKeys = d_parents
		distributer.IncLocs = get_locations(record[1])
		distributer.ExcLocs = get_locations(record[2])
		distributers[distributer.Name] = *distributer
	}
	PrintDistributerMap(distributers)

	/*************************************************************************
	* Search On the populated data.
	**************************************************************************
	*
	 */
}
