package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"

	locationHttp "github.com/RealImage/challenge2016/location/client/locationService/http"
	kitlog "github.com/go-kit/kit/log"
	"github.com/opentracing/opentracing-go"
)

var inputCsv string
var locationServiceAddress string
var logging bool

func init() {
	flag.StringVar(&inputCsv, "i", "cities.csv", "Path to location csv")
	flag.StringVar(&locationServiceAddress, "host", "localhost:8080", "Location Service Address(Host:Port)")
	flag.BoolVar(&logging, "l", false, "Logging Enable")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	if !logging {
		log.SetOutput(ioutil.Discard)
	}

	tracer := opentracing.GlobalTracer()
	locationService, err := locationHttp.New(locationServiceAddress, tracer, kitlog.NewNopLogger())
	if err != nil {
		fmt.Println(err.Error())
	}

	iFile, err := os.Open(inputCsv)
	if err != nil {
		fmt.Printf("Error While Open Input Csv. Error : %s \n", err.Error())
		return
	}

	r := csv.NewReader(iFile)
	r.Comment = '#'

	//Skipping Header
	r.Read()
	if r.FieldsPerRecord != 6 {
		fmt.Printf("Input Csv Number of Column should be 6 \n")
		return
	}

	httpChan := make(chan bool, 2)
	wg := sync.WaitGroup{}

	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("Error While Read Csv Error : %s\n", err.Error())
				return
			}
		}

		httpChan <- true
		wg.Add(1)

		go func(countryName string, countryCode string, stateName string, stateCode string, cityName string, cityCode string) {
			defer wg.Done()
			err = locationService.AddLocation(ctx, countryName, countryCode, stateName, stateCode, cityName, cityCode)
			if err != nil {
				fmt.Printf("Error While Add Location Error : %s \n", err.Error())
			} else {
				log.Printf("Location Country: %s \t State: %s \t City: %s \t Added Sucessfully \n", countryName, stateName, cityName)
			}
			<-httpChan

		}(record[5], record[2], record[4], record[1], record[3], record[0])

	}

	wg.Wait()

}
