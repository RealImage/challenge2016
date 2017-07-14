package distributions /******** AUTHOR: NAGA SAI AAKARSHIT BATCHU ********/

import (
	"flag"
)

// CityMap is a map to store cities and its permissions
type CityMap map[string]bool

//ProvinceMap is a map to store provinces and all its citymap's
type ProvinceMap map[string]CityMap

//CountryMap is a map to store countries and all its provincemap's
type CountryMap map[string]ProvinceMap

var filepath string

func init() {
	flag.StringVar(&filepath, "file", "/root/gowork/src/small-works/mywork/cities.csv", "CSV File Path")
}

//CreateCSVTemplate Takes the Loaded data and Creates a Template which can be used to verify Distributions
func CreateCSVTemplate() CountryMap {
	csvdata, loaderr := LoadCSVData(filepath)
	if loaderr != nil {
		ErrorLog("Failed to Load Data from CSV File", loaderr)
		panic(loaderr)
	}
	csvtemplate := make(CountryMap)
	for _, data := range csvdata {
		_, isCountryListed := csvtemplate[data.CountryCode]
		if !isCountryListed {
			csvtemplate[data.CountryCode] = make(ProvinceMap)
		}
		_, isProvinceListed := csvtemplate[data.CountryCode][data.ProvinceCode]
		if !isProvinceListed {
			csvtemplate[data.CountryCode][data.ProvinceCode] = make(CityMap)
		}
		csvtemplate[data.CountryCode][data.ProvinceCode][data.CityCode] = false
	}
	InfoLog("Loaded Template")
	return csvtemplate
}
