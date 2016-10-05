package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"

	locationHttp "github.com/RealImage/challenge2016/location/client/locationService/http"
	"github.com/go-kit/kit/log"
	"github.com/opentracing/opentracing-go"
)

var inputCsv string
var locationServiceAddress string

func init() {
	flag.StringVar(&inputCsv, "i", "cities.csv", "Path to location csv")
	flag.StringVar(&locationServiceAddress, "host", "localhost:8080", "Location Service Address(Host:Port)")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	tracer := opentracing.GlobalTracer()
	locationService, err := locationHttp.New(locationServiceAddress, tracer, log.NewNopLogger())
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
		fmt.Printf("Input Csv Number of Column should be 6")
		return
	}

	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("Error While Read Csv Error : %s", err.Error())
				return
			}
		}

		err = locationService.AddLocation(ctx, record[5], record[2], record[4], record[1], record[3], record[1])
		if err != nil {
			fmt.Printf("Error While Add Location Error : ", err.Error())
			continue
		}
		//fmt.Printf("Location Country: %s \t State: %s \t City: %s \t Added Sucessfully \n", record[5], record[4], record[3])
	}

}
