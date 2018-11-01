package utils

import (
	"encoding/json"
	"net/http"

	"github.com/bsyed6/challenge2016/model"
)

// Validate - contains validation for permission request
func Validate(w http.ResponseWriter, p model.Permission) bool {
	isError := false
	if len(p.Includes) == 0 {
		isError = true
		sendError(w, "Includes is not present!")
	} else if p.For == "" {
		isError = true
		sendError(w, "Distributor is not mentioned!")
	} else if len(p.Excludes) == 0 {
		isError = true
		sendError(w, "Excludes is not present!")
	}
	return isError

}

// ValidateCheckPermission - to check the permissions for the particular region
func ValidateCheckPermission(w http.ResponseWriter, p model.IsAuthorized) bool {
	isError := false
	if p.City == "" {
		isError = true
		sendError(w, "City is not present!")
	} else if p.For == "" {
		isError = true
		sendError(w, "Distributor is not mentioned!")
	} else if p.Country == "" {
		isError = true
		sendError(w, "Country is not present!")
	} else if p.Province == "" {
		isError = true
		sendError(w, "Province is not present!")
	}
	return isError

}

// ValidateSubDistributor - contains sub distributor info
func ValidateSubDistributor(w http.ResponseWriter, p model.Permission) bool {
	isError := false
	if len(p.SubIncludes) == 0 {
		isError = true
		sendError(w, "Includes is not present!")
	} else if p.From == "" {
		isError = true
		sendError(w, "Distributor is not mentioned!")
	} else if p.For == "" {
		isError = true
		sendError(w, "Sub-Distributor is not mentioned!")
	} else if len(p.Excludes) == 0 {
		isError = true
		sendError(w, "Excludes is not present!")
	}
	return isError

}

func sendError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(struct {
		Error string
	}{Error: message})
}
