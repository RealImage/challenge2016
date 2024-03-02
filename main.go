package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saurabh-sde/challenge2016_saurabh/controller"
	"github.com/saurabh-sde/challenge2016_saurabh/utils"
)

func init() {
	filePath := "cities.csv"
	_, err := utils.LoadCities(filePath)
	if err != nil {
		log.Fatal(err)
	}
	utils.InitDistributors()
}

func main() {
	req := mux.NewRouter()
	req.HandleFunc("/distributors", controller.GetDistributors).Methods("GET")
	req.HandleFunc("/distributor/add", controller.AddDistributor).Methods("POST")
	req.HandleFunc("/distributor/check", controller.CheckDistributorPermission).Methods("GET")
	utils.Println("Running Local Server : ", "http://localhost:8080")
	utils.Println(http.ListenAndServe(":8080", req))
}
