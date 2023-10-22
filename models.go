package main

// Define distributor hierarchy structure
type Distributor struct {
	Name         string
	Parent       string
	Permissions  Permission
}

// Define distributor permission structure
type Permission struct {
	Include []string
	Exclude []string
}

// Define region data structure
type Region struct {
	CityCode     string
	ProvinceCode string
	CountryCode  string
	CityName     string
	ProvinceName string
	CountryName  string
}


var distributorPermissions = make(map[string]Permission)
var distributors = make(map[string]Distributor)

