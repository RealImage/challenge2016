package utils

import (
	"distributor/types"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

var DistributorsMap = make(map[string]*types.Distributor)
var CountryMap = make(map[string]types.LocationIdentifier)
var ProvinceMap = make(map[string]types.LocationIdentifier)
var CityMap = make(map[string]types.LocationIdentifier)

func LoadLocations(filePath string) error {
	fmt.Println("Inside LoadLocations func")
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Csv load failed")
		return err
	}
	defer file.Close()

	fmt.Println("Csv load successfully")

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error while reading scv file")
		return err
	}

	for _, record := range records[1:] {
		loc := NewLocationIdentifier(record[2], record[1], record[0], record[5], record[4], record[3])
		CountryMap[loc.CountryCode] = loc
		ProvinceMap[loc.CountryCode+"$"+loc.ProvinceCode] = loc
		CityMap[loc.CountryCode+"$"+loc.ProvinceCode+"$"+loc.CityCode] = loc
	}
	return nil
}

func NewLocationIdentifier(countryCode, provinceCode, cityCode, countryName, provinceName, cityName string) types.LocationIdentifier {
	return types.LocationIdentifier{
		CountryCode:  countryCode,
		ProvinceCode: provinceCode,
		CityCode:     cityCode,
		CountryName:  countryName,
		ProvinceName: provinceName,
		CityName:     cityName,
	}
}

func NewDistributor(name, parentName string) *types.Distributor {
	var parent *types.Distributor
	if parentName != "" {
		parent = DistributorsMap[parentName]

	}
	dist := types.Distributor{
		Name:     name,
		Parent:   parent,
		Includes: make([]types.LocationIdentifier, 0),
		Excludes: make([]types.LocationIdentifier, 0),
	}
	DistributorsMap[name] = &dist
	return &dist
}

func ValidateAndGetLocations(input types.Location) (types.LocationIdentifier, error) {
	var ok bool
	var errorStr string
	if input.CountryCode != "" && input.ProvinceCode != "" && input.CityCode != "" {
		_, ok = CityMap[input.CountryCode+"$"+input.ProvinceCode+"$"+input.CityCode]
		errorStr = "Country code: " + input.CountryCode + ", Province code: " + input.ProvinceCode + ", City code: " + input.CityCode
	} else if input.CountryCode != "" && input.ProvinceCode != "" {
		_, ok = ProvinceMap[input.CountryCode+"$"+input.ProvinceCode]
		errorStr = "Country code: " + input.CountryCode + ", Province code: " + input.ProvinceCode
	} else if input.CountryCode != "" {
		_, ok = CountryMap[input.CountryCode]
		errorStr = "Country code: " + input.CountryCode
	}
	if !ok {
		return types.LocationIdentifier{}, errors.New("Location Doesnot exists base on the details provided by you. " + errorStr)
	}
	return types.LocationIdentifier{CountryCode: input.CountryCode, ProvinceCode: input.ProvinceCode, CityCode: input.CityCode}, nil
}

func AddPermission(d *types.Distributor, locations []types.LocationIdentifier, include bool) {
	if include {
		d.Includes = append(d.Includes, locations...)
		d.Includes = removeDuplicates(d.Includes)
	} else {
		d.Excludes = append(d.Excludes, locations...)
		d.Excludes = removeDuplicates(d.Excludes)
	}
}

func removeDuplicates(input []types.LocationIdentifier) []types.LocationIdentifier {
	encountered := map[types.LocationIdentifier]bool{}
	result := []types.LocationIdentifier{}

	for _, v := range input {
		if !encountered[v] {
			encountered[v] = true
			result = append(result, v)
		}
	}
	return result
}

func CheckPermission(d *types.Distributor, location types.LocationIdentifier) bool {
	for _, exclude := range d.Excludes {
		if match(exclude, location) {
			return false
		}
	}

	directInclusion := false
	for _, include := range d.Includes {
		if matchInclusive(include, location) {
			directInclusion = true
			break
		}
	}

	if directInclusion {
		if d.Parent != nil && CheckExclusion(d.Parent, location) {
			return false
		}
		return true
	}

	if d.Parent != nil {
		return CheckPermission(d.Parent, location)
	}

	return false
}

func matchInclusive(perm types.LocationIdentifier, loc types.LocationIdentifier) bool {
	if loc.CountryCode != "" && perm.CountryCode != loc.CountryCode {
		return false
	}
	if (loc.ProvinceCode != "" && perm.ProvinceCode != "") && (perm.ProvinceCode != loc.ProvinceCode) {
		return false
	}
	if (loc.CityCode != "" && perm.CityCode != "") && (perm.CityCode != loc.CityCode) {
		return false
	}
	return true
}

func match(perm types.LocationIdentifier, loc types.LocationIdentifier) bool {
	if loc.CountryCode != "" && perm.CountryCode == loc.CountryCode {
		if loc.ProvinceCode == "" && perm.ProvinceCode == loc.ProvinceCode {
			if loc.CityCode == "" && perm.CityCode == loc.CityCode {
				return true
			}
		}
	}
	if loc.CountryCode != "" && perm.CountryCode == loc.CountryCode {
		if loc.ProvinceCode != "" && perm.ProvinceCode == loc.ProvinceCode {
			if loc.CityCode == "" && perm.CityCode == loc.CityCode {
				return true
			}
		}
	}
	if loc.CountryCode != "" && perm.CountryCode == loc.CountryCode {
		if loc.ProvinceCode != "" && perm.ProvinceCode == loc.ProvinceCode {
			if loc.CityCode != "" && perm.CityCode == loc.CityCode {
				return true
			}
		}
	}
	return false
}

func CheckExclusion(d *types.Distributor, location types.LocationIdentifier) bool {
	for _, exclude := range d.Excludes {
		if match(exclude, location) {
			return true
		}
	}
	if d.Parent != nil {
		return CheckExclusion(d.Parent, location)
	}
	return false
}
