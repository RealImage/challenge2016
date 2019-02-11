package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

var u userProcessor

func userController(w http.ResponseWriter, req *http.Request) {
	ok, creds, err := isAlreadyLoggedIn(req)
	if err != nil || !ok {
		respondError(w, http.StatusBadRequest, loginFirst)
		return
	}

	if req.Method == "GET" {
		getUsers(w, req, &creds)
		return
	}
	if req.Method == "POST" {
		createUser(w, req, &creds)
		return
	}

	respondError(w, http.StatusMethodNotAllowed, methodNotAllowed)

}

func getUsers(w http.ResponseWriter, req *http.Request, creds *credential) {

	outUser := u.getUsers(creds)
	respondJSON(w, http.StatusOK, outUser)

}

func createUser(w http.ResponseWriter, req *http.Request, creds *credential) {
	newUser := user{}

	err := json.NewDecoder(req.Body).Decode(&newUser)
	defer req.Body.Close()
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = validateUserRequestBody(&newUser)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = u.createUser(creds.Username, &newUser)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, message{Message: userCreated})
}

func validateUserRequestBody(newUser *user) error {

	if newUser.Name == "" {
		return errors.New(nameCannotBeEmpty)
	}
	if newUser.Role == "" {
		return errors.New(roleCannotBeEmpty)
	}
	return nil

}
