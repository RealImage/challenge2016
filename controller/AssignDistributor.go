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

			for _, v := range permission.Excludes {
				isPresent := false
				for _, element := range distributor.Excludes {
					if v.Province == element.Province && element.City == "ALL" {
						isPresent = true
					} else if v.Province == element.Province && element.City == v.City {
						isPresent = true
					}
				}
				if !isPresent {
					if v.City == "" {
						v.City = "ALL"
					}
					distributor.Excludes = append(distributor.Excludes, v)

				}
			}

			fmt.Println(distributor)
		} else {
			Exclude := []model.Excluded{}
			for _, v := range permission.Excludes {
				if v.City == "" {
					v.City = "ALL"
				}
				Exclude = append(Exclude, v)
			}
			permission.Excludes = Exclude
			d.Distributors[permission.For] = permission
			fmt.Println(d.Distributors)
		}

	}
	response := model.AssignResponse{Status: "Distributor permissions successfully assigned!"}
	json.NewEncoder(w).Encode(&response)

}
