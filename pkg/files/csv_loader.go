package files

import (
	"distribution-mgmnt/pkg/cmaps"
	"distribution-mgmnt/pkg/util"
	"encoding/csv"
	"io"
	"os"

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
	for ; i < 2147483647; i++ {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(fileName, err, i, row)
			errRowcounter++
			continue
		}
		cmaps.DistributorMgmntDB.SetCityMap(util.RemoveSpacesAndToUpper(row[3]), util.RemoveSpacesAndToUpper(row[4]))
		cmaps.DistributorMgmntDB.SetProvinceMap(util.RemoveSpacesAndToUpper(row[4]), util.RemoveSpacesAndToUpper(row[5]))
		cmaps.DistributorMgmntDB.SetCountryMap(util.RemoveSpacesAndToUpper(row[3]), util.RemoveSpacesAndToUpper(row[4]), util.RemoveSpacesAndToUpper(row[5]))
	}
}
