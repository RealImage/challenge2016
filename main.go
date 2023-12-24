package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Region struct {
	Name     string
	Location string
}

type Distributor struct {
	Name           string
	IncludeRegions []Region
	ExcludeRegions []Region
	regionName     string
	location       string
}

func parseRegion(regionStr string) Region {
	parts := strings.Split(regionStr, "-")
	return Region{
		Name:     parts[0],
		Location: parts[1],
	}
}

func readDataFromCSVFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Unable to read input file "+fileName, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+fileName, err)
	}

	return data
}

func parseDistributorData(data [][]string) map[string]Distributor {
	distributors := make(map[string]Distributor)

	for _, row := range data {
		distributorName := row[0]
		regionName := row[1]
		location := row[2]
		includeRegionsStr := row[3]
		excludeRegionsStr := row[4]

		var includeRegions []Region
		for _, regionStr := range strings.Split(includeRegionsStr, ",") {
			includeRegions = append(includeRegions, parseRegion(regionStr))
		}

		var excludeRegions []Region
		for _, regionStr := range strings.Split(excludeRegionsStr, ",") {
			excludeRegions = append(excludeRegions, parseRegion(regionStr))
		}

		distributors[distributorName] = Distributor{
			Name:           distributorName,
			IncludeRegions: includeRegions,
			ExcludeRegions: excludeRegions,
			regionName:     regionName,
			location:       location,
		}
		fmt.Println(distributors)

	}

	return distributors
}

func checkDistributorPermissions(distributorName string, region Region) bool {
	distributorData := readDataFromCSVFile("cities.csv")

	fmt.Println(distributorData, "check")
	distributors := parseDistributorData(distributorData)

	fmt.Println(distributors)

	distributor, ok := distributors[distributorName]
	if !ok {
		return false
	}

	isIncluded := false
	isExcluded := false

	for _, includeRegion := range distributor.IncludeRegions {
		if region.Name == includeRegion.Name && region.Location == includeRegion.Location {
			isIncluded = true
			break
		}
	}

	for _, excludeRegion := range distributor.ExcludeRegions {
		if region.Name == excludeRegion.Name && region.Location == excludeRegion.Location {
			isExcluded = true
			break
		}
	}

	return isIncluded && !isExcluded
}

func main() {
	http.HandleFunc("/checkpermissions", func(w http.ResponseWriter, r *http.Request) {
		distributorName := r.URL.Query().Get("distributor_name")
		regionName := r.URL.Query().Get("region_name")
		location := r.URL.Query().Get("location")

		fmt.Println(regionName, distributorName, location)

		region := Region{
			Name:     regionName,
			Location: location,
		}
		fmt.Println(region, distributorName, location)
		data, err := json.Marshal(region)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if distributorName == "" || regionName == "" || location == "" {
			fmt.Fprint(w, "No Distributor, Region, or Location provided.")
			return
		}
		fmt.Fprint(w, string(data))
		if checkDistributorPermissions(distributorName, region) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Distributor has permission to operate in the region"))
		} else {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Distributor does not have permission to operate in the region"))
		}
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
