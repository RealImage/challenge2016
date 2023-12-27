package main

// Primary building block - Distributor hierarchy structure
type Distributor struct {
	Name        string
	Parent      string
	Permissions Permission
}

// Permission DS
type Permission struct {
	Include []string
	Exclude []string
}

// Wrapper with full path - Region DS
type Region struct {
	CityCode     string
	ProvinceCode string
	CountryCode  string
	CityName     string
	ProvinceName string
	CountryName  string
}
