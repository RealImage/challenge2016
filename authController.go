package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

var a authProcessor

func getAuthenticationToken(w http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		respondError(w, http.StatusMethodNotAllowed, methodNotAllowed)
		return
	}

	creds := credential{}

	err := json.NewDecoder(req.Body).Decode(&creds)
	defer req.Body.Close()
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = validateCredsRequestBody(&creds)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	ok := isValidCreds(&creds)
	if !ok {
		respondError(w, http.StatusUnauthorized, invalidCredentials)
		return
	}

	ok, _, err = isAlreadyLoggedIn(req)
	if err != nil {
		respondError(w, http.StatusBadRequest, loginFirst)
		return
	}

	if ok {
		respondError(w, http.StatusBadRequest, alreadyLogin)
		return
	}

	authToken, err := a.getAuthenticationToken(creds)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == invalidCredentials {
			status = http.StatusUnauthorized
		}
		respondError(w, status, err.Error())
		return
	}

	authCookie := &http.Cookie{
		Name:     authenticationToken,
		Value:    authToken,
		Path:     "/",
		Domain:   domain,
		HttpOnly: true,
	}
	authCookie.MaxAge = cookieMaxAge
	http.SetCookie(w, authCookie)

	respondJSON(w, http.StatusOK, message{Message: successfulLogin})

}

func removeAuthenticationToken(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		respondError(w, http.StatusMethodNotAllowed, methodNotAllowed)
		return
	}
	ok, _, err := isAlreadyLoggedIn(req)
	if err != nil || !ok {
		respondError(w, http.StatusBadRequest, loginFirst)
		return
	}

	authToken, _ := getCookieValue(req, authenticationToken)

	a.removeAuthenticationToken(authToken)

	authCookie := &http.Cookie{
		Name:     authenticationToken,
		Value:    "",
		Path:     "/",
		Domain:   domain,
		HttpOnly: true,
		MaxAge:   -1,
	}
	http.SetCookie(w, authCookie)

	respondJSON(w, http.StatusNoContent, nil)

}

func validateCredsRequestBody(creds *credential) error {

	if creds.Username == "" || creds.Password == "" {
		return errors.New(invalidCredentials)
	}
	return nil

}
