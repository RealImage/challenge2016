package handler

import (
	"distributor/service"
	"encoding/json"
	"github.com/go-chassis/openlog"
	"net/http"
)

func CreateDistributorHandler(w http.ResponseWriter, r *http.Request) {
	openlog.Info("Got a request to create distributor")
	w.Header().Set("Content-Type", "application/json")
	request, err := DecodeCreateDistributorRequest(r)
	if err != nil {
		openlog.Info("Getting Error while DecodingCreateDistributorHandler")
		w.WriteHeader(http.StatusBadRequest)
	}
	svc := service.NewDistributorService()
	response, err := svc.CreateDistributor(request)
	if err != nil {
		openlog.Info("Getting Error while Creating Distributor Entry")
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func GetDistributorLocationDetailsHandler(w http.ResponseWriter, r *http.Request) {
	openlog.Info("Got a request to check the location access of a distributor")
	w.Header().Set("Content-Type", "application/json")
	request, err := DecodeGetDistributorLocationDetails(r)
	if err != nil {
		openlog.Info("Getting Error while DecodeGetDistributorLocationDetails")
		w.WriteHeader(http.StatusBadRequest)
	}
	svc := service.NewDistributorService()
	response, err := svc.GetDistributorLocationDetails(request)
	if err != nil {
		openlog.Info("Getting Error while Getting Distributor Location Details")
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetDistributorDetailsHandler(w http.ResponseWriter, r *http.Request) {
	openlog.Info("Got a request to get the details of distributor")
	w.Header().Set("Content-Type", "application/json")
	request, err := DecodeGetDistributorDetails(r)
	if err != nil {
		openlog.Info("Getting Error while DecodeGetDistributorDetails")
		w.WriteHeader(http.StatusBadRequest)
	}
	svc := service.NewDistributorService()
	response, err := svc.GetDistributorDetails(request)
	if err != nil {
		openlog.Info("Getting Error while Getting DistributorDetails")
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
