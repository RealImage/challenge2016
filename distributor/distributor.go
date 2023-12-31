package distributor

import (
	"fmt"
	"strings"
)

// Distributor represents a distributor object.
type Distributor struct {
	RegionMap map[string]map[string]map[string]bool
	DistributorID int
}

// NewDistributor creates a new Distributor object based on the provided inputs.
func NewDistributor(regionMap map[string]map[string]map[string]bool, includes, excludes []string,id int) (*Distributor, error) {
	newRegionMap := make(map[string]map[string]map[string]bool)
	for _, include := range includes {
		parts := strings.Split(include, "-")
		switch len(parts) {
		case 1:
			// Include country
			country := parts[0]
			if countryMap, exists := regionMap[country]; exists {
				newRegionMap[country] = countryMap
			} else {
				return nil, fmt.Errorf("included country '%s' not found in the original region map", country)
			}
		case 2:
			// Include province-country
			province, country := parts[0], parts[1]
			if countryMap, exists := regionMap[country]; exists {
				if provinceMap, exists := countryMap[province]; exists {
					newRegionMap[country] = map[string]map[string]bool{province: provinceMap}
				} else {
					return nil, fmt.Errorf("included province '%s-%s' not found in the original region map", province, country)
				}
			} else {
				return nil, fmt.Errorf("included country '%s' not found in the original region map", country)
			}
		case 3:
			// Include city-province-country
			city, province, country := parts[0], parts[1], parts[2]
			if countryMap, exists := regionMap[country]; exists {
				if provinceMap, exists := countryMap[province]; exists {
					if cityMap, exists := provinceMap[city]; exists {
						newRegionMap[country] = map[string]map[string]bool{province:{city: cityMap}}
					} else {
						return nil, fmt.Errorf("included city '%s-%s-%s' not found in the original region map", city, province, country)
					}
				} else {
					return nil, fmt.Errorf("included province '%s-%s' not found in the original region map", province, country)
				}
			} else {
				return nil, fmt.Errorf("included country '%s' not found in the original region map", country)
			}
		default:
			return nil, fmt.Errorf("invalid include format: %s", include)
		}
	}

	// Process excludes
	for _, exclude := range excludes {
		parts := strings.Split(exclude, "-")
		switch len(parts) {
		case 1:
			// Exclude country
			country := parts[0]
			delete(newRegionMap, country)
		case 2:
			// Exclude province-country
			province, country := parts[0], parts[1]
			if countryMap, exists := newRegionMap[country]; exists {
				delete(countryMap, province)
				if len(countryMap) == 0 {
					delete(newRegionMap, country)
				}
			}
		case 3:
			// Exclude city-province-country
			city, province, country := parts[0], parts[1], parts[2]
			if countryMap, exists := newRegionMap[country]; exists {
				if provinceMap, exists := countryMap[province]; exists {
					delete(provinceMap, city)
					if len(provinceMap) == 0 {
						delete(countryMap, province)
						if len(countryMap) == 0 {
							delete(newRegionMap, country)
						}
					}
				}
			}
		default:
			return nil, fmt.Errorf("invalid exclude format: %s", exclude)
		}
	}

	return &Distributor{RegionMap: newRegionMap,DistributorID: id}, nil
}

// CheckCityProvinceCountry checks whether a specific city-province-country combination exists in the distributor's region map.
func (d *Distributor) CheckCityProvinceCountry(query string) bool {
	parts := strings.Split(query, "-")
	if len(parts) != 3 {
		return false
	}

	city, province, country := parts[0], parts[1], parts[2]
	if provinces, exists := d.RegionMap[country]; exists {
		if cities, exists := provinces[province]; exists {
			return cities[city]
		}
	}

	return false
}