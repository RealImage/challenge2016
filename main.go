package main

import (
	"fmt"
	"log"
	"net/http"
)

var defaultAdminHash, defaultDistributorHash []byte
var sessionObject session
var credentialsObject credentials

var users = make([]*user, 0)
var countries = make([]*country, 0)

func init() {
	err := prepareAllLocations()
	if err != nil {
		log.Fatal(err)
	}

	credentialsObject.credentialMap = make(map[string]credential, 0)
	sessionObject.sessionMap = make(map[string]credential, 0)

	err = initAdminCredsAndHashes()
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/v1/authenticate-token", getAuthenticationToken)
	router.HandleFunc("/v1/users", userController)
	router.HandleFunc("/v1/invalidate-token", removeAuthenticationToken)
	fmt.Println("Starting local server at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

	//TODO: Gracefully shutdown server
}
