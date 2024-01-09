package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mitchellh/copystructure"
)

type ParsedDistributor struct {
	ID                int                  `json:"id"`
	Inclusions        []string             `json:"includes"`
	Exclusions        []string             `json:"excludes"`
	ChildDistributors []*ParsedDistributor `json:"children"`
}

// Distributor represents a distributor object.
type Distributor struct {
	RegionMap     map[string]map[string]map[string]bool // Mapping of regions (country, province, city)
	DistributorID int                                   // Unique identifier for the distributor
}

// Structure for query data
type QueryInfo struct {
	ID      int      `json:"id"`
	Queries []string `json:"queries"`
}

// ParseCSVFile parses the CSV file and returns a map representing the region structure.
func ParseCSVFile(filePath string) (map[string]map[string]map[string]bool, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening CSV file: %v", err)
	}
	defer file.Close()

	// Create a CSV reader
	csvReader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV: %v", err)
	}

	// Create a map to store regions
	regionMap := make(map[string]map[string]map[string]bool)

	// Iterate through the CSV records and populate the region map
	for _, record := range records {
		cityCode := record[0]
		provinceCode := record[1]
		regionCode := record[2]

		// Check if the region exists in the map
		region, regionExists := regionMap[regionCode]
		if !regionExists {
			// If the region doesn't exist, create a new entry
			region = make(map[string]map[string]bool)
			regionMap[regionCode] = region
		}

		// Check if the province exists in the region map
		province, provinceExists := region[provinceCode]
		if !provinceExists {
			// If the province doesn't exist, create a new entry
			province = make(map[string]bool)
			region[provinceCode] = province
		}

		// Add the city to the province
		province[cityCode] = true
	}

	return regionMap, nil
}

// NewDistributor creates a new Distributor object based on the provided inputs.
func NewDistributor(originalRegionMap map[string]map[string]map[string]bool, includes, excludes []string, distributorID int) (*Distributor, error) {
	newRegionMap := make(map[string]map[string]map[string]bool)

	// Process includes
	for _, include := range includes {
		parts := strings.Split(include, "-")
		switch len(parts) {
		case 1:
			// Include country
			country := parts[0]
			if countryMap, exists := originalRegionMap[country]; exists {
				newRegionMap[country] = countryMap
			} else {
				return nil, fmt.Errorf("included country '%s' not found in the original region map", country)
			}
		case 2:
			// Include province-country
			province, country := parts[0], parts[1]
			if countryMap, exists := originalRegionMap[country]; exists {
				if provinceMap, exists := countryMap[province]; exists {
					copiedMapInterface, _ := copystructure.Copy(provinceMap)
					tempProvinceMap, _ := copiedMapInterface.(map[string]bool)
					newRegionMap[country] = map[string]map[string]bool{province: tempProvinceMap}
				} else {
					return nil, fmt.Errorf("included province '%s-%s' not found in the original region map", province, country)
				}
			} else {
				return nil, fmt.Errorf("included country '%s' not found in the original region map", country)
			}
		case 3:
			// Include city-province-country
			city, province, country := parts[0], parts[1], parts[2]
			if countryMap, exists := originalRegionMap[country]; exists {
				if provinceMap, exists := countryMap[province]; exists {
					if cityMap, exists := provinceMap[city]; exists {
						copiedMapInterface, _ := copystructure.Copy(cityMap)
						tempCityMap, _ := copiedMapInterface.(bool)
						newRegionMap[country] = map[string]map[string]bool{province: {city: tempCityMap}}
					} else {
						return nil, fmt.Errorf("included city '%s-%s-%s' not found in the original region map", city, province, country)
					}
				} else {
					return nil, fmt.Errorf("included province '%s-%s' not found in the original region map", province, country)
				}
			} else {
				return nil, fmt.Errorf("included country '%s' not found in the original region map", country)
			}
		default:
			return nil, fmt.Errorf("invalid include format: %s", include)
		}
	}

	// Process excludes
	for _, exclude := range excludes {
		parts := strings.Split(exclude, "-")
		switch len(parts) {
		case 1:
			// Exclude country
			country := parts[0]
			delete(newRegionMap, country)
		case 2:
			// Exclude province-country
			province, country := parts[0], parts[1]
			if countryMap, exists := newRegionMap[country]; exists {
				delete(countryMap, province)
				if len(countryMap) == 0 {
					delete(newRegionMap, country)
				}
			}
		case 3:
			// Exclude city-province-country
			city, province, country := parts[0], parts[1], parts[2]
			if city == "YELUR" {
				fmt.Println("Deleting", city, distributorID)
			}
			if countryMap, exists := newRegionMap[country]; exists {
				if provinceMap, exists := countryMap[province]; exists {
					delete(provinceMap, city)
					if len(provinceMap) == 0 {
						delete(countryMap, province)
						if len(countryMap) == 0 {
							delete(newRegionMap, country)
						}
					}
				}
			}
		default:
			return nil, fmt.Errorf("invalid exclude format: %s", exclude)
		}
	}

	return &Distributor{RegionMap: newRegionMap, DistributorID: distributorID}, nil
}

