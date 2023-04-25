package handler

import (
	"math"

	"github.com/ishhyoboytarun/challenge2016/csv"
	"github.com/ishhyoboytarun/challenge2016/distributor"
)

// Distributors Define a struct for the distributor class
type Distributors struct {
	CountryLevelDistributorsMap  map[*distributor.Distributor]bool
	ProvinceLevelDistributorsMap map[*distributor.Distributor]bool
	CityLevelDistributorsMap     map[*distributor.Distributor]bool
	CountryWiseDistributorsCount int
}

func FetchCountriesToRegionMap(regions map[string]*csv.Location) map[string][]*csv.Location {

	countryToRegionMap := make(map[string][]*csv.Location)
	for _, country := range regions {
		countryToRegionMap[country.CountryName] = append(countryToRegionMap[country.CountryName], country)
	}
	return countryToRegionMap
}

func (d Distributors) AllotRegionsToDistributorsCountryWise(countriesToRegionsMap map[string][]*csv.Location) {

	for _, regions := range countriesToRegionsMap {
		d.AllotRegionsToCountryLevelDistributors(regions)
	}
}

func (d Distributors) AllotRegionsToCountryLevelDistributors(regions []*csv.Location) {

	var included []string
	currIndex := 0
	length := len(regions)
	for i := 0; i < d.CountryWiseDistributorsCount-1; i++ {
		nextIndex := currIndex + length/d.CountryWiseDistributorsCount
		included = getIncludedRegionsForCountryLevelDistributors(regions, currIndex, nextIndex)
		newDistributor := distributor.NewDistributor(included, []string{}, []*distributor.Distributor{}, 4)
		d.CountryLevelDistributorsMap[newDistributor] = true
		currIndex = nextIndex
	}

	included = getIncludedRegionsForCountryLevelDistributors(regions, currIndex, length)
	newDistributor := distributor.NewDistributor(included, []string{}, []*distributor.Distributor{}, 4)
	d.CountryLevelDistributorsMap[newDistributor] = true
}

func getIncludedRegionsForCountryLevelDistributors(regions []*csv.Location, start, end int) []string {

	var includedRegions []string
	for i := start; i < end; i++ {
		region := regions[i].CityName + regions[i].ProvinceName + regions[i].CountryName
		includedRegions = append(includedRegions, region)
	}
	return includedRegions
}

func (d Distributors) AddProvinceLevelDistributors() {

	for Distributor := range d.CountryLevelDistributorsMap {
		subDistributorsCount := Distributor.GetSubDistributorsCount()
		includedRegions := d.AllotRegionsToProvinceLevelDistributors(Distributor.GetIncludedRegion(),
			subDistributorsCount)
		for i := 0; i < len(includedRegions); i++ {
			subDistributor := Distributor.AddSubDistributor(includedRegions[i], []string{},
				[]*distributor.Distributor{}, subDistributorsCount)
			d.ProvinceLevelDistributorsMap[subDistributor] = true
		}
	}
}

func (d Distributors) AllotRegionsToProvinceLevelDistributors(regionsMap map[string]bool,
	subDistributorsCount int) [][]string {

	if len(regionsMap) == 0 {
		return [][]string{}
	}

	if len(regionsMap) < subDistributorsCount {
		subDistributorsCount = len(regionsMap)
	}

	included := make([][]string, subDistributorsCount)
	for i := 0; i < subDistributorsCount; i++ {
		included[i] = []string{}
	}

	regions := []string{}
	for city := range regionsMap {
		regions = append(regions, city)
	}

	batchSize := int(math.Ceil(float64(len(regions) / subDistributorsCount)))

	for i := range regions {
		index := i / batchSize
		if index == len(included) {
			included = append(included, []string{})
		}
		included[index] = append(included[index], regions[i])
	}

	return included
}
