package files

import (
	"distribution-mgmnt/app"
	"distribution-mgmnt/pkg/cmaps"
	"distribution-mgmnt/pkg/util"
	"encoding/csv"
	"io"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
)

func CSVLoader(fileName string, headerNumber int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Errorln("error in os.Open()", err)
		return
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	for i := 0; i < headerNumber; i++ {
		header, _ := csvReader.Read()
		log.Infof("csv filename : %v header : %v", fileName, header)
	}
	i := int32(0)
	errRowcounter := 0
	locs := make([]app.Location, 0)
	for ; i < 2147483647; i++ {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(fileName, err, i, row)
			errRowcounter++
			continue
		}
		loc := util.ConvertStructToUpper(app.Location{
			City:       row[3],
			CityCD:     row[0],
			Country:    row[5],
			CountryCD:  row[2],
			Province:   row[4],
			ProvinceCD: row[1],
		})
		locs = append(locs, loc)
	}
	wg := sync.WaitGroup{}
	wg.Add(3)
	go cmaps.DistributorMgmntDB.SetCityMap(locs, &wg)
	go cmaps.DistributorMgmntDB.SetProvinceMap(locs, &wg)
	go cmaps.DistributorMgmntDB.SetCountryMap(locs, &wg)
	wg.Wait()
}
