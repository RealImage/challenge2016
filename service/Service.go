package service

import (
	"encoding/csv"
	"fmt"
	"golang/dto"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Service interface {
	Distributors(req []dto.DistributorRequest) ([]dto.DistributorResponse, []dto.ValidateResponse)
}

type DefaultService struct{}

// This is the main function which is responsible for returning the response or errors to the handler layer.
func (s DefaultService) Distributors(req []dto.DistributorRequest) ([]dto.DistributorResponse, []dto.ValidateResponse) {
	errorsArr := make([]dto.ValidateResponse, 0)

	// The working directory path is obtained.
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		errorsArr = append(errorsArr, dto.ValidateResponse{Code: "400", Message: err.Error()})
		return nil, errorsArr
	}
	filePath := filepath.Join(dir, "service", "cities.csv")

	//ReadCSV function is called by sending the csv file's path as request & errors are returned if any.
	allLocations, err := ReadCSV(filePath)
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		errorsArr = append(errorsArr, dto.ValidateResponse{Code: "400", Message: err.Error()})
		return nil, errorsArr
	}

	// Once the data is retrieved from the CSV file, it is sent to the validateRequest
	// function along with the request body.
	finalResponse, errValidationResp := validateRequest(req, allLocations)
	if len(errValidationResp) > 0 {
		return nil, errValidationResp
	}

	return finalResponse, nil
}

// This function is to read the CSV file and return the city,province & country data from the CSV file.
func ReadCSV(filePath string) ([]dto.Location, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var locationsArray []dto.Location
	reader := csv.NewReader(file)

	// The csv data is read once to skip the headers before storing data in the array.
	if _, err := reader.Read(); err != nil {
		return nil, err
	}
	for {
		row, err := reader.Read()
		if err != nil {
			break
		}

		if row[0] != "" {
			location := dto.Location{
				City:     row[3],
				Province: row[4],
				Country:  row[5],
			}
			locationsArray = append(locationsArray, location)
		}
	}

	return locationsArray, nil
}

