package handlers

import (
	"encoding/json"
	"net/http"
	"task/common"
	"task/service"

	"github.com/go-chassis/openlog"
	"github.com/gorilla/mux"
)

type Handler struct {
	Service service.Service
}

type Response struct {
	Msg    string      `json:"_msg"`
	Status int         `json:"_status"`
	Data   interface{} `json:"data"`
}

func (h *Handler) CreateDistributor(w http.ResponseWriter, r *http.Request) {
	openlog.Info("Got a request to create distributor")
	w.Header().Set("Content-Type", "application/json")

	data := common.DistributorInput{}
	json.NewDecoder(r.Body).Decode(&data)
	// valres, err := validator.ValidatePaylaod("./../payloadschemas/transaction.json", data)
	// if err != nil {
	// 	openlog.Error(err.Error())
	// 	response := Response{Msg: err.Error(), Data: valres, Status: 400}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }
	res := h.Service.AddDistributorAndPermissions(data)
	w.WriteHeader(res.Status)
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) GetLocationAccess(w http.ResponseWriter, r *http.Request) {
	openlog.Info("Got a check the location access of a distributor")

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	distributor := vars["distributor"]

	queryParams := r.URL.Query()

	// Access specific query parameter
	CountryCode := queryParams.Get("countrycode")
	ProvinceCode := queryParams.Get("provincecode")
	CityCode := queryParams.Get("citycode")

	data := common.LocationAccessInput{
		Distributorname: distributor,
		CountryCode:     CountryCode,
		ProvinceCode:    ProvinceCode,
		CityCode:        CityCode,
	}
	res := h.Service.GetAccessDetailsOfDistributor(data)
	w.WriteHeader(res.Status)
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) GetDistributorDetails(w http.ResponseWriter, r *http.Request) {
	openlog.Info("Got a check the location access of a distributor")

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	distributor := vars["distributor"]

	res := h.Service.GetDistributorDetails(distributor)
	w.WriteHeader(res.Status)
	json.NewEncoder(w).Encode(res)
}
