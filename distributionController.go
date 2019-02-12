package main

import (
	"encoding/json"
	"net/http"
)

var d distributionProcessor

func validateDistribution(w http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		respondError(w, http.StatusMethodNotAllowed, methodNotAllowed)
		return
	}

	ok, creds, err := isAlreadyLoggedIn(req)
	if err != nil || !ok {
		respondError(w, http.StatusBadRequest, loginFirst)
		return
	}

	newLocation := location{}

	err = json.NewDecoder(req.Body).Decode(&newLocation)
	defer req.Body.Close()
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = validateLocation(newLocation)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = d.validateDistribution(&newLocation, creds.Username)
	if err != nil {
		respondError(w, http.StatusForbidden, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, message{Message: distributionPermitted})
}
