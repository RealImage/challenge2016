package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func respondJSON(w http.ResponseWriter, status int, data interface{}) {

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error in Marshalling data: ", err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonBytes)
}

func respondError(w http.ResponseWriter, status int, err string) {
	respondJSON(w, status, errorMessage{err})
}

func getPasswordHash(password string) ([]byte, error) {

	if len(password) == 0 {
		return nil, errors.New("Empty Password")
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	return hashBytes, nil

}

func isAlreadyLoggedIn(req *http.Request) (bool, credential, error) {

	if len(req.Cookies()) == 0 {
		return false, credential{}, nil
	}
	_, authToken, err := getCookieAndValue(req, authenticationToken)
	if err != nil {
		return false, credential{}, err
	}

	creds, ok := sessionObject.getFromSessionMap(authToken)
	return ok, creds, nil
}

func getCookieAndValue(req *http.Request, cookieName string) (*http.Cookie, string, error) {

	authCookie, err := req.Cookie(cookieName)
	if err != nil {
		return nil, "", err
	}
	return authCookie, authCookie.Value, nil
}

func isValidCreds(inputCreds *credential) bool {
	if creds, ok := credentialsObject.getFromCredentialMap(inputCreds.Username); ok {
		inputCreds.EncryptedPassword = creds.EncryptedPassword
		return true
	}
	return false
}

func getUserFromUsers(username string) *user {
	for _, u := range users {
		if u.Name == username {
			return u
		}
	}

	for _, user := range users {
		if outUser := getUserFromUsersHelper(user, username); outUser != nil {
			return outUser
		}
	}

	return nil
}

func getUserFromUsersHelper(inUser *user, username string) *user {

	for _, child := range inUser.Children {
		if child.Name == username {
			return child
		}
		if len(child.Children) == 0 {
			return nil
		}
		if outUser := getUserFromUsersHelper(child, username); outUser != nil {
			return outUser
		}

	}

	return nil

}
