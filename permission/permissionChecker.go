package permission

import (
	"challenge2016/dto" // Importing DTO package for data transfer objects
	"strings"
)

// The CheckPermission function checks if a distributor has access to certain test data based on their
// inclusion and exclusion lists.
func CheckPermission(distributorName string, testData []string, origin string, distributorInformation []dto.Distributor) []string {
	var validationResult []string
	var errorMsg []string

	// Get distributor data by name
	var distributorData dto.Distributor
	for _, distributor := range distributorInformation {
		if distributor.Name == strings.ToUpper(distributorName) {
			distributorData = distributor
			break
		}
	}

	for _, data := range testData {
		switch len(strings.Split(data, "-")) {
		case 1:
			if contains(distributorData.Include, data) {
				validationResult = append(validationResult, distributorData.Name+" has access to "+data)
			} else {
				validationResult = append(validationResult, distributorData.Name+" does not have access to "+data)
				errorMsg = append(errorMsg, distributorData.Name+" does not have access to "+data)
			}
		case 2:
			newTestData := strings.Split(data, "-")
			if contains(distributorData.Include, newTestData[1]) {
				if contains(distributorData.Exclude, data) {
					validationResult = append(validationResult, distributorData.Name+" does not have access to "+data)
					errorMsg = append(errorMsg, distributorData.Name+" does not have access to "+data)
				} else {
					validationResult = append(validationResult, distributorData.Name+" has access to "+data)
				}
			} else if contains(distributorData.Include, data) {
				validationResult = append(validationResult, distributorData.Name+" has access to "+data)
			} else {
				validationResult = append(validationResult, distributorData.Name+" does not have access to "+data)
				errorMsg = append(errorMsg, distributorData.Name+" does not have access to "+data)
			}
		case 3:
			newTestData := strings.Split(data, "-")
			if contains(distributorData.Include, newTestData[2]) {
				if contains(distributorData.Include, newTestData[1]+"-"+newTestData[2]) {
					if contains(distributorData.Include, data) {
						validationResult = append(validationResult, distributorData.Name+" has access to "+data)
					} else {
						if contains(distributorData.Exclude, data) {
							validationResult = append(validationResult, distributorData.Name+" does not have access to "+data)
							errorMsg = append(errorMsg, distributorData.Name+" does not have access to "+data)
						} else {
							validationResult = append(validationResult, distributorData.Name+" has access to "+data)
						}
					}
				} else {
					if contains(distributorData.Exclude, newTestData[1]+"-"+newTestData[2]) {
						validationResult = append(validationResult, distributorData.Name+" does not have access to "+data)
						errorMsg = append(errorMsg, distributorData.Name+" does not have access to "+data)
					} else {
						if contains(distributorData.Exclude, data) {
							validationResult = append(validationResult, distributorData.Name+" does not have access to "+data)
							errorMsg = append(errorMsg, distributorData.Name+" does not have access to "+data)
						} else {
							validationResult = append(validationResult, distributorData.Name+" has access to "+data)
						}
					}
				}
			} else {
				if contains(distributorData.Include, newTestData[1]+"-"+newTestData[2]) {
					if contains(distributorData.Exclude, data) {
						validationResult = append(validationResult, distributorData.Name+" does not have access to "+data)
						errorMsg = append(errorMsg, distributorData.Name+" does not have access to "+data)
					} else {
						validationResult = append(validationResult, distributorData.Name+" has access to "+data)
					}
				} else if contains(distributorData.Include, data) {
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

// contains checks if a string is present in a slice of strings
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
