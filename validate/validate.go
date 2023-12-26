// Validates if given region is valid
package validate

import (
	"example.com/realimage_2016/constants"

	"strings"
)

// IsValidRegion function validates if a region is valid by comparing with the data collected from CSV file
func IsValidRegion(region string) bool {
	strList := strings.Split(region, "-")
	switch len(strList) {
	case 1:
	// If only country code is provided
		if _, ok := constants.RegionData[strList[0]]; ok {
			return true
		}
	case 2:
	// If country code and province code provided
		if _, ok := constants.RegionData[strList[1]]; ok {
			if _, ok := constants.RegionData[strList[1]][strList[0]]; ok {
				return true
			}
		}
	case 3:
	// If country code, province code and city code provided
		if _, ok := constants.RegionData[strList[2]]; ok {
			if _, ok := constants.RegionData[strList[2]][strList[1]]; ok {
				for _,city := range constants.RegionData[strList[2]][strList[1]] {
					if city == strList[0] {
						return true
					}
				}
			}
		}
	default:
		return false
	}
	return false
}