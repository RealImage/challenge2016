package main

import (
	"log"
	"net/http"

	"./distributor"
)

func main() {

	http.HandleFunc("/createDistributor", distributor.CreateDistributor)
	http.HandleFunc("/verifyDistribution", distributor.VerifyDistributorRights)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
