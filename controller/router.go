package controller

import (
	"context"
	"github.com/gorilla/mux"
)

func NewRouter(ctx context.Context) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		handler := route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
