// I am assuming there is  a file called distributers.csv which contains distributer's name, parent distributer and regions
//Also I am assuming that you are sending codes instead of names of the region as it could lead to confusion due to spelling  and cases mistakes

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

type Region struct {
	City     string `json:"city"`
	Province string `json:"province"`
	Country  string `json:"country"`
}
type City struct {
	CityCode string `json:"city_code"`
	CityName string `json:"city_name"`
}
type Province struct {
	ProvinceCode string `json:"province_code"`
	ProvinceName string `json:"province_name"`
}
type Country struct {
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
}
type PermissionType string

const (
	Include PermissionType = "INCLUDE"
	Exclude PermissionType = "EXCLUDE"
)

type Permission struct {
	Region string         `json:"region"`
	Type   PermissionType `json:"type"`
}

type Distributor struct {
	Name string `json:"name"`
	//Permissions []Permission `json:"permissions"`
	Included string `json:"included"`
	Excluded string `json:"excluded"`
	Parent   string `json:"parent"`
}

func loadRegions(filePath string) (map[string]*Region, error) {
	city := make(map[string]*City)
	province := make(map[string]*Province)
	country := make(map[string]*Country)
	regions := make(map[string]*Region)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Read()

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		city[record[0]] = &City{CityCode: record[0], CityName: record[3]}
		province[record[1]] = &Province{ProvinceCode: record[1], ProvinceName: record[4]}
		country[record[2]] = &Country{CountryCode: record[2], CountryName: record[5]}
		regionCode := fmt.Sprintf("%s-%s-%s", city[record[0]].CityCode, province[record[1]].ProvinceCode, country[record[2]].CountryCode)
		regions[regionCode] = &Region{City: city[record[0]].CityCode, Province: province[record[1]].ProvinceCode, Country: country[record[2]].CountryCode}
	}

	return regions, nil
}

func loadDistributors(filePath string, regions map[string]*Region) (map[string]*Distributor, error) {
	distributors := make(map[string]*Distributor)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	reader.Read()

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		distributor, parent, include, exclude := record[0], record[1], record[2], record[3]

		processedInclude, processedExclude := "", ""
		if include != "" {
			for _, region := range strings.Split(include, ",") {
				var cityCode, provinceCode, countryCode string
				splitted := strings.Split(region, "-")
				if len(splitted) == 3 {
					cityCode, provinceCode, countryCode = splitted[0], splitted[1], splitted[2]
				} else if len(splitted) == 2 {
					provinceCode, countryCode = splitted[0], splitted[1]
				} else if len(splitted) == 1 {
					countryCode = splitted[0]
				}
				regionCode := fmt.Sprintf("%s-%s-%s", cityCode, provinceCode, countryCode)
				city, province, country := regions[regionCode].City, regions[regionCode].Province, regions[regionCode].Country
				processedInclude += fmt.Sprintf("%s-%s-%s,", city, province, country)
			}
			processedInclude = strings.TrimSuffix(processedInclude, ",")
		}

		if exclude != "" {
			for _, region := range strings.Split(exclude, ",") {
				var cityCode, provinceCode, countryCode string
				splitted := strings.Split(region, "-")
				if len(splitted) == 3 {
					cityCode, provinceCode, countryCode = splitted[0], splitted[1], splitted[2]
				} else if len(splitted) == 2 {
					provinceCode, countryCode = splitted[0], splitted[1]
				} else if len(splitted) == 1 {
					countryCode = splitted[0]
				}
				regionCode := fmt.Sprintf("%s-%s-%s", cityCode, provinceCode, countryCode)
				city, province, country := regions[regionCode].City, regions[regionCode].Province, regions[regionCode].Country
				processedExclude += fmt.Sprintf("%s-%s-%s,", city, province, country)
			}
			processedExclude = strings.TrimSuffix(processedExclude, ",")
		}

		distributors[distributor] = &Distributor{Name: distributor, Parent: parent, Included: processedInclude, Excluded: processedExclude}
	}

	return distributors, nil

}

func contains(arr string, str string) bool {
	for _, item := range strings.Split(arr, ",") {
		if item == str {
			return true

		}
	}
	return false
}
func checkPermission(distributorName string, location string, distributorsList map[string]*Distributor) bool {
	distributor, ok := distributorsList[distributorName]
	if !ok {
		return false // Distributor not found
	}

	// Check if region is included
	if contains(distributor.Included, location) {

		// Check if region is excluded (overrides include)
		if !contains(distributor.Excluded, location) {
			return true
		}
	}

	// Check parent's permission if any
	if distributor.Parent != "" {
		return checkPermission(distributor.Parent, location, distributorsList)
	}

	// No permission found
	return false

}

func main() {
	// Load regions and distributors from CSV files (replace with your file paths)
	regions, err := loadRegions("cities.csv")
	if err != nil {
		fmt.Println("Error loading regions:", err)
		return
	}

	distributors, err := loadDistributors("distributors.csv", regions)
	if err != nil {
		fmt.Println("Error loading distributors:", err)
		return
	}

	// Get user input for checking permissions
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter distributor name:")
	for scanner.Scan() {
		distributorName := strings.ToUpper(scanner.Text())

		fmt.Println("Enter location (CITY-STATE-COUNTRY):")
		if scanner.Scan() {
			location := strings.ToUpper(scanner.Text())
			if checkPermission(distributorName, location, distributors) {
				fmt.Println("Yes, distribution is allowed in", location)
			} else {
				fmt.Println("No, distribution is not allowed in", location)
			}
		}

		fmt.Println("Enter another distributor name (or 'q' to quit):")
	}
}
