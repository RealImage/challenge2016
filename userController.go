package main

import "net/http"

func userController(w http.ResponseWriter, req *http.Request) {
	ok, creds, err := isAlreadyLoggedIn(req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !ok {
		respondError(w, http.StatusUnauthorized, loginFirst)
		return
	}

	if req.Method == "GET" {
		getUsers(w, req, creds)
		return
	}
	if req.Method == "POST" {
		createUser(w, req, creds)
		return
	}

	respondError(w, http.StatusMethodNotAllowed, methodNotAllowed)

}

func getUsers(w http.ResponseWriter, req *http.Request, creds credential) {
	var u userProcessor
	outUser := u.getUsers(creds)
	respondJSON(w, http.StatusOK, outUser)

}

func createUser(w http.ResponseWriter, req *http.Request, creds credential) {
}
