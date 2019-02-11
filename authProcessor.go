package main

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type authProcessor int

func (a *authProcessor) getAuthenticationToken(creds credential) (string, error) {

	err := bcrypt.CompareHashAndPassword(creds.EncryptedPassword, []byte(creds.Password))
	if err != nil {
		return "", errors.New(invalidCredentials)
	}
	authToken, err := sessionObject.putIntoSessionMap(creds)
	if err != nil {
		return "", err
	}
	return authToken, nil

}

func (a *authProcessor) removeAuthenticationToken(authToken string) {

	sessionObject.deleteFromSessionMap(authToken)

}
