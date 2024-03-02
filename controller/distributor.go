package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/saurabh-sde/challenge2016_saurabh/model"
	"github.com/saurabh-sde/challenge2016_saurabh/service"
	"github.com/saurabh-sde/challenge2016_saurabh/utils"
)

var DistributorMap map[string]model.Distributor

func GetDistributors(w http.ResponseWriter, r *http.Request) {
	list := service.GetDistributors()
	json.NewEncoder(w).Encode(list)
}

func AddDistributor(w http.ResponseWriter, r *http.Request) {
	var req utils.NewDistributorRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.Error(err)
		return
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		utils.Error(err)
		return
	}
	utils.Println("Creating Distributor: ", req)
	reslt, err := service.AddDistributor(&req)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(reslt)
}

func CheckDistributorPermission(w http.ResponseWriter, r *http.Request) {
	var req utils.CheckDistributorPermissionRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.Error(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		utils.Error(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	utils.Println("Checking Permission for Location: ", req)
	permission, err := service.GetDistributorPermission(req.Name, req.Locations)
	if err != nil {
		utils.Error(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	reslt := make(map[string]interface{})
	reslt["distributorName"] = req.Name
	reslt["permission"] = permission
	json.NewEncoder(w).Encode(reslt)
}
