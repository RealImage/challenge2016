package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
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

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	credentialsObject.credentialMap = make(map[string]credential, 0)
	sessionObject.sessionMap = make(map[string]credential, 0)

	err = initAdminCredsAndHashes()
	if err != nil {
		log.Fatal(err)
	}

}

func panciHandle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
				respondError(w, http.StatusInternalServerError, internalServerError)
			}

		}()
		h.ServeHTTP(w, req)
	})
}

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/v1/authenticate-token", getAuthenticationToken)
	router.HandleFunc("/v1/users", userController)
	router.HandleFunc("/v1/distribution", validateDistribution)
	router.HandleFunc("/v1/invalidate-token", removeAuthenticationToken)

	handler := panciHandle(router)

	osSignalChannel := make(chan os.Signal, 1)
	signal.Notify(osSignalChannel, os.Interrupt, os.Kill)
	fmt.Println("Starting local server at 8080")

	go func() {
		<-osSignalChannel
		log.Println("Shutting down local server at 8080")
		os.Exit(0)
	}()

	log.Fatal(http.ListenAndServe(":8080", handler))
}
