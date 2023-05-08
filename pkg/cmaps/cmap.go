package cmaps

import (
	"distribution-mgmnt/app"
	"strings"
	"sync"
)

var DistributorMgmntDB *DistributionMaps

type DistributionMaps struct {
	CityMap     map[string]app.Location
	ProvinceMap map[string]app.Location
	CountryMap  map[string]map[string][]app.Location
	Distributor map[string]app.DistributorDetails
	mu          sync.Mutex
}

func NewDistributionMaps() *DistributionMaps {
	return &DistributionMaps{
		CityMap:     make(map[string]app.Location),
		ProvinceMap: make(map[string]app.Location),
		CountryMap:  make(map[string]map[string][]app.Location),
		Distributor: make(map[string]app.DistributorDetails),
		mu:          sync.Mutex{},
	}
}

func (d *DistributionMaps) SetCityMap(locs []app.Location, wg *sync.WaitGroup) {
	d.mu.Lock()
	for _, loc := range locs {
		d.CityMap[strings.ToUpper(loc.City)] = loc
	}
	d.mu.Unlock()
	wg.Done()
}
func (d *DistributionMaps) GetProvinceFromCityMap(city string) (app.Location, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	res, ok := d.CityMap[strings.ToUpper(city)]
	return res, ok
}
func (d *DistributionMaps) SetProvinceMap(locs []app.Location, wg *sync.WaitGroup) {
	d.mu.Lock()
	for _, loc := range locs {
		d.ProvinceMap[strings.ToUpper(loc.Province)] = loc
	}
	d.mu.Unlock()
	wg.Done()
}
func (d *DistributionMaps) GetCountryFromProvinceMap(province string) (app.Location, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	res, ok := d.ProvinceMap[strings.ToUpper(province)]
	return res, ok
}

func (d *DistributionMaps) SetCountryMap(locs []app.Location, wg *sync.WaitGroup) {
	d.mu.Lock()
	for _, loc := range locs {
		if d.CountryMap == nil {
			d.CountryMap = make(map[string]map[string][]app.Location)
		}
		if d.CountryMap[strings.ToUpper(loc.Country)] == nil {
			d.CountryMap[strings.ToUpper(loc.Country)] = make(map[string][]app.Location)
		}
		if d.CountryMap[strings.ToUpper(loc.Country)][strings.ToUpper(loc.Province)] == nil {
			d.CountryMap[strings.ToUpper(loc.Country)][strings.ToUpper(loc.Province)] = make([]app.Location, 0)
		}
		d.CountryMap[strings.ToUpper(loc.Country)][strings.ToUpper(loc.Province)] = append(d.CountryMap[strings.ToUpper(loc.Country)][strings.ToUpper(loc.Province)], loc)
	}
	d.mu.Unlock()
	wg.Done()
}

func (d *DistributionMaps) GetLocationFromCountryMap(country string) (map[string][]app.Location, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	res, ok := d.CountryMap[strings.ToUpper(country)]
	return res, ok
}

func (d *DistributionMaps) SetDistributionDB(details *app.DistributorDetails) {
	d.mu.Lock()
	d.Distributor[strings.ToUpper(details.Name)] = *details
	d.mu.Unlock()
}

func (d *DistributionMaps) GetDistributorDetails(distributorName string) (app.DistributorDetails, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	res, ok := d.Distributor[strings.ToUpper(distributorName)]
	return res, ok
}
