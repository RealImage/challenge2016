package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"bufio"
	"strings"
	// "io"
)

type Location struct {
	City_code string
	Province_code string
	Country_code string
	City_name string
	Province_name string
	Country_name string
}

type Distributer struct {
	Name string
	IncLocs []Location
	ExcLocs []Location
	SubDistributerKeys []string // Sub distributer keys, we keep map of distrbuter -> its permission logic indepndendly.
}

// Single depth map to quickly locate the availability of the distributer
// for a given location.
type Distributers map[string]Distributer;

func PrintDistributerMap(distributers Distributers) {
	for k, v := range distributers {
		fmt.Print("===> Key: ", k)
		fmt.Println("  Value: ", v)
	}
}

func get_locations(location string) []Location {
	location = strings.TrimSpace(location)
	locs := strings.Split(location, ":")
	loc_objs := make([]Location, len(locs))
	for i, l := range locs {
		loc_obj := new(Location)
		// Get sub locations.
		sub_locs := strings.Split(l, "-");
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
	return loc_objs;
}

func get_distributer_name(name string) string {
	d_name := strings.Split(name, "<")
	if len(d_name) >= 1 {
		return d_name[0]
	} else {
		return ""
	}
}
// Split the distributer hiearachy by spliting with delimiter as "<"
func get_sub_distributers(distributer string) []string {
	sub_distributers := make([]string, 0)
	return sub_distributers;
}

func main() {

	// CSV Format
	// Distributers, Included location, Excluded locations
	// D1, CITY-ST-COUN:..., CITY-ST-COUNT:...
	// D2 < D1, .....
	dist_permission_csv, _ := os.Open("./dist_permission.csv")

	distributers := make(Distributers)

	csv_reader := csv.NewReader(bufio.NewReader(dist_permission_csv))
	records, _ := csv_reader.ReadAll()

	// Skip the headers.
	for _ , record := range records[1:] {
		//fmt.Println(index, record)
		distributer := new (Distributer)
		distributer.Name = get_distributer_name(record[0])
		distributer.SubDistributerKeys = get_sub_distributers(record[0])
		distributer.IncLocs = get_locations(record[1])
		distributer.ExcLocs = get_locations(record[2])
		distributers[record[0]] = *distributer;
	}
	//fmt.Println(records[1:])
	PrintDistributerMap(distributers)
}