package resouces

import (
	"bufio"
	"encoding/csv"
	"os"
	"path/filepath"
	"sync"

	"chng2016/pkg/datasource"
	"chng2016/pkg/models"
)

// CSVLoader ...
type CSVLoader struct {
	dataSource datasource.Datasource

	sync.RWMutex
}

// NewLoader ...
func NewLoader(dataSource datasource.Datasource) *CSVLoader {
	return &CSVLoader{
		dataSource: dataSource,
	}
}

func (l *CSVLoader) LoadCSV() error {
	directory, err := os.Getwd() // get the current directory using the built-in function
	if err != nil {
		return err
	}

	// Construct the file path of the file in the same directory
	filePath := filepath.Join(directory, "resouces/cities.csv")
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	reader := csv.NewReader(bufio.NewReader(file))
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for _, rec := range records[1:] {
		d := &models.City{}
		d.CityCode = rec[0]
		d.ProvinceCode = rec[1]
		d.CountryCode = rec[2]
		d.CityName = rec[3]
		d.ProvinceName = rec[4]
		d.CountryName = rec[5]

		wg.Add(1)
		go func(city *models.City) {
			defer wg.Done()
			l.dataSource.SetCountryDetails(&models.Country{
				CountryCode: city.CountryCode,
				StateCode:   city.ProvinceCode,
				CityCode:    city.CityCode,
			})
		}(d)
	}
	wg.Wait()

	return nil
}
