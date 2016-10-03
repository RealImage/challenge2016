package repository

import (
	"log"
	"sync"

	"github.com/RealImage/challenge2016/location/domain"
)

type distributorRepository struct {
	mtx          sync.RWMutex
	distributors map[domain.DistributorId]*distributor
}

func NewDistributorRepository() *distributorRepository {
	return &distributorRepository{
		distributors: make(map[domain.DistributorId]*distributor),
	}
}

type distributor struct {
	id domain.DistributorId

	mtx       sync.RWMutex
	countries map[domain.CountryCode]*countryDis
}

func newDistributor(distibutionId domain.DistributorId) *distributor {
	return &distributor{
		id:        distibutionId,
		countries: make(map[domain.CountryCode]*countryDis),
	}
}

type countryDis struct {
	code       domain.CountryCode
	permission domain.Permission

	mtx    sync.RWMutex
	states map[domain.StateCode]*stateDis
}

func newCountryDis(code domain.CountryCode, permission domain.Permission) *countryDis {
	return &countryDis{
		code:       code,
		permission: permission,
		states:     make(map[domain.StateCode]*stateDis),
	}
}

type stateDis struct {
	code       domain.StateCode
	permission domain.Permission

	mtx    sync.RWMutex
	cities map[domain.CityCode]*cityDis
}

func newStateDis(code domain.StateCode, permission domain.Permission) *stateDis {
	return &stateDis{
		code:       code,
		permission: permission,
		cities:     make(map[domain.CityCode]*cityDis),
	}
}

type cityDis struct {
	code       domain.CityCode
	permission domain.Permission
}

func newCityDist(code domain.CityCode, permission domain.Permission) *cityDis {
	return &cityDis{
		code:       code,
		permission: permission,
	}
}

func (r *distributorRepository) FindCountryPermission(distributorId domain.DistributorId, countryCode domain.CountryCode) (countryPermission domain.Permission, err error) {
	country, err := r.findCountry(distributorId, countryCode)
	if err != nil {
		return
	}
	return country.permission, err
}

func (r *distributorRepository) FindStatePermission(distributorId domain.DistributorId, countryCode domain.CountryCode, stateCode domain.StateCode) (statePermission domain.Permission, err error) {
	state, err := r.findState(distributorId, countryCode, stateCode)
	if err != nil {
		return
	}

	return state.permission, nil
}

func (r *distributorRepository) FindCityPermission(distributorId domain.DistributorId, countryCode domain.CountryCode, stateCode domain.StateCode, cityCode domain.CityCode) (cityPermission domain.Permission, err error) {
	city, err := r.findCity(distributorId, countryCode, stateCode, cityCode)
	if err != nil {
		return
	}
	return city.permission, nil
}

//TODO FindAll Proper
func (r *distributorRepository) FindAll() {
	r.mtx.RLock()
	for _, d := range r.distributors {
		d.mtx.RLock()
		for _, c := range d.countries {
			c.mtx.RLock()
			for _, s := range c.states {
				s.mtx.RLock()
				for _, city := range s.cities {
					log.Println(d.id, "\t", "\t", c.code, ":", c.permission, "\t", s.code, ":", s.permission, "\t", city.code, ":", city.permission)
				}
				s.mtx.RUnlock()
			}
			c.mtx.RUnlock()
		}
		d.mtx.RUnlock()
	}
	r.mtx.RUnlock()
}

func (r *distributorRepository) StoreCountry(distibutionId domain.DistributorId, countryCode domain.CountryCode, countryPermission domain.Permission) (err error) {
	r.storeDistributor(distibutionId)
	r.storeCountry(distibutionId, countryCode, countryPermission)
	return
}

func (r *distributorRepository) StoreState(distibutionId domain.DistributorId, countryCode domain.CountryCode, stateCode domain.StateCode, statePermission domain.Permission) (err error) {
	r.storeDistributor(distibutionId)
	r.storeState(distibutionId, countryCode, stateCode, statePermission)
	return
}

func (r *distributorRepository) StoreCity(distibutionId domain.DistributorId, countryCode domain.CountryCode, stateCode domain.StateCode, cityCode domain.CityCode, cityPermission domain.Permission) (err error) {
	r.storeDistributor(distibutionId)
	r.storeCity(distibutionId, countryCode, stateCode, cityCode, cityPermission)
	return
}

