package main

import (
	"encoding/json"
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

	ok := isValidCreds(&creds)
	if !ok {
		respondError(w, http.StatusUnauthorized, invalidCreds)
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
	ok, creds, err := isAlreadyLoggedIn(req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !ok {
		respondError(w, http.StatusBadRequest, loginFirst)
		return
	}

	_, authToken, err := getCookieAndValue(req, authenticationToken)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	a.removeAuthenticationToken(&creds, authToken)

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
