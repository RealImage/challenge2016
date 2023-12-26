// Perform actions related to distributor
package distributor

import (
	"example.com/realimage_2016/constants"
	"example.com/realimage_2016/logger"
	"example.com/realimage_2016/validate"

	"strings"
)


// IsAuthorized function checks if the distributor has permission in the region
// distributor - distributor struct for whom we need to check permission
// region      - region where we need to check permission 
func IsAuthorized(distributor constants.Distributor, region string) bool {

	// To store logs inside log file
	log, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}
	defer log.Close()

	currentDistributor := distributor
	currentAddr := &currentDistributor
	regionList := strings.Split(region, "-")
	regionLen := len(regionList)

	if !validate.IsValidRegion(region) {
		log.Log("Error while checking distributor has permission. Invalid region - ", region)
		return false
	}

	for currentAddr != nil {
		switch regionLen {
		case 1:
		// To check if the distributor has permission for whole country
			// Checking if the country code is present in Include
			if incProvinces, ok := currentDistributor.Permissions.Include[regionList[0]]; ok {
				// If incProvinces is not empty, then we can assume distributor don't have permission
				// for the whole country
				if incProvinces != nil {
					return false
				}
				// If the country code is present in Exclude, then distributor don't have permission
				// for whole country or part of the country
				if _, ok := currentDistributor.Permissions.Exclude[regionList[0]]; ok {
					return false
				}
			} else {
				return false
			}

		case 2:
		// To check if the distributor has permission for the province inside the country
			
			// Checking if the country code is present in Include
			if incProvinces, ok := currentDistributor.Permissions.Include[regionList[1]]; ok {
				// If incProvinces is empty, then we can assume distributor have permission
				// for the whole country
				if incProvinces != nil {
					if incCities, ok := incProvinces[regionList[0]]; ok {
						// If incCities is not emty then distributor don't have permission to whole province
						if len(incCities) != 0 {
							return false
						}
					} else {
						return false
					}
				}
			} else {
				return false
			}

			// Checking if country code and province is part of exclude, if yes
			// Distributor don't have permission
			if excProvinces, ok := currentDistributor.Permissions.Exclude[regionList[1]]; ok {
				// If excProvinces is empty then distributor don't have permission to entire country
				if excProvinces != nil {
					if _, ok := excProvinces[regionList[0]]; ok {
						return false
					}
				} else {
					return false
				}
			}

		case 3:
		// To check if the distributor has permission for the city inside the country
			
			// Checking if the country code is present in Include
			if incProvinces, ok := currentDistributor.Permissions.Include[regionList[2]]; ok {
				// If incProvinces is empty, then we can assume distributor have permission
				// for the whole country
				if incProvinces != nil {
					if incCities, ok := incProvinces[regionList[1]]; ok {
						// If incCities is not emty then distributor don't have permission to whole province
						if incCities != nil {
							isCityIncluded := false
							for _,city := range incCities {
								if regionList[0] == city {
									isCityIncluded = true
									break
								}
							}
							// City is not found inside Include, so ditributor don't have permission to the city
							if !isCityIncluded {
								return false
							}
						}
					} else {
						return false
					}
				}
			} else {
				return false
			}

			// Checking if country code and province is part of exclude, if yes
			// Distributor don't have permission
			if excProvinces, ok := currentDistributor.Permissions.Exclude[regionList[2]]; ok {
				// If excProvinces is empty then distributor don't have permission to entire country
				if excProvinces != nil {
					if excCities, ok := excProvinces[regionList[1]]; ok {
						// Don't have permission to entire province
						if excCities == nil {
							return false
						} else {
							for _, city := range excCities {
								// Don't have permission to this city
								if regionList[0] == city {
									return false
								}
							}
						}
					}
				} else {
					return false
				}
			}
		}
		if currentDistributor.Parent != nil {
			currentAddr = currentDistributor.Parent
			currentDistributor = *currentDistributor.Parent
		} else {
			currentAddr = nil
		}
	}
	return true
}

// addRegion function adds the region to inlude/exclude map if the region is valid
// region    - Region to be added to map
// regionMap - Map where region will be added if valid
func addRegion(region string, regionMap map[string]map[string][]string) {

	// To store logs inside log file
	log, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}
	defer log.Close()
	
	// Checking if the region is valid by comparing with the data in CSV file
	if !validate.IsValidRegion(region) {
		log.Log("Error: Invalid region - ", region)
		return
	}
	regionList := strings.Split(region, "-")
	regionLen := len(regionList)
	switch regionLen {
	case 1:
	// If only country code is provided
		// If the key is not present create new entry
		if _, ok := regionMap[regionList[0]]; !ok {
			regionMap[regionList[0]] = nil
		} 
	case 2:
	// If country code and province code provided
		// If the key is not present create new entry
		if provinces, ok := regionMap[regionList[1]]; !ok || provinces == nil {
			regionMap[regionList[1]] = make(map[string][]string)
			regionMap[regionList[1]][regionList[0]] = nil
		}  else {
			if _, ok := regionMap[regionList[1]][regionList[0]]; !ok {
				regionMap[regionList[1]][regionList[0]] = nil
			}
		}
	case 3:
	// If country code, province code and city code provided
		// If the key is not present create new entry
		if provinces, ok := regionMap[regionList[2]]; !ok || provinces == nil {
			regionMap[regionList[2]] = make(map[string][]string)
			regionMap[regionList[2]][regionList[1]] = []string{regionList[0]}
		}  else {
			if cities, ok := regionMap[regionList[2]][regionList[1]]; !ok || cities == nil{
				regionMap[regionList[2]][regionList[1]] = []string{regionList[0]}
			} else {
				regionMap[regionList[2]][regionList[1]] = append(regionMap[regionList[2]][regionList[1]], regionList[0])
			}
		}
	}
}

// AddDistributor functions adds the distributor details to struct
// Id      - Id of the distributor
// include - List of regions where the distributor has permission
// exclude - List of regions where the distributor don't have permission
// parent  - Address of parent distributor
func AddDistributor(Id string, include []string, exclude[]string, parent *constants.Distributor) constants.Distributor {
	var distributor constants.Distributor
	var permission constants.Permissions
	includeMap := make(map[string]map[string][]string)
	excludeMap := make(map[string]map[string][]string)

	// Adding region to include and exclude map
	for _, region := range include {
		addRegion(region, includeMap)
	}
	for _, region := range exclude {
		addRegion(region, excludeMap)
	}

	permission.Include = includeMap
	permission.Exclude = excludeMap

	distributor.Id = Id
	distributor.Permissions = permission
	if parent != nil {
		distributor.Parent = parent
	}

	return distributor
}