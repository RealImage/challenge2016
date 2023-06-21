package models


// DB struct
type DistributionMaps struct{
	CityMap map[string]*Location
	ProvinceMap map[string]*Location
	CountryMap map[string]*Location
}

// factory function
func NewDistributionMaps() *DistributionMaps{
	return &DistributionMaps{
		CityMap: make(map[string]*Location),
		ProvinceMap: make(map[string]*Location),
		CountryMap: make(map[string]*Location),
	}
}