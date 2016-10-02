package repository

import (
	"sync"

	"github.com/RealImage/challenge2016/location/domain"
)

type LocationRepository struct {
	mtx       sync.RWMutex
	countries map[domain.CountryCode]*country
}

func NewLocationRepository() *LocationRepository {
	return &LocationRepository{
		countries: make(map[domain.CountryCode]*country),
	}
}

type country struct {
	name string
	code domain.CountryCode

	mtx    sync.RWMutex
	states map[domain.StateCode]*state
}

func newCountry(name string, code domain.CountryCode) *country {
	return &country{
		name:   name,
		code:   code,
		states: make(map[domain.StateCode]*state),
	}
}

type state struct {
	name string
	code domain.StateCode

	mtx    sync.RWMutex
	cities map[domain.CityCode]*city
}

type city struct {
	name string
	code domain.CityCode
}

func newState(name string, code domain.StateCode) *state {
	return &state{
		name:   name,
		code:   code,
		cities: make(map[domain.CityCode]*city),
	}
}

func (r *LocationRepository) Store(c *domain.Location) (err error) {
	r.storeCountry(c.CountryName, c.CountryCode)
	r.storeState(c.CountryCode, c.StateName, c.StateCode)
	r.storeCity(c.CountryCode, c.StateCode, c.CityName, c.CityCode)
	return nil
}

func (r *LocationRepository) Find(countryCode domain.CountryCode, stateCode domain.StateCode, cityCode domain.CityCode) (c *domain.Location, err error) {
	country, err := r.findCountry(countryCode)
	if err != nil {
		return
	}

	state, err := findStateHelper(country, stateCode)
	if err != nil {
		return
	}

	city, err := findCityHelper(state, cityCode)
	if err != nil {
		return
	}

	return &domain.Location{
		CountryName: country.name,
		CountryCode: country.code,
		StateName:   state.name,
		StateCode:   state.code,
		CityName:    city.name,
		CityCode:    city.code,
	}, nil

}

func (r *LocationRepository) FindAll() (ls []*domain.Location, err error) {
	r.mtx.RLock()
	for _, country := range r.countries {
		country.mtx.RLock()
		for _, state := range country.states {
			state.mtx.RLock()
			for _, city := range state.cities {
				ls = append(ls, &domain.Location{
					CountryName: country.name,
					CountryCode: country.code,
					StateName:   state.name,
					StateCode:   state.code,
					CityName:    city.name,
					CityCode:    city.code,
				})
			}
			state.mtx.RUnlock()
		}
		country.mtx.RUnlock()
	}
	r.mtx.RUnlock()
	return
}

func (r *LocationRepository) storeCountry(countryName string, countryCode domain.CountryCode) error {
	_, err := r.findCountry(countryCode)
	if err == nil {
		return domain.ErrExists
	}

	newC := newCountry(countryName, countryCode)
	r.storeCountryHelper(newC)

	return nil
}

func (r *LocationRepository) storeCountryHelper(c *country) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.countries[c.code] = c
	return
}

func (r *LocationRepository) findCountry(code domain.CountryCode) (c *country, err error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	c, ok := r.countries[code]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return c, nil
}

func (r *LocationRepository) storeState(countryCode domain.CountryCode, stateName string, stateCode domain.StateCode) error {
	c, err := r.findCountry(countryCode)
	if err != nil {
		return err
	}

	_, err = findStateHelper(c, stateCode)
	if err == nil {
		return domain.ErrExists
	}

	newS := newState(stateName, stateCode)
	storeStateHelper(c, newS)

	return nil
}

func storeStateHelper(c *country, s *state) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.states[s.code] = s
	return
}

func (r *LocationRepository) findState(countryCode domain.CountryCode, stateCode domain.StateCode) (s *state, err error) {
	c, err := r.findCountry(countryCode)
	if err != nil {
		return nil, err
	}

	s, err = findStateHelper(c, stateCode)
	if err != nil {
		return
	}

	return s, nil
}

func findStateHelper(c *country, stateCode domain.StateCode) (s *state, err error) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	s, ok := c.states[stateCode]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return s, nil
}

func (r *LocationRepository) storeCity(countryCode domain.CountryCode, stateCode domain.StateCode, cityName string, cityCode domain.CityCode) error {
	s, err := r.findState(countryCode, stateCode)
	if err != nil {
		return err
	}

	_, err = findCityHelper(s, cityCode)
	if err == nil {
		return domain.ErrExists
	}

	newCity := &city{
		name: cityName,
		code: cityCode,
	}

	storeCityHelper(s, newCity)

	return nil
}

func storeCityHelper(s *state, city *city) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.cities[city.code] = city
	return
}

func (r *LocationRepository) findCity(countryCode domain.CountryCode, stateCode domain.StateCode, cityCode domain.CityCode) (city *city, err error) {
	s, err := r.findState(countryCode, stateCode)
	if err != nil {
		return nil, err
	}

	city, err = findCityHelper(s, cityCode)
	if err != nil {
		return
	}

	return city, nil
}

func findCityHelper(s *state, cityCode domain.CityCode) (city *city, err error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	city, ok := s.cities[cityCode]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return city, nil
}
