package controller

import (
	"distributor/constants"
	"distributor/handler"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	// Create Distributor
	Route{
		"CreateDistributor",
		constants.POST,
		constants.CreateDistributorURI,
		handler.CreateDistributorHandler,
	},
	Route{
		"GetDistributorLocationDetails",
		constants.GET,
		constants.GetDistributorLocationDetailsURI,
		handler.GetDistributorLocationDetailsHandler,
	},
	Route{
		"GetDistributorDetails",
		constants.GET,
		constants.GetDistributorDetailsURI,
		handler.GetDistributorDetailsHandler,
	},
}
