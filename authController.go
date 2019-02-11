package main

import (
	"encoding/json"
	"net/http"
)

func getAuthenticationToken(w http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		respondError(w, http.StatusMethodNotAllowed, methodNotAllowed)
		return
	}

	creds := credential{}

	err := json.NewDecoder(req.Body).Decode(&creds)
	defer req.Body.Close()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ok := isValidCreds(&creds)
	if !ok {
		respondError(w, http.StatusUnauthorized, err.Error())
		return
	}

	ok, _, err = isAlreadyLoggedIn(req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if ok {
		respondError(w, http.StatusBadRequest, alreadyLogin)
		return
	}
	var a authProcessor

	authToken, err := a.getAuthenticationToken(creds)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == invalidPassword {
			status = http.StatusUnauthorized
		}
		respondError(w, status, err.Error())
		return
	}

	authCookie := &http.Cookie{
		Name:     authorizationToken,
		Value:    authToken,
		Path:     "/",
		Domain:   domain,
		HttpOnly: true,
	}
	authCookie.MaxAge = cookieMaxAge
	http.SetCookie(w, authCookie)

	respondJSON(w, http.StatusOK, message{Message: successfulLogin})

}
