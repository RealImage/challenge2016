package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

// DistributorInfo ...
type DistributorInfo struct {
	include               bool
	exclude               bool
	city                  string
	state                 string
	country               string
	cityStateCountryCount int
}

// main ...
func main() {
	fileInfo, err := ReadCsvFileInfo("sample.csv")
	if err != nil {
		log.Println("error occure while processing the distructor file information,err:", err)
		os.Exit(1)
	}
	if len(os.Args) != 3 {
		log.Println("invalid arguments are passed, please provide enough details to validate")
		os.Exit(0)
	}
	inputName := os.Args[1]
	inputLocation := os.Args[2]

	isMatch := IsDistributorMatched(inputLocation, fileInfo, inputName)
	if isMatch {
		log.Println("YES")
	} else {
		log.Println("NO")
	}
}

// IsDistributorMatched ... checking the input distributor matches with the file information of distributors
func IsDistributorMatched(inputLocation string, fileInfo map[string][]DistributorInfo, inputName string) bool {
	//fetching the request distributorInfo information
	reqDistributorInfo := fetchRegionInformation(inputLocation)
	distributorDetails := fileInfo[inputName]
	isMatched := false
	for _, distributor := range distributorDetails {
		if distributor.include && distributor.country == reqDistributorInfo.country {
			isMatched = true
			break
		}
	}
	return checkExcludePermission(distributorDetails, reqDistributorInfo, isMatched)
}

// checkExcludePermission ...
func checkExcludePermission(distributorDetails []DistributorInfo, reqDistributorInfo DistributorInfo, isMatch bool) bool {

	for _, distributor := range distributorDetails {
		if distributor.include {
			continue
		}
		switch distributor.cityStateCountryCount {
		case 3:
			if distributor.city == reqDistributorInfo.city && distributor.state == reqDistributorInfo.state && distributor.country == reqDistributorInfo.country {
				isMatch = false
			}
		case 2:
			if distributor.city == reqDistributorInfo.city && distributor.state == reqDistributorInfo.state {
				isMatch = false
			}
		default:
			if distributor.city == reqDistributorInfo.city {
				isMatch = false
			}
		}
	}

	return isMatch
}

// ReadCsvFileInfo ... reading the distributors information and storing it into one map
func ReadCsvFileInfo(fileName string) (map[string][]DistributorInfo, error) {
	file, err := os.Open(fileName) //open csv file
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	fileInfo := enrichFileInformationToMap(records)
	return fileInfo, nil
}

// enrichFileInformationToMap ...
func enrichFileInformationToMap(records [][]string) map[string][]DistributorInfo {
	fileInfo := make(map[string][]DistributorInfo)
	for i, record := range records {
		//skip the header
		if i == 0 {
			continue
		}
		distributorInfo := fetchRegionInformation(record[2])
		CheckDistributorIsIncludedOrNot(record[1], &distributorInfo)

		listOfDistributors := strings.Split(record[0], "<")
		//adding the distributors info to the fileInfo map
		for _, distributor := range listOfDistributors {
			fileInfo[distributor] = append(fileInfo[distributor], distributorInfo)
		}
	}
	return fileInfo
}

// CheckDistributorIsIncludedOrNot ...
func CheckDistributorIsIncludedOrNot(record string, distributorInfo *DistributorInfo) {
	if strings.ToUpper(record) == "INCLUDE" {
		distributorInfo.include = true
	} else if strings.ToUpper(record) == "EXCLUDE" {
		distributorInfo.exclude = true
	}
}

// fetchRegionInformation ... fetching the country, state,city details from region info
func fetchRegionInformation(region string) DistributorInfo {
	distributorInfo := DistributorInfo{}
	parts := strings.Split(region, "-")
	city, state, country := "", "", ""
	switch len(parts) {
	case 3:
		city, state, country = parts[0], parts[1], parts[2]
	case 2:
		state, country = parts[0], parts[1]
	default:
		country = parts[0]
	}
	distributorInfo.city = city
	distributorInfo.country = country
	distributorInfo.state = state
	distributorInfo.cityStateCountryCount = len(parts)
	return distributorInfo
}
