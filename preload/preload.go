package preload

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"

	"github.com/challenge2016/models"
)

func Preload(dMap *models.DistributionMaps,fileName string) error{

	// reading the csv file
	reader,err := os.OpenFile(fileName,os.O_RDONLY,0777)
	if err != nil{
		log.Println("Error in opening file, Err",err)
		return err
	}

	csvReader := csv.NewReader(reader)

	for {
		record,err := csvReader.Read()
		if err == io.EOF{
			break
		}

		if err != nil{
			log.Printf("Err :%v",err)
			return err
		}

		if len(record) != 6{
			log.Println("Record length is less than 6")
		}

		loc := models.Location{
			CityCode: record[0],
			ProvinceCode: record[1],
			CountryCode: record[2],
			City: record[3],
			Province: record[4],
			Country: record[5],
		}

		setCityMap(dMap,loc)
		setProvinceMap(dMap,loc)
		setCountryMap(dMap,loc)
	}

	return nil
}

func setCityMap(dMap *models.DistributionMaps,loc models.Location){
	dMap.CityMap[strings.ToUpper(loc.City)] = &loc
}

func setProvinceMap(dMap *models.DistributionMaps,loc models.Location){
	loc.City = "ALL"
	loc.CityCode = "ALL"
	dMap.ProvinceMap[strings.ToUpper(loc.Province)] = &loc
}

func setCountryMap(dMap *models.DistributionMaps,loc models.Location){
	loc.City = "ALL"
	loc.CityCode = "ALL"
	loc.Province = "ALL"
	loc.ProvinceCode = "ALL"
	dMap.CountryMap[strings.ToUpper(loc.Country)] = &loc
}