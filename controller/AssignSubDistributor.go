package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bsyed6/challenge2016/model"
	"github.com/bsyed6/challenge2016/utils"
)

// AssignSubDistributor - contains sub distributor info
type AssignSubDistributor struct {
	Countries    map[string][]string
	Cities       map[string][]string
	Distributors map[string]model.Permission
}

func (d *AssignSubDistributor) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		var subPermission model.Permission
		err := decoder.Decode(&subPermission)
		if err != nil {
			log.Fatal(err)
		}
		if isError := utils.ValidateSubDistributor(w, subPermission); isError {
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		data := d.Distributors[subPermission.From]
		var unauthorized []model.Included
		var authorizedIncludes []model.Included
		var excludes []model.Excluded

		//Check if Distributor is present already!
		if distributor, ok := d.Distributors[subPermission.For]; ok {
			authorizedIncludes = distributor.SubIncludes
			excludes = distributor.Excludes

		}
		for _, include := range subPermission.SubIncludes {
			for _, country := range data.Includes {
				if include.Country == country {
					isEligible := true
					for _, excluded := range data.Excludes {
						if include.Province == excluded.Province && include.City == excluded.City {
							isEligible = false
							unauthorized = append(unauthorized, include)
						} else if include.Province == excluded.Province && excluded.City == "ALL" {
							isEligible = false
							unauthorized = append(unauthorized, include)
						}
					}
					if isEligible {
						isPresent := false
						if len(authorizedIncludes) > 0 {
							for _, i := range authorizedIncludes {
								if i == include {
									isPresent = true
								}
							}
							if !isPresent {
								authorizedIncludes = append(authorizedIncludes, include)
							}
						} else {
							authorizedIncludes = append(authorizedIncludes, include)

						}

					}

				}
			}
		}
		for _, exclude := range subPermission.Excludes {
			isPresent := false
			if len(excludes) > 0 {

				for _, i := range excludes {
					if i == exclude {
						isPresent = true
					}
				}
				if !isPresent {
					excludes = append(excludes, exclude)
				}
			} else {
				excludes = append(excludes, exclude)

			}

		}
		for _, e := range data.Excludes {
			isPresent := false
			if len(excludes) > 0 {
				for _, i := range excludes {
					if i == e {
						isPresent = true
					}
				}
				if !isPresent {
					excludes = append(excludes, e)
				}
			} else {
				excludes = append(excludes, e)

			}
		}
		d.Distributors[subPermission.For] = model.Permission{SubIncludes: authorizedIncludes, Excludes: excludes}

		response := model.AssignResponse{Status: "Distributor permissions successfully assigned!"}
		json.NewEncoder(w).Encode(&response)
	}
}
