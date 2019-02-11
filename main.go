package main

import (
	"log"
	"net/http"
)

func init() {
	// err := prepareAllLocations()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	credentialsObject.credentialMap = make(map[string]credential, 1)
	sessionObject.sessionMap = make(map[string]credential, 1)

	err := initAdminCreds()
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/v1/authenticate-token", getAuthenticationToken)
	router.HandleFunc("/v1/users", userController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
