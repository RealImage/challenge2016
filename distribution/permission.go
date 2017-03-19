package distribution

import (
	"strings"
)

type PermissionMatrix map[string]map[string]map[string]bool

// load the base permissions from the file
func (permissions *PermissionMatrix) Initialize(filePath string) ApplicationError {
	cities, err := LoadCitiesFromCSV(filePath)
	if err != nil {
		return err
	}
	for _, city := range cities {
		_, isCountryThere := permissions[city.Country.Code]
		if !isCountryThere {
			permissions[city.Country.Code] = make(map[string]map[string]bool)
		}
		_, isProvinceThere := permissions[city.Country.Code][city.Provide.Code]
		if !isProvinceThere {
			permissions[city.Country.Code][city.Provide.Code] = make(map[string]bool)
		}
		permissions[city.Country.Code][city.Provide.Code][city.Code] = false
	}
	return nil
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
	}
	if len(parts) == 2 {
		country, province := parts[0], parts[1]
		if len(permissions[country]) == 0 || len(permissions[country][province]) == 0 {
			return nil //TODO: Raise error if necessary
		}
		for city, _ := range permissions[country][province] {
			permissions[country][province][city] = flag
		}
	}
	if len(parts) == 3 {
		country, province, city := parts[0], parts[1], parts[2]
		permissions[country][province][city] = flag
	}
}

// include the location in the permissions
func (permissions *PermissionMatrix) Include(location string) ApplicationError {
	return Update(location, true)
}

// exclude the location in the permissions
func (permissions *PermissionMatrix) Exclude(location string) ApplicationError {
	return Update(location, false)
}
