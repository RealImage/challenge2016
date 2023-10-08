package main

import "fmt"

func (l *LocationData) ExistenceCheck(PermissionStruct map[string]struct{}, Val, City, Province, Country string) bool {

	// for checking the existence of region in PermissionStruct. If yes return error
	if _, ok := PermissionStruct[Val]; ok {
		return false
	}

	// for checking the existence of region in PermissionStruct. If yes return error
	for key := range PermissionStruct {
		if l.CheckCountryExistance(key, fmt.Sprintf("-%s", Val)) {
			return false
		}
		if City != "" && Province != "" && Country != "" {
			if l.CheckCountryExistance(key, fmt.Sprintf("%s-%s", Province, Country)) {
				return false
			}
			if _, ok := PermissionStruct[Country]; ok {
				return false
			}
		} else if Province != "" && Country != "" {
			if _, ok := PermissionStruct[Country]; ok {
				return false
			}
		}
	}
	return true
}
func (l *LocationData) CheckCountryExistance(s, substr string) bool {
	return (len(s) >= len(substr) && s[len(s)-len(substr):] == substr)
}

// for validating the provided regions
func (l *LocationData) ValidateLocation(Country string, Province string, City string) bool {

	if _, ok := l.AvailableLocations[Country]; !ok {
		return false
	}

	if Province != "" {
		if _, ok := l.AvailableLocations[Country][Province]; !ok {
			return false
		}
		if City != "" {
			if _, ok := l.AvailableLocations[Country][Province][City]; !ok {
				return false
			}
		}
	}

	return true
}

// for checking whether the distributor's regions are under parent distributor's regions
func (l *LocationData) ParentDistributorValidate(ParentDistributor string, Country string, Province string, City string) bool {
	if _, ok := l.Distributors[ParentDistributor]; !ok {
		return false
	}

	if _, ok := l.Distributors[ParentDistributor].Excluded[Country]; ok {
		return false
	}

	if Province != "" {
		if _, ok := l.Distributors[ParentDistributor].Excluded[fmt.Sprintf("%s-%s", Province, Country)]; ok {
			return false
		}
		if City != "" {
			if _, ok := l.Distributors[ParentDistributor].Excluded[fmt.Sprintf("%s-%s-%s", City, Province, Country)]; ok {
				return false
			}
		}
	}

	if City != "" && Province != "" && Country != "" {
		if _, ok := l.Distributors[ParentDistributor].Included[fmt.Sprintf("%s-%s-%s", City, Province, Country)]; ok {
			return true
		}
	}

	if Province != "" && Country != "" {
		if _, ok := l.Distributors[ParentDistributor].Included[fmt.Sprintf("%s-%s", Province, Country)]; ok {
			return true
		}
	}

	if Country != "" {
		if _, ok := l.Distributors[ParentDistributor].Included[fmt.Sprintf("%s", Country)]; ok {
			return true
		}
	}

	return false
}

// ensuring the excluded regions comes under included regions
func (l *LocationData) ValidateExcludedRegions(Included map[string]struct{}, Country string, Province string, City string) bool {
	if len(Included) <= 0 {
		return false
	}
	if City != "" && Province != "" && Country != "" {
		if _, ok := Included[fmt.Sprintf("%s-%s-%s", City, Province, Country)]; ok {
			return false
		}
		if _, ok := Included[fmt.Sprintf("%s-%s", Province, Country)]; ok {
			return true
		}
		if _, ok := Included[fmt.Sprintf("%s", Country)]; ok {
			return true
		}
		return false
	}

	if Province != "" && Country != "" {
		if _, ok := Included[fmt.Sprintf("%s-%s", Province, Country)]; ok {
			return false
		}
		if _, ok := Included[fmt.Sprintf("%s", Country)]; ok {
			return true
		}
		return false
	}

	if Country != "" {
		if _, ok := Included[fmt.Sprintf("%s", Country)]; ok {
			return false
		}
	}

	return false
}

// Checking permission available for the distributor or not
func (l *LocationData) PermissionCheckForDistributor(Distributor, Country, Province, City string) bool {
	Excluded := l.Distributors[Distributor].Excluded
	if _, ok := Excluded[fmt.Sprintf("%s-%s-%s", City, Province, Country)]; ok {
		return false
	}
	if _, ok := Excluded[fmt.Sprintf("%s-%s", Province, Country)]; ok {
		return false
	}
	if _, ok := Excluded[fmt.Sprintf("%s", Country)]; ok {
		return false
	}

	Parent, ok := l.DistributorParent[Distributor]
	if ok {
		return l.PermissionCheckForDistributor(Parent, Country, Province, City)
	}

	Included := l.Distributors[Distributor].Included
	if _, ok := Included[fmt.Sprintf("%s-%s-%s", City, Province, Country)]; ok {
		return true
	}
	if _, ok := Included[fmt.Sprintf("%s-%s", Province, Country)]; ok {
		return true
	}
	if _, ok := Included[fmt.Sprintf("%s", Country)]; ok {
		return true
	}
	return false
}
