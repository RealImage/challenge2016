package main

import (
	"golang/app"
	"golang/constant"
	Service "golang/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// In this function the connection to handler layer is established and the endpoint for the API is defined.
func main() {

	//This router will be used to define the routes and handle the incoming HTTP requests for the API.
	router := mux.NewRouter()

	h := app.Handlers{Service: &Service.DefaultService{}}

	// The API endpoint & method is defined and connected to handler layer.
	router.HandleFunc("/distributors/permissions", h.DistributorPermissions).Methods(http.MethodPost)

	// Since port number is a sensitive information, it is stored in the constant file, imported & used here.
	listenAddr := constant.PORT
	log.Printf("About to listen on port %s.", listenAddr)
	log.Fatal(http.ListenAndServe(":"+listenAddr, router))
}
