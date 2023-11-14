package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	MsgErr = "ERROR_MESSAGE"
	Msg    = "MESSAGE"
)

type RequestData struct {
	w http.ResponseWriter
	r *http.Request
}

// GetRouter creates a router and registers all the routes for the
// service and returns it.
func GetRouter() http.Handler {
	router := httprouter.New()
	setAuthorisationRoutes(router)

	router.Handler("GET", "/swagger/*path", httpSwagger.WrapHandler)
	return router
}
