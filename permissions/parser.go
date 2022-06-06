package permissions

import (
	"encoding/csv"
	"os"
)

const (
	CitiCodeColumn = iota
	ProvinceCodeColumn
	CountryCodeColumn
	CityNameColumn
	ProvinceNameColumn
	CountryNameColumn
)

type (
	info struct {
		Code, Name string
	}

	RegionInfo struct {
		City, Province, Country info
	}
)

func parseCsvToRegionsInfos(file string) ([]*RegionInfo, error) {
	data, err := readCsv(file)
	if err != nil {
		return nil, err
	}

	regions := make([]*RegionInfo, len(data))
	// omit header
	for i := 1; i < len(data); i++ {
		row := data[i]
		ri := &RegionInfo{
			City: info{
				row[CitiCodeColumn],
				row[CityNameColumn],
			},
			Province: info{
				row[ProvinceCodeColumn],
				row[ProvinceNameColumn],
			},
			Country: info{
				row[CountryCodeColumn],
				row[CountryNameColumn],
			},
		}
		regions[i] = ri
	}

	return regions, nil
}

func readCsv(file string) ([][]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	return csvReader.ReadAll()
}

func (ri *RegionInfo) String() string {
	str := ""
	if ri.City.Name != "" {
		str += ri.City.Name
	}
	if ri.Province.Name != "" {
		str += separator + ri.City.Name
	}
	if ri.Country.Name != "" {
		str += separator + ri.City.Name
	}

	return str
}
