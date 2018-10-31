package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bsyed6/challenge2016/model"
)

// AssignDistributor - contains distributor info
type AssignDistributor struct {
	Countries    map[string][]string
	Cities       map[string][]string
	Distributors map[string]model.Permission
}

func (d *AssignDistributor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		var permission model.Permission
		err := decoder.Decode(&permission)
		if err != nil {
			log.Fatal(err)
		}

		//Validating if the country is present in our countries datastore
		for _, country := range permission.Includes {
			if provinces, ok := d.Countries[country]; ok {
				// Checking if the Distributor is already present in the distributor datastore
				if distributor, ok := d.Distributors[permission.For]; ok {

					for _, v := range permission.Includes {
						isPresent := false
						for _, element := range distributor.Includes {
							if v == element {
								isPresent = true
							}
						}
						if !isPresent {
							distributor.Includes = append(distributor.Includes, v)
						}
					}

					for _, e := range permission.Excludes {

						// check e.Country is present in existing distributor.Includes
						isCountryPresent := false
						for _, existingCountry := range distributor.Includes {
							if existingCountry == e.Country {
								isCountryPresent = true
							}
						}
						// check e.Country is present in existing permission.Includes

						for _, newCountry := range permission.Includes {
							if newCountry == e.Country {
								isCountryPresent = true
							}
						}

						if isCountryPresent {
							// checking v.Province is present in  Provinces(countries datastore)
							for _, p := range provinces {
								if e.Province == p {
									isPresent := false
									for _, element := range distributor.Excludes {
										if e.Province == element.Province && element.City == "ALL" {
											isPresent = true
										} else if e.Province == element.Province && element.City == e.City {
											isPresent = true
										}
									}
									if !isPresent {
										if e.City == "" {
											e.City = "ALL"
											distributor.Excludes = append(distributor.Excludes, e)
										} else {

											// checking if the city is present in that province
											for _, c := range d.Cities[p] {
												if e.City == c {
													distributor.Excludes = append(distributor.Excludes, e)

												}
											}
										}

									}

								}
							}
						}
					}
					d.Distributors[permission.For] = distributor
					fmt.Println(d.Distributors)

				} else {
					var newPermission model.Permission
					for _, country := range permission.Includes {
						if provinces, ok := d.Countries[country]; ok {
							// distributor.Includes = permission.Includes
							newPermission.Includes = append(newPermission.Includes, country)
							for _, e := range permission.Excludes {
								isCountryPresent := false
								for _, newCountry := range permission.Includes {
									if newCountry == e.Country {
										isCountryPresent = true
									}
								}
								if isCountryPresent {

									// checking v.Province is present in  Provinces(countries datastore)
									for _, p := range provinces {
										if e.Province == p {
											for _, c := range d.Cities[p] {

												if e.City == c {
													newPermission.Excludes = append(newPermission.Excludes, e)
												}
											}
											if e.City == "" {
												e.City = "ALL"
												newPermission.Excludes = append(newPermission.Excludes, e)

											}
										}
									}
								}
							}
							d.Distributors[permission.For] = newPermission

						}

					}

					fmt.Println(d.Distributors)
				}

			} else {
				return
			}
		}

	}
	response := model.AssignResponse{Status: "Distributor permissions successfully assigned!"}
	json.NewEncoder(w).Encode(&response)

}
