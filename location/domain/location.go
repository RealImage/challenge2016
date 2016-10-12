package domain

type LocationRepository interface {
	Store(c *Location) (err error)
	Find(countryCode CountryCode, stateCode StateCode, cityCode CityCode) (c *Location, err error)
	FindAll() (ls []*Location, err error)
	CountryExists(countryCode CountryCode) (ok bool, err error)
	StateExists(countryCode CountryCode, stateCode StateCode) (stateOk bool, err error)
	CityExists(countryCode CountryCode, stateCode StateCode, cityCode CityCode) (cityOk bool, err error)
}

type CityCode string
type StateCode string
type CountryCode string

type Location struct {
	CountryName string
	CountryCode CountryCode

	StateName string
	StateCode StateCode

	CityName string
	CityCode CityCode
}

func (l *Location) Validate() error {
	if len(l.CountryCode) < 1 || len(l.CountryName) < 1 {
		return ErrInvalidArgument
	}

	if len(l.StateName) < 1 || len(l.StateCode) < 1 {
		return ErrInvalidArgument
	}

	if len(l.CityName) < 1 || len(l.CityCode) < 1 {
		return ErrInvalidArgument
	}

	return nil
}
