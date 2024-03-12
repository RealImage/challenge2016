package handler

import (
	"distributor/types"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func DecodeCreateDistributorRequest(r *http.Request) (types.DistributorRequest, error) {
	var payload types.DistributorRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return types.DistributorRequest{}, err
	}
	return payload, nil
}

func DecodeGetDistributorLocationDetails(r *http.Request) (types.LocationDetailsReq, error) {
	fmt.Println("Inside the DecodeGetDistributorLocationDetails")
	vars := mux.Vars(r)
	distributorName, ok := vars["distributor"]
	if !ok {
		fmt.Println("distributor name is missing in the params")
		return types.LocationDetailsReq{}, errors.New("distributor name is missing in the params")
	}
	queryParams := r.URL.Query()
	CountryCode := queryParams.Get("countryCode")
	ProvinceCode := queryParams.Get("provinceCode")
	CityCode := queryParams.Get("cityCode")

	data := types.LocationDetailsReq{
		DistributorName: distributorName,
		CountryCode:     CountryCode,
		ProvinceCode:    ProvinceCode,
		CityCode:        CityCode,
	}
	return data, nil
}

func DecodeGetDistributorDetails(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	distributor, ok := vars["distributor"]
	if !ok {
		fmt.Println("distributor name is missing in the params")
		return "", errors.New("distributor name is missing in the params")
	}
	return distributor, nil
}
