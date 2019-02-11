package main

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type authProcessor int

func (a *authProcessor) getAuthenticationToken(creds credential) (string, error) {

	err := bcrypt.CompareHashAndPassword(creds.EncryptedPassword, []byte(creds.Password))
	if err != nil {
		return "", errors.New(invalidPassword)
	}
	authToken, err := sessionObject.putIntoSessionMap(creds)
	if err != nil {
		return "", err
	}
	return authToken, nil

}
