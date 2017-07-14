package distributions /******** AUTHOR: NAGA SAI AAKARSHIT BATCHU ********/

import (
	disterror "../err"
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

// CSVData stuct for the incoming data from file
type CSVData struct {
	CityCode     string `csv:"City Code"`
	CityName     string `csv:"City Name"`
	ProvinceCode string `csv:"Province Code"`
	ProvinceName string `csv:"Province Name"`
	CountryCode  string `csv:"Country Code"`
	CountryName  string `csv:"Country Name"`
}

// LoadCSVData from the file
func LoadCSVData(filepath string) ([]*CSVData, error) {

	csvdata := []*CSVData{}
	csvFile, fileErr := os.Open(filepath)
	if fileErr != nil {
		ErrorLog("Failed to Open the File", fileErr)
		return nil, disterror.OsError(fileErr.Error())
	}
	records := csv.NewReader(bufio.NewReader(csvFile))
	_, firstrecordErr := records.Read()
	if firstrecordErr != nil {
		ErrorLog("Failed to Read First Record from CSV File", firstrecordErr)
		return nil, disterror.ReadError(firstrecordErr.Error())
	}

	for {
		record, csvReadErr := records.Read()
		if csvReadErr == io.EOF {
			//			InfoLog("Reached End of the File")
			break
		}
		if csvReadErr != nil {
			ErrorLog("Failed to Read Data from CSV File", csvReadErr)
			return nil, disterror.ReadError(csvReadErr.Error())
		}
		csvdata = append(csvdata, &CSVData{
			CityCode:     record[0],
			ProvinceCode: record[1],
			CountryCode:  record[2],
			CityName:     record[3],
			ProvinceName: record[4],
			CountryName:  record[5],
		})
	}
	return csvdata, nil
}
