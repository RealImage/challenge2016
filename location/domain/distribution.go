package domain

import "errors"

var (
	ErrNotFound                = errors.New("Not Found")
	ErrExists                  = errors.New("Exists")
	ErrInvalidArgument         = errors.New("Invalid Argument")
	ErrParentHaveNotPermission = errors.New("Parent have not this permission")
)

type DistributionRepository interface {
	FindCountryPermission(distributorId DistributorId, countryCode CountryCode) (countryPermission Permission, err error)
	FindStatePermission(distributorId DistributorId, countryCode CountryCode, stateCode StateCode) (statePermission Permission, err error)
	FindCityPermission(distributorId DistributorId, countryCode CountryCode, stateCode StateCode, cityCode CityCode) (cityPermission Permission, err error)
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

type Distributor struct {
	ParentId     DistributorId
	Id           DistributorId
	LocationType LocationType
	Permission   Permission
	CountryCode  CountryCode
	StateCode    StateCode
	CityCode     CityCode
}

func (d *Distributor) Validate() error {
	if len(d.Id) < 1 {
		return ErrInvalidArgument
	}

	switch d.LocationType {
	case City:
		if len(d.CityCode) < 1 {
			return ErrInvalidArgument
		}
		fallthrough
	case State:
		if len(d.StateCode) < 1 {
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
