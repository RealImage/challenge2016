package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bsyed6/challenge2016/model"
)

// CheckDistribution - Contains permission info
type CheckDistribution struct {
	Countries    map[string][]string
	Cities       map[string][]string
	Distributors map[string]model.Permission
}

func (d *CheckDistribution) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		var authorize model.IsAuthorized

		err := decoder.Decode(&authorize)
		if err != nil {
			log.Fatal(err)
		}
		if _, ok := d.Distributors[authorize.For]; ok {
			country := authorize.Country
			city := authorize.City
			province := authorize.Province
			distributor := d.Distributors[authorize.For]
			isCountryPresent := false
			if len(distributor.Includes) == 0 && len(distributor.SubIncludes) > 0 {
				for _, i := range distributor.SubIncludes {
					fmt.Println(i)
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
			}

		} else {
			w.WriteHeader(http.StatusBadRequest)
			response := model.AssignResponse{Status: "Unknown Distributor!"}
			json.NewEncoder(w).Encode(&response)
			return
		}

	} else {
		w.WriteHeader(http.StatusNoContent)
		response := model.AssignResponse{Status: "Invalid Info!"}
		json.NewEncoder(w).Encode(&response)
	}

}
