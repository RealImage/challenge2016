package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bsyed6/challenge2016/model"
	"github.com/bsyed6/challenge2016/utils"
)

// CheckDistribution - Contains permission info
type CheckDistribution struct {
	Countries    map[string][]string
	Cities       map[string][]string
	Distributors map[string]model.Permission
}

func (d *CheckDistribution) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		var authorize model.IsAuthorized

		err := decoder.Decode(&authorize)
		if err != nil {
			log.Fatal(err)
		}
		if isError := utils.ValidateCheckPermission(w, authorize); isError {
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		if _, ok := d.Distributors[authorize.For]; ok {
			country := authorize.Country
			city := authorize.City
			province := authorize.Province
			distributor := d.Distributors[authorize.For]

			isCountryPresent := false
			if len(distributor.Includes) == 0 && len(distributor.SubIncludes) > 0 {

				for _, i := range distributor.SubIncludes {
					if i.Country == country {
						if i.City == city && i.Province == province {
							response := model.CheckDistributionResponse{IsAuthorized: "yes"}
							json.NewEncoder(w).Encode(&response)
							return
						}

					}
				}
				isProvinceExcluded := false
				for _, pro := range distributor.Excludes {
					if pro.Province == province && pro.City == "ALL" {
						isProvinceExcluded = true
					} else if pro.Province == province && pro.City == city {
						isProvinceExcluded = true
					}
				}
				if isProvinceExcluded {
					response := model.CheckDistributionResponse{IsAuthorized: "no"}
					json.NewEncoder(w).Encode(&response)
					return
				} else {
					response := model.CheckDistributionResponse{IsAuthorized: "no"}
					json.NewEncoder(w).Encode(&response)
					return
				}

			}

			for _, con := range distributor.Includes {
				if con == country {
					isCountryPresent = true
					isProvinceExcluded := false
					for _, pro := range distributor.Excludes {
						if pro.Province == province && pro.City == "ALL" {
							isProvinceExcluded = true
						} else if pro.Province == province && pro.City == city {
							isProvinceExcluded = true
						}
					}
					if isProvinceExcluded {
						response := model.CheckDistributionResponse{IsAuthorized: "no"}
						json.NewEncoder(w).Encode(&response)
						return
					} else {
						isCityPresent := false
						for _, c := range d.Cities[province] {
							if c == city {
								isCityPresent = true
							}
						}
						if isCityPresent {
							response := model.CheckDistributionResponse{IsAuthorized: "yes"}
							json.NewEncoder(w).Encode(&response)
							return
						}

					}

				}
			}
			if !isCountryPresent {
				response := model.CheckDistributionResponse{IsAuthorized: "no"}
				json.NewEncoder(w).Encode(&response)
				return
			} else {
				isValid := false
				isProvincePresent := false
				provinces := d.Countries[authorize.Country]
				for _, p := range provinces {
					if authorize.Province == p {
						isProvincePresent = true
						for _, c := range d.Cities[p] {
							if authorize.City == c {
								isValid = true
							}
						}

					}
				}
				if isValid || isProvincePresent && authorize.City == "" {
					response := model.CheckDistributionResponse{IsAuthorized: "yes"}
					json.NewEncoder(w).Encode(&response)
					return
				} else {
					response := model.CheckDistributionResponse{IsAuthorized: "no"}
					json.NewEncoder(w).Encode(&response)
					return
				}

			}

		} else {
			w.WriteHeader(http.StatusBadRequest)
			response := model.AssignResponse{Status: "Unknown Distributor!"}
			json.NewEncoder(w).Encode(&response)
			return
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		response := model.AssignResponse{Status: "Invalid Info!"}
		json.NewEncoder(w).Encode(&response)
	}

}
