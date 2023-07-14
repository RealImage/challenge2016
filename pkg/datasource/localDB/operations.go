package datasource

import (
	"chng2016/pkg/models"
)

func (l *LocalDBClient) IsCountryCodeValid(countryCode string) bool {
	l.Lock()
	defer l.Unlock()
	_, ok := l.countryCodes[countryCode]
	return ok
}

func (l *LocalDBClient) IsStateCodeValid(stateCode string) bool {
	l.Lock()
	defer l.Unlock()
	_, ok := l.stateCodes[stateCode]
	return ok
}

func (l *LocalDBClient) IsCityCodeValid(cityCode string) bool {
	l.Lock()
	defer l.Unlock()
	_, ok := l.cityCodes[cityCode]
	return ok
}

func (l *LocalDBClient) SetCountryDetails(country *models.Country) {
	l.Lock()
	l.countryCodes[country.CountryCode] = struct{}{}
	l.stateCodes[country.StateCode] = struct{}{}
	l.cityCodes[country.CityCode] = struct{}{}
	l.Unlock()
}
