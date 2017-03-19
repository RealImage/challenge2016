package distribution

import (
	"flag"
	"strings"
)

type PermissionMatrix map[string]map[string]map[string]bool

var basePermissions PermissionMatrix
var filePath string

func init() {
	flag.StringVar(&filePath, "file", "cities.csv", "Cities list file")
	cities, err := LoadCitiesFromCSV(filePath)
	if err != nil {
		panic(err)
	}
	basePermissions = make(PermissionMatrix)
	for _, city := range cities {
		_, isCountryThere := basePermissions[city.Country.Code]
		if !isCountryThere {
			basePermissions[city.Country.Code] = make(map[string]map[string]bool)
		}
		_, isProvinceThere := basePermissions[city.Country.Code][city.Provide.Code]
		if !isProvinceThere {
			basePermissions[city.Country.Code][city.Provide.Code] = make(map[string]bool)
		}
		basePermissions[city.Country.Code][city.Provide.Code][city.Code] = false
	}
}

// check whether the location is permitted
func (permissions *PermissionMatrix) IsAllowed(location string) bool {
	location = strings.TrimSpace(location)
	if len(location) == 0 {
		return false
	}
	parts := strings.Split(location, "-")
	if len(parts) > 3 {
		return false
	}

	var city, province, country string
	if len(parts) == 1 {
		country := parts[0]
		if len(permissions[country] == 0) {
			return false
		}
		for _, cities := range permissions[country] {
			for _, isAllowed := range cities {
				if !isAllowed {
					return false
				}
			}
		}
		return true
	}
	if len(parts) == 2 {
		country, province := parts[0], parts[1]
		if len(permissions[country]) == 0 || len(permissions[country][province]) == 0 {
			return false
		}
		for _, isAllowed := range permissions[country][province] {
			if !isAllowed {
				return false
			}
		}
		return true
	}
	if len(parts) == 3 {
		country, province, city := parts[0], parts[1], parts[2]
		if len(permissions[country]) == 0 || len(permissions[country][province]) == 0 || permissions[country][province][city] == nil {
			return false
		}
		return permissions[country][province][city]
	}
}

// update the location in the permissions as either true/false
func (permissions *PermissionMatrix) Update(location string, flag bool) ApplicationError {
	location = strings.TrimSpace(location)
	if len(location) == 0 {
		return nil //TODO: Raise error if necessary
	}
	parts := strings.Split(location, "-")
	if len(parts) > 3 {
		return nil //TODO: Raise error if necessary
	}

	var city, province, country string
	if len(parts) == 1 {
		country := parts[0]
		if len(permissions[country] == 0) {
			return nil //TODO: Raise error if necessary
		}
		for province, cities := range permissions[country] {
			for city, _ := range cities {
				permissions[country][province][city] = flag
			}
		}
	} else if len(parts) == 2 {
		country, province := parts[0], parts[1]
		if len(permissions[country]) == 0 || len(permissions[country][province]) == 0 {
			return nil //TODO: Raise error if necessary
		}
		for city, _ := range permissions[country][province] {
			permissions[country][province][city] = flag
		}
	} else if len(parts) == 3 {
		country, province, city := parts[0], parts[1], parts[2]
		permissions[country][province][city] = flag
	}
	return nil
}

// include the location in the permissions
func (permissions *PermissionMatrix) Include(location string) ApplicationError {
	return Update(location, true)
}

// exclude the location in the permissions
func (permissions *PermissionMatrix) Exclude(location string) ApplicationError {
	return Update(location, false)
}
