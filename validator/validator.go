package validator

import (
	"challenge2016/dto"
	"challenge2016/permission"
	"strings"
)

// The function `ValidateDistributorData` validates distributor data by checking for empty fields,
// duplicate names, and valid regions.
func ValidateDistributorData(data dto.Distributor, groupedData []dto.Country, distributorInformation []dto.Distributor) []string {
	var errorMsg []string

	if strings.TrimSpace(data.Name) == "" {
		errorMsg = append(errorMsg, "Distributor Name must not be empty, please enter a valid distributor name")
	} else if ValidateDistributorName(strings.ToUpper(strings.TrimSpace(data.Name)), distributorInformation) {
		errorMsg = append(errorMsg, "Distributor Name already exists")
	}

	if len(data.Include) == 0 {
		errorMsg = append(errorMsg, "Include Regions must not be empty, please enter valid regions")
	} else {
		for _, region := range data.Include {
			if !ValidateRegion(region, groupedData) {
				errorMsg = append(errorMsg, "Include Region '"+region+"' is not present in csv, please enter a valid region")
			}
		}
	}

	if len(data.Exclude) > 0 {
		for _, region := range data.Exclude {
			if !ValidateRegion(region, groupedData) {
				errorMsg = append(errorMsg, "Exclude Region '"+region+"' is not present in csv, please enter a valid region")
			}
		}
	}

	if len(data.Exclude) > 0 && len(data.Include) > 0 {
		for _, ExcludeRegion := range data.Exclude {
			for _, IncludeRegion := range data.Include {
				if strings.EqualFold(ExcludeRegion, IncludeRegion) {
					errorMsg = append(errorMsg, "Exclude Region and Include Region should not be Same, please enter a valid region")
				}

			}

		}
	}

	return errorMsg
}

// The function `ValidateSubDistributorData` validates the data of a sub-distributor, checking for
// errors such as empty fields, duplicate names, invalid regions, and incorrect parent distributor
// name.
func ValidateSubDistributorData(data dto.Distributor, groupedData []dto.Country, distributorInformation []dto.Distributor) []string {
	var errorMsg []string

	if strings.TrimSpace(data.Name) == "" {
		errorMsg = append(errorMsg, "Distributor Name must not be empty, please enter a valid distributor name")
	} else if ValidateDistributorName(strings.ToUpper(strings.TrimSpace(data.Name)), distributorInformation) {
		errorMsg = append(errorMsg, "Distributor Name already exists")
	}

	if len(data.Include) == 0 {
		errorMsg = append(errorMsg, "Include Regions must not be empty, please enter valid regions")
	} else {
		for _, region := range data.Include {
			if !ValidateRegion(region, groupedData) {
				errorMsg = append(errorMsg, "Include Region '"+region+"' is not present in csv, please enter a valid region")
			}
		}
	}

	if len(data.Exclude) > 0 {
		for _, region := range data.Exclude {
			if !ValidateRegion(region, groupedData) {
				errorMsg = append(errorMsg, "Exclude Region '"+region+"' is not present in csv, please enter a valid region")
			}
		}
	}

	if len(data.Exclude) > 0 && len(data.Include) > 0 {
		for _, ExcludeRegion := range data.Exclude {
			for _, IncludeRegion := range data.Include {
				if strings.EqualFold(ExcludeRegion, IncludeRegion) {
					errorMsg = append(errorMsg, "Exclude Region and Include Region should not be Same, please enter a valid region")
				}

			}

		}
	}

	if strings.TrimSpace(data.Parent) == "" {
		errorMsg = append(errorMsg, "Parent distributor Name must not be empty, please enter a valid parent distributor name")
	} else if !ValidateDistributorName(strings.ToUpper(strings.TrimSpace(data.Parent)), distributorInformation) {
		errorMsg = append(errorMsg, "Parent distributor Name does not exist, please enter an existing parent distributor name")
	}

	if len(errorMsg) == 0 {
		testData := append(data.Include, data.Exclude...)
		checkPermissionWithParent := permission.CheckPermission(strings.ToUpper(strings.TrimSpace(data.Parent)), testData, "subDistributionCreation", distributorInformation)
		if len(checkPermissionWithParent) > 0 {
			errorMsg = append(errorMsg, checkPermissionWithParent...)
		}
	}

	return errorMsg
}

// The function "ValidateDistributorName" checks if a given distributor name exists in a list of
// distributor information.
func ValidateDistributorName(distributorName string, distributorInformation []dto.Distributor) bool {
	for _, distributor := range distributorInformation {
		if distributor.Name == distributorName {
			return true
		}
	}
	return false
}

// The function `ValidateRegion` checks if a given region is valid based on a list of grouped data.
func ValidateRegion(data string, groupedData []dto.Country) bool {
	splitTestData := strings.Split(data, ",")
	for _, region := range splitTestData {
		testData := strings.Split(strings.ToUpper(region), "-")

		if len(testData) > 0 && len(testData) < 4 {
			if len(testData) == 1 {
				for _, country := range groupedData {
					if country.Name == testData[0] {
						return true
					}
				}
				return false
			} else if len(testData) == 2 {
				for _, country := range groupedData {
					if country.Name == testData[1] {
						for _, state := range country.States {
							if state.Name == testData[0] {
								return true
							}
						}
					}
				}
				return false
			} else if len(testData) == 3 {
				for _, country := range groupedData {
					if country.Name == testData[2] {
						for _, state := range country.States {
							if state.Name == testData[1] {
								for _, city := range state.Cities {
									if city.Name == testData[0] {
										return true
									}
								}
							}
						}
					}
				}
				return false
			}
		} else {
			return false
		}
	}
	return false
}

// The function `ValidateCheckPermissionData` validates the `CheckPermissionData` object by checking if
// the distributor name is not empty and exists in the distributor information, and if all the regions
// in the data exist in the grouped data.
func ValidateCheckPermissionData(data dto.CheckPermissionData, groupedData []dto.Country, distributorInformation []dto.Distributor) []string {
	var errorMsg []string

	if strings.TrimSpace(data.DistributorName) == "" {
		errorMsg = append(errorMsg, "Distributor Name must not be empty, please enter a valid distributor name")
	} else if !ValidateDistributorName(strings.ToUpper(strings.TrimSpace(data.DistributorName)), distributorInformation) {
		errorMsg = append(errorMsg, "Distributor name does not exist")
	}

	for _, region := range data.Regions {
		if !ValidateRegion(region, groupedData) {
			errorMsg = append(errorMsg, strings.ToUpper(region)+" does not exist in the csv file, please enter a valid region")
		}
	}

	return errorMsg
}