//Distributor
func (r *distributorRepository) storeDistributor(distibutionId domain.DistributorId) (err error) {
	_, err = r.findDistributor(distibutionId)
	if err == nil {
		return domain.ErrExists
	}

	newD := newDistributor(distibutionId)
	r.storeDistributorHelper(newD)

	return nil
}
func (r *distributorRepository) storeDistributorHelper(d *distributor) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.distributors[d.id] = d
	return
}

func (r *distributorRepository) findDistributor(distibutionId domain.DistributorId) (d *distributor, err error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	d, ok := r.distributors[distibutionId]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return d, nil
}

//Country
func (r *distributorRepository) storeCountry(distibutionId domain.DistributorId, countryCode domain.CountryCode, permission domain.Permission) error {
	d, err := r.findDistributor(distibutionId)
	if err != nil {
		return err
	}

	c, err := findCountryHelper(d, countryCode)
	if err != nil {
		c = newCountryDis(countryCode, permission)
	} else {
		c.code = countryCode
		c.permission = permission
	}

	storeCountryHelper(d, c)

	return nil
}

func storeCountryHelper(d *distributor, c *countryDis) {
	d.mtx.Lock()
	defer d.mtx.Unlock()
	d.countries[c.code] = c
	return
}

func (r *distributorRepository) findCountry(distibutionId domain.DistributorId, code domain.CountryCode) (c *countryDis, err error) {
	d, err := r.findDistributor(distibutionId)
	if err != nil {
		return nil, err
	}

	c, err = findCountryHelper(d, code)
	if err != nil {
		return
	}

	return c, nil
}

func findCountryHelper(d *distributor, code domain.CountryCode) (c *countryDis, err error) {
	d.mtx.RLock()
	defer d.mtx.RUnlock()
	c, ok := d.countries[code]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return c, nil
}

//State
func (r *distributorRepository) storeState(distibutionId domain.DistributorId, countryCode domain.CountryCode, stateCode domain.StateCode, permission domain.Permission) error {
	c, err := r.findCountry(distibutionId, countryCode)
	if err != nil {
		return err
	}

	s, err := findStateDisHelper(c, stateCode)
	if err != nil {
		s = newStateDis(stateCode, permission)
	} else {
		s.code = stateCode
		s.permission = permission
	}

	storeStateDisHelper(c, s)

	return nil
}

func storeStateDisHelper(c *countryDis, s *stateDis) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.states[s.code] = s
	return
}

func (r *distributorRepository) findState(distibutionId domain.DistributorId, countryCode domain.CountryCode, stateCode domain.StateCode) (s *stateDis, err error) {
	c, err := r.findCountry(distibutionId, countryCode)
	if err != nil {
		return nil, err
	}

	s, err = findStateDisHelper(c, stateCode)
	if err != nil {
		return
	}

	return s, nil
}

func findStateDisHelper(c *countryDis, stateCode domain.StateCode) (s *stateDis, err error) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	s, ok := c.states[stateCode]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return s, nil
}

func (r *distributorRepository) storeCity(distibutionId domain.DistributorId, countryCode domain.CountryCode, stateCode domain.StateCode, cityCode domain.CityCode, permission domain.Permission) error {
	s, err := r.findState(distibutionId, countryCode, stateCode)
	if err != nil {
		return err
	}

	c, err := findCityDisHelper(s, cityCode)
	if err != nil {
		c = newCityDist(cityCode, permission)
	} else {
		c.code = cityCode
		c.permission = permission
	}

	storeCityDisHelper(s, c)
	return nil
}

func storeCityDisHelper(s *stateDis, cityDis *cityDis) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.cities[cityDis.code] = cityDis
	return
}

func (r *distributorRepository) findCity(distibutionId domain.DistributorId, countryCode domain.CountryCode, stateCode domain.StateCode, cityCode domain.CityCode) (cityDis *cityDis, err error) {
	s, err := r.findState(distibutionId, countryCode, stateCode)
	if err != nil {
		return nil, err
	}

	cityDis, err = findCityDisHelper(s, cityCode)
	if err != nil {
		return
	}

	return cityDis, nil
}

func findCityDisHelper(s *stateDis, cityCode domain.CityCode) (cityDis *cityDis, err error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	cityDis, ok := s.cities[cityCode]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return cityDis, nil
}