func main() {
	citiesFilePath := "cities.csv"

	// Read country and region information from the CSV file
	countryRegionMap, err := ParseCSVFile(citiesFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//fmt.Println(countryRegionMap)

	distributorJSON, err := ioutil.ReadFile("distribution_config.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Unmarshal JSON data
	var distributorsParsed []ParsedDistributor
	err = json.Unmarshal(distributorJSON, &distributorsParsed)
	if err != nil {
		fmt.Println("Error unmarshalling JSON parsed distributors data:", err)
		return
	}

	// Create distributors
	var distributors []*Distributor
	for _, data := range distributorsParsed {
		// Assuming createRegionMap is a function that creates the initial region map

		parentDistributor, err := NewDistributor(countryRegionMap, data.Inclusions, data.Exclusions, data.ID)
		if err != nil {
			fmt.Println("Error creating distributor:", err)
			return
		}

		// Recursive creation of child distributors
		var recCreateChild func(*[]*Distributor, ParsedDistributor, *Distributor)
		recCreateChild = func(distributors *[]*Distributor, data ParsedDistributor, parentDistributor *Distributor) {
			// Base case: no children
			if data.ChildDistributors == nil {
				return
			}

			for _, childData := range data.ChildDistributors {
				childRegionMap := parentDistributor.RegionMap // Child inherits parent's region map
				childDistributor, err := NewDistributor(childRegionMap, childData.Inclusions, childData.Exclusions, childData.ID)
				if err != nil {
					fmt.Println("Error creating child distributor:", err)
					return
				}
				*distributors = append(*distributors, childDistributor)
				recCreateChild(distributors, *childData, childDistributor)
			}
		}

		recCreateChild(&distributors, data, parentDistributor)

		distributors = append(distributors, parentDistributor)
	}

	// Read query data from JSON file
	queryFileData, err := ioutil.ReadFile("distribution_queries.json")
	if err != nil {
		fmt.Println("Error reading query JSON file:", err)
		return
	}

	// Unmarshal JSON data
	var queryInfoList []QueryInfo
	err = json.Unmarshal(queryFileData, &queryInfoList)
	if err != nil {
		fmt.Println("Error unmarshalling query JSON:", err)
		return
	}

	// Loop through query info and check if they exist in each distributor's region map
	for _, qInfo := range queryInfoList {
		for _, distributor := range distributors {
			if distributor.DistributorID == qInfo.ID {
				fmt.Printf("Results for Distributor %d:\n", qInfo.ID)
				for _, query := range qInfo.Queries {
					if distributor.CheckCityProvinceCountry(query) {
						fmt.Printf("  The region %s is included in the distribution area for distributor %d.\n", query, distributor.DistributorID)
					} else {
						fmt.Printf("  The region %s is not included in the distribution area for distributor %d.\n", query, distributor.DistributorID)
					}
				}
				fmt.Println()
				break
			}
		}
	}

}

func recursivelyCreateChild(distributors *[]*Distributor, data ParsedDistributor, parentDistributor *Distributor) {
	// Base case: no children
	if data.ChildDistributors == nil {
		return
	}
	fmt.Println("Inside Recursive Function")

	for _, childData := range data.ChildDistributors {
		fmt.Println("Inside For Loop")
		childRegionMap := parentDistributor.RegionMap // Child inherits parent's region map
		childDistributor, err := NewDistributor(childRegionMap, childData.Inclusions, childData.Exclusions, childData.ID)
		// fmt.Println(childDistributor.DistributorID, childDistributor.RegionMap)
		if err != nil {
			fmt.Println("Error creating child distributor:", err)
			return
		}
		*distributors = append(*distributors, childDistributor)
		// fmt.Println(len(distributors))
		recursivelyCreateChild(distributors, *childData, childDistributor)
	}
}

// CheckCityProvinceCountry checks whether a specific city-province-country combination exists in the distributor's region map.
func (d *Distributor) CheckCityProvinceCountry(query string) bool {
	parts := strings.Split(query, "-")
	if len(parts) != 3 {
		return false
	}

	city, province, country := parts[0], parts[1], parts[2]
	if provinces, exists := d.RegionMap[country]; exists {
		if cities, exists := provinces[province]; exists {
			if _, exists := cities[city]; exists {
				return true
			}
			return false
		}
	}

	return false
}
