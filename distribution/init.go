package distribution

import (
	"flag"
)

type CityPermissionMatrix map[string]bool
type ProvincePermissionMatrix map[string]CityPermissionMatrix
type PermissionMatrix map[string]ProvincePermissionMatrix

var filePath string

func init() {
	flag.StringVar(&filePath, "file", "cities.csv", "Cities list file")
}

func loadBasePermissions() PermissionMatrix {
	cities, err := LoadCitiesFromCSV(filePath)
	if err != nil {
		panic(err)
	}
	basePermissions := make(PermissionMatrix)
	for _, city := range cities {
		_, isCountryThere := basePermissions[city.Country.Code]
		if !isCountryThere {
			basePermissions[city.Country.Code] = make(ProvincePermissionMatrix)
		}
		_, isProvinceThere := basePermissions[city.Country.Code][city.Province.Code]
		if !isProvinceThere {
			basePermissions[city.Country.Code][city.Province.Code] = make(CityPermissionMatrix)
		}
		basePermissions[city.Country.Code][city.Province.Code][city.Code] = false
	}
	return basePermissions
}
