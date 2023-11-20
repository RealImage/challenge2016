package handler

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nikhilsiwach28/Cinema-Distribution-System/models"
	"github.com/nikhilsiwach28/Cinema-Distribution-System/service"
)

type DistributorHandler struct {
	distributorSvc service.DistributorService
}

func NewDistributorHandler(DistributorService service.DistributorService) *DistributorHandler {
	return &DistributorHandler{distributorSvc: DistributorService}
}

func (h DistributorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var response apiResponse

	switch r.Method {
	case http.MethodPost:
		response = h.handleCreate(r)
	case http.MethodGet:
		response = h.handleGet(r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := response.Write(w); err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
	}
}

func (h *DistributorHandler) handleCreate(r *http.Request) apiResponse {
	var request models.CreateDistributorRequest
	if err := request.Parse(r); err != nil {
		log.Println(err)
		return models.NewAPIError(400, "ErrorBadData")
	}

	dist, err := h.distributorSvc.RegisterNewDistributor(request)
	if err != nil {
		log.Println(err)
		return models.NewAPIError(500, "Internal Server Error")
	}

	return models.NewCreateDistributorAPIResponse(&dist)
}

func (h *DistributorHandler) handleGet(r *http.Request) apiResponse {
	vars := mux.Vars(r)
	if id, err := uuid.Parse(vars["id"]); err != nil { // Extracting "id" from path parameters and parsing to uuid
		return models.NewAPIError(402, "Bad Request")
	} else {

		queryParams := r.URL.Query()
		region := queryParams.Get("region") // Extracting region from query param (City-State-Country , State-Country, Country)

		if ok := h.distributorSvc.CheckAuthorization(id, region); ok != true {
			return models.NewAPIError(404, "Not Found")
		} else {
			return models.NewAPIError(200, "Yes! Distributor is Authorised for the Region")
		}

	}
}

// Handler for split distribution

type splitDistributorHandler struct {
	distributorSvc service.SplitDistributorService
}

func NewSplitDistributorHandler(splitDistributorService service.SplitDistributorService) *splitDistributorHandler {
	return &splitDistributorHandler{distributorSvc: splitDistributorService}
}

func (h splitDistributorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var response apiResponse

	switch r.Method {
	case http.MethodPost:
		response = h.handleSplitCreate(r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := response.Write(w); err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
	}
}

func (h *splitDistributorHandler) handleSplitCreate(r *http.Request) apiResponse {
	var request models.CreateSplitDistributorRequest
	if err := request.Parse(r); err != nil {
		log.Println(err)
		return models.NewAPIError(400, "ErrorBadData")
	}

	dist, err := h.distributorSvc.SplitDistribution(request)
	if err != nil {
		log.Println(err)
		return models.NewAPIError(500, "Internal Server Error")
	}

	return models.NewCreateDistributorAPIResponse(&dist) // Reusing same Response for create and split for now
}
