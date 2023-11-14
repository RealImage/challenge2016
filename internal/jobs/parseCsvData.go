package jobs

import (
	"challengeQube/dtos"
	"challengeQube/internal/globals"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

func ParseCsvData() error {
	file, err := os.Open(globals.CsvFileName)
	if err != nil {
		return errors.New("failed to open csv data file")
	}

	defer file.Close()
	//parse the file
	fileReader := csv.NewReader(file)

	finalMap := make(map[string]*dtos.Country, 0)
	for {
		//read the file
		record, err := fileReader.Read()
		//EOF check for end of csv file
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println("failed to read from csv", err)
			return err
		}
		// ctc is shorthand for city code, sc for state code and cc for country code
		ctc := strings.ToUpper(strings.ReplaceAll(record[0], " ", ""))
		sc := strings.ToUpper(strings.ReplaceAll(record[1], " ", ""))
		cc := strings.ToUpper(strings.ReplaceAll(record[2], " ", ""))

		// considering country code as mandatory
		if cc == "" {
			continue
		}

		// checking if the country code exists
		countryCode, exists := finalMap[cc]
		if !exists {
			countryCode = &dtos.Country{
				States: make(map[string]*dtos.State, 0),
			}
			finalMap[cc] = countryCode
		}

		// checking if the state code exists
		stateCode, exists := finalMap[cc].States[sc]
		if !exists {
			stateCode = &dtos.State{
				Cities: make(map[string]bool, 0),
			}
			countryCode.States[sc] = stateCode
		}

		// adding the city
		stateCode.Cities[ctc] = true
	}
	globals.MasterData = finalMap
	return nil
}
