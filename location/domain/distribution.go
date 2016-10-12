package domain

import "errors"

var (
	ErrNotFound                = errors.New("Not Found")
	ErrExists                  = errors.New("Exists")
	ErrInvalidArgument         = errors.New("Invalid Argument")
	ErrParentHaveNotPermission = errors.New("Parent have not this permission")
	ErrAlreadyHavePermission   = errors.New("Already have permission")
	ErrInvalidLocation         = errors.New("Invalid Location")
	ErrDistributorNotFound     = errors.New("Distributor Not Found")
)

type DistributionRepository interface {
	DistributorExits(distributorId DistributorId) (ok bool, err error)
	GetCountryPermission(distributorId DistributorId, countryCode CountryCode) (countryPermission Permission, err error)
	GetStatePermission(distributorId DistributorId, countryCode CountryCode, stateCode StateCode) (statePermission Permission, err error)
	GetCityPermission(distributorId DistributorId, countryCode CountryCode, stateCode StateCode, cityCode CityCode) (cityPermission Permission, err error)
	ListCountryPermission(distributorId DistributorId) (countyPermissions []CountryPermission, err error)
	ListStatePermission(distributorId DistributorId, countyCode CountryCode) (statePermissions []StatePermission, err error)
	ListCityPermission(distributorId DistributorId, countyCode CountryCode, stateCode StateCode) (cityPermissions []CityPermission, err error)
	StoreCountry(distibutionId DistributorId, countryCode CountryCode, countryPermission Permission) (err error)
	StoreState(distibutionId DistributorId, countryCode CountryCode, stateCode StateCode, statePermission Permission) (err error)
	StoreCity(distibutionId DistributorId, countryCode CountryCode, stateCode StateCode, cityCode CityCode, cityPermission Permission) (err error)
}

type DistributorId string

//go:generate stringer -type=LocationType
//go:generate jsonenums -type=LocationType
type LocationType int8

const (
	Country LocationType = iota + 1
	State
	City
)

//go:generate stringer -type=Permission
//go:generate jsonenums -type=Permission
type Permission int8

const (
	Granted Permission = iota + 1
	Denied
	NotDefined
)

type DistributorPermission struct {
	LocationType LocationType
	Permission   Permission
	CountryCode  CountryCode
	StateCode    StateCode
	CityCode     CityCode
}

func NewDistributorPermission(permission Permission, countryCode CountryCode, stateCode StateCode, cityCode CityCode) *DistributorPermission {
	var locationType LocationType
	if len(cityCode) < 1 {
		if len(stateCode) < 1 {
			locationType = Country
		} else {
			locationType = State
		}
	} else {
		locationType = City
	}

	return &DistributorPermission{
		LocationType: locationType,
		Permission:   permission,
		CountryCode:  countryCode,
		StateCode:    stateCode,
		CityCode:     cityCode,
	}

}

type CityPermission struct {
	Permission Permission
	CityCode   CityCode
}

type StatePermission struct {
	Permission Permission
	StateCode  StateCode
}

type CountryPermission struct {
	Permission  Permission
	CountryCode CountryCode
}

func (d *DistributorPermission) Validate() error {
	switch d.LocationType {
	case City:
		if len(d.CityCode) < 1 || len(d.StateCode) < 1 || len(d.CountryCode) < 1 {
			return ErrInvalidArgument
		}
		fallthrough
	case State:
		if len(d.StateCode) < 1 || len(d.CountryCode) < 1 {
			return ErrInvalidArgument
		}
		fallthrough
	case Country:
		if len(d.CountryCode) < 1 {
			return ErrInvalidArgument
		}
		break
	default:
		return ErrInvalidArgument
	}

	if d.Permission != Granted && d.Permission != Denied {
		return ErrInvalidArgument
	}

	return nil

}