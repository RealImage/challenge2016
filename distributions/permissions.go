package distributions /******** AUTHOR: NAGA SAI AAKARSHIT BATCHU ********/

import (
	dicterr "../err"
)

// IsAllowed function will let us verify the Permissions
func (csvtemplate CountryMap) IsAllowed(locations []Location) bool {
	for _, locationselect := range locations {
		city, province, country := locationselect.City, locationselect.Province, locationselect.Country
		if city == "" || province == "" || country == "" {
			CustomErrorLog("Empty Fields Not Accepted")
			return false
		}
		if city == "*" && province == "*" && country == "*" {
			for countryselect := range csvtemplate {
				for _, cities := range csvtemplate[countryselect] {
					for _, isAllowed := range cities {
						if !isAllowed {
							return false
						}
					}
				}
			}
			continue
		} else if city == "*" && province == "*" && country != "*" {
			if _, countrycheck := csvtemplate[country]; countrycheck {
				for _, cities := range csvtemplate[country] {
					for _, isAllowed := range cities {
						if !isAllowed {
							return false
						}
					}
				}
			} else {
				CustomErrorLog("Not Allowed for : " + country)
				return false
			}
			continue
		} else if city == "*" && province != "*" && country != "*" {
			if _, provincecheck := csvtemplate[country][province]; provincecheck {
				for _, isAllowed := range csvtemplate[country][province] {
					if !isAllowed {
						return false
					}
				}
			} else {
				CustomErrorLog("Not Allowed for : " + country + "-" + province)
				return false
			}
			continue
		} else if city != "*" && province != "*" && country != "*" {
			if csvtemplate[country][province][city] {
				continue
			} else {
				CustomErrorLog("Not Allowed for : " + country + "-" + province + "-" + city)
				return false
			}
		}
	}
	return true
}

// Verify verifies the permission for one specific location
func (csvtemplate CountryMap) Verify(location Location) bool {
	city, province, country := location.City, location.Province, location.Country
	if city == "" || province == "" || country == "" {
		CustomErrorLog("Empty Fields Not Accepted")
		return false
	}
	if city == "*" && province == "*" && country == "*" {
		for countryselect := range csvtemplate {
			for _, cities := range csvtemplate[countryselect] {
				for _, isAllowed := range cities {
					if !isAllowed {
						return false
					}
				}
			}
		}
	} else if city == "*" && province == "*" && country != "*" {
		if _, countrycheck := csvtemplate[country]; countrycheck {
			for _, cities := range csvtemplate[country] {
				for _, isAllowed := range cities {
					if !isAllowed {
						return false
					}
				}
			}
		} else {
			CustomErrorLog("Not Allowed for : " + country)
			return false
		}
	} else if city == "*" && province != "*" && country != "*" {
		if _, provincecheck := csvtemplate[country][province]; provincecheck {
			for _, isAllowed := range csvtemplate[country][province] {
				if !isAllowed {
					return false
				}
			}
		} else {
			CustomErrorLog("Not Allowed for : " + country + "-" + province)
			return false
		}
	} else if city != "*" && province != "*" && country != "*" {
		if csvtemplate[country][province][city] {
		} else {
			CustomErrorLog("Not Allowed for : " + country + "-" + province + "-" + city)
			return false
		}
	}
	return true
}

//Update the location in csvtemplate as either true/false
func (csvtemplate CountryMap) Update(locations []Location, flag bool) error {
	for _, locationselect := range locations {
		city, province, country := locationselect.City, locationselect.Province, locationselect.Country
		if city == "*" && province == "*" && country == "*" {
			for countryselect := range csvtemplate {
				for provinceselect, cities := range csvtemplate[country] {
					for cityselect := range cities {
						csvtemplate[countryselect][provinceselect][cityselect] = flag
					}
				}
			}
		} else if city == "*" && province == "*" && country != "*" {
			if _, ok := csvtemplate[country]; ok {
				for provinceselect, cities := range csvtemplate[country] {
					for cityselect := range cities {
						csvtemplate[country][provinceselect][cityselect] = flag
					}
				}
			} else {
				return dicterr.UpdateError("Cannot Update the location: " + country)
			}
		} else if city == "*" && province != "*" && country != "*" {
			if _, ok := csvtemplate[country][province]; ok {
				for cityselect := range csvtemplate[country][province] {
					csvtemplate[country][province][cityselect] = flag
				}
			} else {
				return dicterr.UpdateError("Cannot Update the location: " + province + "-" + country)
			}
		} else if city != "*" && province != "*" && country != "*" {
			if _, ok := csvtemplate[country][province][city]; ok {
				csvtemplate[country][province][city] = flag
			} else {
				return dicterr.UpdateError("Cannot Update the location: " + city + "-" + province + "-" + country)
			}
		}
	}
	return nil
}

//Copy the Permissions for Include or Exclude from parent distributor Template
func (csvtemplate CountryMap) Copy(locations []Location, parent CountryMap) error {
	for _, locationselect := range locations {
		city, province, country := locationselect.City, locationselect.Province, locationselect.Country
		if city == "*" && province == "*" && country == "*" {
			for countryselect := range parent {
				for provinceselect, cities := range parent[countryselect] {
					for cityselect, flag := range cities {
						csvtemplate[countryselect][provinceselect][cityselect] = flag
					}
				}
			}
		} else if city == "*" && province == "*" && country != "*" {
			if _, countrycheck := parent[country]; countrycheck {
				for provinceselect, cities := range parent[country] {
					for cityselect, flag := range cities {
						csvtemplate[country][provinceselect][cityselect] = flag
					}
				}
			} else {
				return dicterr.CopyError("Cannot Copy the location from parent: " + country)
			}
		} else if city == "*" && province != "*" && country != "*" {
			if _, provincecheck := parent[country][province]; provincecheck {
				for cityselect, flag := range parent[country][province] {
					csvtemplate[country][province][cityselect] = flag
				}
			} else {
				return dicterr.CopyError("Cannot Copy the location from parent : " + country + "-" + province)
			}
		} else if city != "*" && province != "*" && country != "*" {
			if _, citycheck := parent[country][province][city]; citycheck {
				csvtemplate[country][province][city] = parent[country][province][city]
			} else {
				return dicterr.CopyError("Cannot Copy the location from parent : " + country + "-" + province + "-" + city)
			}
		}
	}
	return nil
}

// Include Locations in csvtemplate of the distributor
func (csvtemplate CountryMap) Include(location []Location) error {

	return csvtemplate.Update(location, true)
}

// Exclude Locations in csvtemplate of the distributor
func (csvtemplate CountryMap) Exclude(location []Location) error {

	return csvtemplate.Update(location, false)
}
