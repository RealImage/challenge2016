package permission

import (
	"challenge2016/dto" // Importing DTO package for data transfer objects
	"strings"

	"golang.org/x/exp/slices"
)

// The CheckPermission function checks if a distributor has access to certain test data based on their
// inclusion and exclusion lists.
func CheckPermission(distributorName string, InputData []string, origin string, distributorInformation []dto.Distributor) []string {
	var validationResult []string
	var errorMsg []string

	// Get distributor data by name
	var distributorData dto.Distributor
	for _, distributor := range distributorInformation {
		if strings.EqualFold(distributor.Name, distributorName) {
			distributorData = distributor
			break
		}
	}

	for _, data := range InputData {
		switch len(strings.Split(data, "-")) {
		case 1:
			if slices.Contains(distributorData.Include, data) {
				validationResult = append(validationResult, distributorData.Name+" has access to "+data)
			} else {
				validationResult = append(validationResult, distributorData.Name+" does not have access to "+data)
				errorMsg = append(errorMsg, distributorData.Name+" does not have access to "+data)
			}
		case 2:
			newTestData := strings.Split(data, "-")
			if slices.Contains(distributorData.Include, newTestData[1]) {
				if slices.Contains(distributorData.Exclude, data) {
					validationResult = append(validationResult, distributorData.Name+" does not have access to "+data)
					errorMsg = append(errorMsg, distributorData.Name+" does not have access to "+data)
				} else {
					validationResult = append(validationResult, distributorData.Name+" has access to "+data)
				}
			} else if slices.Contains(distributorData.Include, data) {
				validationResult = append(validationResult, distributorData.Name+" has access to "+data)
			} else {
				validationResult = append(validationResult, distributorData.Name+" does not have access to "+data)
				errorMsg = append(errorMsg, distributorData.Name+" does not have access to "+data)
			}
		case 3:
			newTestData := strings.Split(data, "-")
			if slices.Contains(distributorData.Include, newTestData[2]) {
				if slices.Contains(distributorData.Include, newTestData[1]+"-"+newTestData[2]) {
					if slices.Contains(distributorData.Include, data) {
						validationResult = append(validationResult, distributorData.Name+" has access to "+data)
					} else {
						if slices.Contains(distributorData.Exclude, data) {
							validationResult = append(validationResult, distributorData.Name+" does not have access to "+data)
							errorMsg = append(errorMsg, distributorData.Name+" does not have access to "+data)
						} else {
							validationResult = append(validationResult, distributorData.Name+" has access to "+data)
						}
					}
				} else {
					if slices.Contains(distributorData.Exclude, newTestData[1]+"-"+newTestData[2]) {
						validationResult = append(validationResult, distributorData.Name+" does not have access to "+data)
						errorMsg = append(errorMsg, distributorData.Name+" does not have access to "+data)
					} else {
						if slices.Contains(distributorData.Exclude, data) {
							validationResult = append(validationResult, distributorData.Name+" does not have access to "+data)
							errorMsg = append(errorMsg, distributorData.Name+" does not have access to "+data)
						} else {
							validationResult = append(validationResult, distributorData.Name+" has access to "+data)
						}
					}
				}
			} else {
				if slices.Contains(distributorData.Include, newTestData[1]+"-"+newTestData[2]) {
					if slices.Contains(distributorData.Exclude, data) {
						validationResult = append(validationResult, distributorData.Name+" does not have access to "+data)
						errorMsg = append(errorMsg, distributorData.Name+" does not have access to "+data)
					} else {
						validationResult = append(validationResult, distributorData.Name+" has access to "+data)
					}
				} else if slices.Contains(distributorData.Include, data) {
					validationResult = append(validationResult, distributorData.Name+" has access to "+data)
				} else {
					validationResult = append(validationResult, distributorData.Name+" does not have access to "+data)
					errorMsg = append(errorMsg, distributorData.Name+" does not have access to "+data)
				}
			}
		}
	}

	if origin == "subDistributionCreation" {
		return errorMsg
	}
	return validationResult
}
