package service

import (
	"strings"

	"github.com/saurabh-sde/challenge2016_saurabh/utils"
)

func ValidateLocation(includes, excludes []string) bool {
	for _, c := range includes {
		found := false
		locStr := strings.Split(c, "-")
		if len(locStr) == 3 {
			_, cityFound := utils.CityMap[locStr[2]]
			_, provinceFound := utils.ProvinceMap[locStr[1]]
			_, countryFound := utils.CountryMap[locStr[0]]
			if !cityFound || !provinceFound || !countryFound {
				return found
			}
		} else if len(locStr) == 2 {
			_, provinceFound := utils.ProvinceMap[locStr[1]]
			_, countryFound := utils.CountryMap[locStr[0]]
			if !provinceFound || !countryFound {
				return found
			}
		} else if len(locStr) == 1 {
			_, countryFound := utils.CountryMap[locStr[0]]
			if !countryFound {
				return found
			}
		}
	}
	for _, c := range excludes {
		found := false
		locStr := strings.Split(c, "-")
		if len(locStr) == 3 {
			_, cityFound := utils.CityMap[locStr[0]]
			_, provinceFound := utils.ProvinceMap[locStr[1]]
			_, countryFound := utils.CountryMap[locStr[3]]
			if !cityFound || !provinceFound || !countryFound {
				return found
			}
		} else if len(locStr) == 2 {
			_, provinceFound := utils.ProvinceMap[locStr[0]]
			_, countryFound := utils.CountryMap[locStr[1]]
			if !provinceFound || !countryFound {
				return found
			}
		} else if len(locStr) == 1 {
			_, countryFound := utils.CountryMap[locStr[0]]
			if !countryFound {
				return found
			}
		}
	}
	return true
}
