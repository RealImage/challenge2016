package service

import (
	"errors"
	"task/common"
)

type Service struct {
}

func removeDuplicates(input []common.LocationIdentifier) []common.LocationIdentifier {
	encountered := map[common.LocationIdentifier]bool{}
	result := []common.LocationIdentifier{}

	for _, v := range input {
		if !encountered[v] {
			encountered[v] = true
			result = append(result, v)
		}
	}
	return result
}

func AddPermission(d *common.Distributor, locations []common.LocationIdentifier, include bool) {
	if include {
		d.Includes = append(d.Includes, locations...)
		d.Includes = removeDuplicates(d.Includes)
	} else {
		d.Excludes = append(d.Excludes, locations...)
		d.Excludes = removeDuplicates(d.Excludes)
	}
}

func ValidateAndGetLocations(input common.Location) (common.LocationIdentifier, error) {
	var ok bool
	var errorStr string
	if input.CountryCode != "" && input.ProvinceCode != "" && input.CityCode != "" {
		_, ok = common.CityMap[input.CountryCode+"$"+input.ProvinceCode+"$"+input.CityCode]
		errorStr = "Country code: " + input.CountryCode + ", Province code: " + input.ProvinceCode + ", City code: " + input.CityCode
	} else if input.CountryCode != "" && input.ProvinceCode != "" {
		_, ok = common.ProvinceMap[input.CountryCode+"$"+input.ProvinceCode]
		errorStr = "Country code: " + input.CountryCode + ", Province code: " + input.ProvinceCode
	} else if input.CountryCode != "" {
		_, ok = common.CountryMap[input.CountryCode]
		errorStr = "Country code: " + input.CountryCode
	}
	if !ok {
		return common.LocationIdentifier{}, errors.New("Location Doesnot exists base on the details provided by you. " + errorStr)
	}
	return common.LocationIdentifier{CountryCode: input.CountryCode, ProvinceCode: input.ProvinceCode, CityCode: input.CityCode}, nil
}

func (h *Service) AddDistributorAndPermissions(input common.DistributorInput) common.Response {
	var distributor *common.Distributor
	if common.DistributorsMap[input.Distributorname] == nil {
		distributor = common.NewDistributor(input.Distributorname, input.ParentDistributorName)
	} else {
		distributor = common.DistributorsMap[input.Distributorname]
	}
	IncludeLocations := make([]common.LocationIdentifier, 0)
	ExcludeLocations := make([]common.LocationIdentifier, 0)

	for _, include := range input.Includes {
		loc, err := ValidateAndGetLocations(include)
		if err != nil {
			return common.Response{Msg: err.Error(), Status: 400}
		}
		IncludeLocations = append(IncludeLocations, loc)
	}
	for _, exclude := range input.Excludes {
		loc, err := ValidateAndGetLocations(exclude)
		if err != nil {
			return common.Response{Msg: err.Error(), Status: 400}
		}
		ExcludeLocations = append(ExcludeLocations, loc)
	}
	AddPermission(distributor, IncludeLocations, true)
	AddPermission(distributor, ExcludeLocations, false)

	return common.Response{Msg: "Created Distributor", Data: input, Status: 201}

}

func (h *Service) GetAccessDetailsOfDistributor(input common.LocationAccessInput) common.Response {

	if common.DistributorsMap[input.Distributorname] == nil {
		return common.Response{Msg: "distributor does not exists", Data: input, Status: 400}
	}
	distributor := common.DistributorsMap[input.Distributorname]

	location, err := ValidateAndGetLocations(common.Location{
		CountryCode:  input.CountryCode,
		ProvinceCode: input.ProvinceCode,
		CityCode:     input.CityCode,
	})
	if err != nil {
		return common.Response{Msg: err.Error(), Status: 400}
	}

	exists := CheckPermission(distributor, location)
	var msg string
	if exists {
		msg = "Has access"
	} else {
		msg = "Does not have access"
	}
	return common.Response{Msg: msg, Data: input, Status: 200}
}

func CheckPermission(d *common.Distributor, location common.LocationIdentifier) bool {
	for _, exclude := range d.Excludes {
		if match(exclude, location) {
			return false
		}
	}

	directInclusion := false
	for _, include := range d.Includes {
		if match(include, location) {
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

func match(perm common.LocationIdentifier, loc common.LocationIdentifier) bool {
	if loc.CountryCode != "" && perm.CountryCode != loc.CountryCode {
		return false
	}
	if loc.ProvinceCode != "" && perm.ProvinceCode != "" && perm.ProvinceCode != loc.ProvinceCode {
		return false
	}
	if loc.CityCode != "" && perm.CityCode != "" && perm.CityCode != loc.CityCode {
		return false
	}
	return true
}

func CheckExclusion(d *common.Distributor, location common.LocationIdentifier) bool {
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

func (h *Service) GetDistributorDetails(name string) common.Response {

	if common.DistributorsMap[name] == nil {
		return common.Response{Msg: "Distributor " + name + " does not exists", Status: 400}
	}
	return common.Response{Msg: "Distributor details fetched succesfully", Data: common.DistributorsMap[name], Status: 200}
}
