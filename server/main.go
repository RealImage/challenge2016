package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"task/common"
	"task/handlers"

	"github.com/go-chassis/openlog"
	"github.com/gorilla/mux"

	"task/service"
)

func GetService() service.Service {
	return service.Service{}
}
func LoadLocationsFromCSV(csvFile string) error {
	file, err := os.Open(csvFile)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %w", err)
	}

	for _, record := range records[1:] {
		loc := common.NewLocationIdentifier(record[2], record[1], record[0], record[5], record[4], record[3])
		common.CountryMap[loc.CountryCode] = loc
		common.ProvinceMap[loc.CountryCode+"$"+loc.ProvinceCode] = loc
		common.CityMap[loc.CountryCode+"$"+loc.ProvinceCode+"$"+loc.CityCode] = loc
	}

	return nil
}
func main() {
	err := LoadLocationsFromCSV("./../cities.csv")
	if err != nil {
		fmt.Println("Error loading location data:", err)
		return
	}

	r := mux.NewRouter()

	service := GetService()
	h := handlers.Handler{Service: service}

	r.HandleFunc("/distributors", h.CreateDistributor).Methods("POST")
	r.HandleFunc("/distributors/{distributor}/locationaccess", h.GetLocationAccess).Methods("GET")
	r.HandleFunc("/distributors/{distributor}", h.GetDistributorDetails).Methods("GET")
	openlog.Info("Started listening at http://localhost:8070")
	log.Fatal(http.ListenAndServe(":8070", r))

}
