package estimation

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"os"
)

type Data struct {
	CityCode     string
	ProvinceCode string
	CountryCode  string
	CityName     string
	ProvinceName string
	CountryName  string
}

func parseCSV(file multipart.File) ([][]string, error) {

	// read csv values using csv.Reader
	csvReader, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	return csvReader, nil
}

func SetDetails() (detail []Data) {
	file, err := os.Open("cities.csv")
	if err != nil {
		fmt.Errorf("error while opening the file")
	}
	defer file.Close()
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		fmt.Errorf("Error reading file")
	}
	var details []Data
	for i := 1; i < len(data); i++ {
		document := Data{
			CityCode:     data[i][0],
			ProvinceCode: data[i][1],
			CountryCode:  data[i][2],
			CityName:     data[i][3],
			ProvinceName: data[i][4],
			CountryName:  data[i][5],
		}
		details = append(details, document)

	}
	return details
}