// In this function, all sorts of validations take place & errors are stored in an array and returned.
// In case of no errors, the response object is formed & returned.
func validateRequest(req []dto.DistributorRequest, allLocations []dto.Location) ([]dto.DistributorResponse, []dto.ValidateResponse) {
	errorsArray := make([]dto.ValidateResponse, 0)
	response := make([]dto.DistributorResponse, 0)

	// This is a helper function that compares 2 location objects & returns a boolean response
	isRegionMatch := func(location dto.Location, region dto.Location) bool {
		if region.Country != "" && !strings.EqualFold(region.Country, location.Country) {
			return false
		}
		if region.Province != "" && !strings.EqualFold(region.Province, location.Province) {
			return false
		}
		if region.City != "" && !strings.EqualFold(region.City, location.City) {
			return false
		}
		return true
	}

	// This helper function is used to check if the given location is in the excluded locations list and returns a boolean response.
	isLocationExcluded := func(location dto.Location, exclude []dto.Location) bool {
		for _, region := range exclude {
			if isRegionMatch(location, region) {
				return true
			}
		}
		return false
	}

	// This helper function is used to check if the given location is in the included locations list and returns a boolean response.
	isLocationIncluded := func(location dto.Location, include []dto.Location, exclude []dto.Location) bool {
		if isLocationExcluded(location, exclude) {
			return false
		}

		for _, region := range include {
			if isRegionMatch(location, region) {
				return true
			}
		}
		return false
	}

	// This helper function checks if the region given in the request is present in the data from CSV file and returns a boolean response.
	isRegionPresent := func(region dto.Location) bool {
		for _, loc := range allLocations {
			if isRegionMatch(loc, region) {
				return true
			}
		}
		return false
	}

	// This function is to get the distributor number to ensure the
	// subset calculation works properly even if the distributors are not in sequential order.
	getDistributorNumber := func(distributor string) int {
		num := distributor[len(distributor)-1:]
		number, _ := strconv.Atoi(num)
		return number
	}

	// This function compares the properties of the given two regions and returns a boolean response.
	isRegionSubset := func(regionA, regionB dto.Location) bool {
		if regionA.Country != "" && regionA.Country != regionB.Country {
			return false
		}
		if regionA.Province != "" && regionA.Province != regionB.Province {
			return false
		}
		if regionA.City != "" && regionA.City != regionB.City {
			return false
		}
		return true
	}

	// This helper function is used to check if a given region is a subset of the other region and returns a boolean response.
	isSubset := func(distributorA, distributorB dto.DistributorRequest) bool {
		// Check the distributor number in their names
		numA := getDistributorNumber(distributorA.Distributor)
		numB := getDistributorNumber(distributorB.Distributor)
		if numA >= numB {
			return false
		}

		// Check if all include regions of distributorB are present in distributorA
		for _, includeB := range distributorB.Include {
			found := false
			for _, includeA := range distributorA.Include {
				if isRegionSubset(includeB, includeA) {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}

		// Check if all exclude regions of distributorB are present in distributorA
		for _, excludeB := range distributorB.Exclude {
			found := false
			for _, excludeA := range distributorA.Exclude {
				if isRegionSubset(excludeB, excludeA) {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}

		return true
	}

	sort.SliceStable(req, func(i, j int) bool {
		return getDistributorNumber(req[i].Distributor) < getDistributorNumber(req[j].Distributor)
	})

	// Chekcs if the object is a subset of the previous object and returns a boolean response.
	isSubsetOfPrevious := func(index int) (bool, string) {
		currentDistributor := req[index]

		for i := index - 1; i >= 0; i-- {
			previousDistributor := req[i]
			if isSubset(previousDistributor, currentDistributor) {
				return true, "" // It's a subset of the previous distributor.
			}

			// Check if any region in currentDistributor's Include array is not present in previousDistributor's Include array.
			for _, includeRegion := range currentDistributor.Include {
				found := false
				for _, prevIncludeRegion := range previousDistributor.Include {
					if isRegionMatch(includeRegion, prevIncludeRegion) {
						found = true
						break
					}
				}
				if !found {
					return false, previousDistributor.Distributor // Not a subset, return the distributor causing the error.
				}
			}
		}

		return false, "" // Not a subset of any previous distributor.
	}

	// Compares the objects from include & exclude & returns an error if both are exacty same.
	hasConflictingPermissions := func(include []dto.Location, exclude []dto.Location) bool {
		for _, includeRegion := range include {
			for _, excludeRegion := range exclude {
				if isRegionMatch(includeRegion, excludeRegion) {
					return true
				}
			}
		}
		return false
	}

	for i, request := range req {
		if hasConflictingPermissions(request.Include, request.Exclude) {
			errorsArray = append(errorsArray, dto.ValidateResponse{Code: "400", Message: "Same location cannot be used in both include and exclude for " + request.Distributor})
		}

		// Validates the locations in include array and returns error if any.
		for _, include := range request.Include {
			if !isRegionPresent(include) {
				errorsArray = append(errorsArray, dto.ValidateResponse{Code: "400", Message: "One of the included regions is not found in the database for " + request.Distributor})
			}
		}

		// Validates the locations in exclude array and returns error if any.
		for _, exclude := range request.Exclude {
			if !isRegionPresent(exclude) {
				errorsArray = append(errorsArray, dto.ValidateResponse{Code: "400", Message: "One of the excluded regions is not found in the database for " + request.Distributor})
			}
		}

		// Validates the locations in Locations array and returns error if any.
		// Also checks if the locations can be distributed by the distributor and returns the response accordingly.
		for _, location := range request.Locations {
			if !isRegionPresent(location) {
				errorsArray = append(errorsArray, dto.ValidateResponse{Code: "400", Message: "One of the selected distribution regions is not included for " + request.Distributor})
			}

			if isLocationIncluded(location, request.Include, request.Exclude) {
				if location.City != "" && location.Province != "" && location.Country != "" {
					response = append(response, dto.DistributorResponse{Distributor: request.Distributor, Permission: request.Distributor + " can distribute in " + location.City + "-" + location.Province + "-" + location.Country})
				} else if location.City == "" && location.Province != "" && location.Country != "" {
					response = append(response, dto.DistributorResponse{Distributor: request.Distributor, Permission: request.Distributor + " can distribute in " + location.Province + "-" + location.Country})
				} else if location.City == "" && location.Province == "" && location.Country != "" {
					response = append(response, dto.DistributorResponse{Distributor: request.Distributor, Permission: request.Distributor + " can distribute in " + location.Country})
				}
			} else {
				if location.City != "" && location.Province != "" && location.Country != "" {
					response = append(response, dto.DistributorResponse{Distributor: request.Distributor, Permission: request.Distributor + " cannot distribute in " + location.City + "-" + location.Province + "-" + location.Country})
				} else if location.City == "" && location.Province != "" && location.Country != "" {
					response = append(response, dto.DistributorResponse{Distributor: request.Distributor, Permission: request.Distributor + " cannot distribute in " + location.Province + "-" + location.Country})
				} else if location.City == "" && location.Province == "" && location.Country != "" {
					response = append(response, dto.DistributorResponse{Distributor: request.Distributor, Permission: request.Distributor + " cannot distribute in " + location.Country})
				}
			}
		}

		_, nonSubsetDistributor := isSubsetOfPrevious(i)
		if nonSubsetDistributor != "" {
			errorsArray = append(errorsArray, dto.ValidateResponse{
				Code:    "400",
				Message: "The " + request.Distributor + " should be a subset of " + nonSubsetDistributor,
			})
		}

	}

	if len(errorsArray) > 0 {
		return nil, errorsArray
	}

	return response, nil
}
