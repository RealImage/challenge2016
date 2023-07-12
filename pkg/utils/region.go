package utils

import (
	"strings"

	"chng2016/pkg/models"
)

/*
	GetRegionDetails get region details like country,city and state code from the input region

@Parameter : Input parameter should be of CountryCode,StateCode,CityCode format
*/
func (a *AppUtil) GetRegionDetails(region string) (*models.Country, error) {
	if len(strings.TrimSpace(region)) == 0 {
		return nil, ErrInvalidRegion
	}

	segments := strings.Split(region, ",")
	res := &models.Country{}
	switch len(segments) {
	case 1:
		if !a.IsValidCountryCode(segments[0]) {
			return nil, ErrInValidCountryCode
		}
		res.CountryCode = segments[0]
	case 2:
		if !a.IsValidCountryCode(segments[0]) {
			return nil, ErrInValidCountryCode
		}
		if !a.IsValidStateCode(segments[1]) {
			return nil, ErrInValidStateCode
		}
		res.CountryCode = segments[0]
		res.StateCode = segments[1]
	case 3:
		if !a.IsValidCountryCode(segments[0]) {
			return nil, ErrInValidCountryCode
		}
		if !a.IsValidStateCode(segments[1]) {
			return nil, ErrInValidStateCode
		}
		if !a.IsValidCityCode(segments[2]) {
			return nil, ErrInValidCityCode
		}
		res.CountryCode = segments[0]
		res.StateCode = segments[1]
		res.CityCode = segments[2]
	default:
		return nil, ErrInvalidRegion
	}

	return res, nil
}

// IsValidStateCode ...
func (a *AppUtil) IsValidStateCode(value string) bool {
	return a.localDB.IsStateCodeValid(value)
}

// IsValidCountryCode ...
func (a *AppUtil) IsValidCountryCode(value string) bool {
	return a.localDB.IsCountryCodeValid(value)
}

// IsValidCityCode ...
func (a *AppUtil) IsValidCityCode(value string) bool {
	return a.localDB.IsCityCodeValid(value)
}
