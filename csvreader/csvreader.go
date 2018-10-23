package csvreader

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"

	"github.com/RealImageChallenge/models"
)

func MakeDataStore(csvFileName string, countryStateMap models.CountryMap) {
	csvFile, _ := os.Open(csvFileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("error reading all lines: %v", err)
	}

	for i, line := range lines {
		if i == 0 {
			// skip header line
			continue
		}
		sMap, cok := countryStateMap[line[5]]
		if cok {
			ctMap, stok := sMap[line[4]]
			if stok {
				ctMap[line[3]] = models.City{
					CityCode:     line[0],
					ProvinceCode: line[1],
					CountryCode:  line[2],
					CityName:     line[3],
					ProvinceName: line[4],
					CountryName:  line[5],
				}
			} else {
				cityMa := map[string]models.City{line[3]: models.City{
					CityCode:     line[0],
					ProvinceCode: line[1],
					CountryCode:  line[2],
					CityName:     line[3],
					ProvinceName: line[4],
					CountryName:  line[5],
				},
				}
				sMap[line[4]] = cityMa

			}

		} else {
			cityMa := map[string]models.City{line[3]: models.City{
				CityCode:     line[0],
				ProvinceCode: line[1],
				CountryCode:  line[2],
				CityName:     line[3],
				ProvinceName: line[4],
				CountryName:  line[5],
			},
			}
			stateM := map[string]models.CityMap{line[4]: cityMa}
			countryStateMap[line[5]] = stateM
		}
	}
}
