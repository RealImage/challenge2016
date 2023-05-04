package cmaps

import (
	"distribution-mgmnt/app"
	"strings"
	"sync"
)

var DistributorMgmntDB *DistributionMaps

type DistributionMaps struct {
	CityMap     map[string]string
	ProvinceMap map[string]string
	CountryMap  map[string]map[string][]string
	Distributor map[string]app.DistributorDetails
	mu          sync.Mutex
}

func NewDistributionMaps() *DistributionMaps {
	return &DistributionMaps{
		CityMap:     make(map[string]string),
		ProvinceMap: make(map[string]string),
		CountryMap:  make(map[string]map[string][]string),
		Distributor: make(map[string]app.DistributorDetails),
		mu:          sync.Mutex{},
	}
}

func (d *DistributionMaps) SetCityMap(city, province string) {
	d.mu.Lock()

	d.CityMap[strings.ToUpper(city)] = strings.ToUpper(province)
	d.mu.Unlock()
}
func (d *DistributionMaps) GetProvinceFromCityMap(city string) (string, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	res, ok := d.CityMap[strings.ToUpper(city)]
	return res, ok
}
func (d *DistributionMaps) SetProvinceMap(province, country string) {
	d.mu.Lock()
	d.ProvinceMap[strings.ToUpper(province)] = strings.ToUpper(country)
	d.mu.Unlock()
}
func (d *DistributionMaps) GetCountryFromProvinceMap(province string) (string, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	res, ok := d.ProvinceMap[strings.ToUpper(province)]
	return res, ok
}

func (d *DistributionMaps) SetCountryMap(city, province, country string) {
	d.mu.Lock()
	if d.CountryMap == nil {
		d.CountryMap = make(map[string]map[string][]string)
	}
	if d.CountryMap[strings.ToUpper(country)] == nil {
		d.CountryMap[strings.ToUpper(country)] = make(map[string][]string)
	}
	if d.CountryMap[strings.ToUpper(country)][strings.ToUpper(province)] == nil {
		d.CountryMap[strings.ToUpper(country)][strings.ToUpper(province)] = make([]string, 0)
	}
	d.CountryMap[strings.ToUpper(country)][strings.ToUpper(province)] = append(d.CountryMap[strings.ToUpper(country)][strings.ToUpper(province)], strings.ToUpper(city))
	d.mu.Unlock()
}

func (d *DistributionMaps) GetLocationFromCountryMap(country string) (map[string][]string, bool) {
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
