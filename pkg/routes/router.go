package routes

import (
	"net/http"

	"chng2016/pkg/handlers"

	"github.com/gin-gonic/gin"
)

// Route ...
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

func NewRoutes(handler handlers.Handler) Routes {
	routes := Routes{
		{
			Name:        "heath check point",
			Method:      http.MethodGet,
			Pattern:     "/health",
			HandlerFunc: handler.HealthHandler,
		},
		{
			Name:        "add distributor details",
			Method:      http.MethodPost,
			Pattern:     "/add-distributor",
			HandlerFunc: handler.AddDistributor,
		},
		{
			Name:        "check distributor permission",
			Method:      http.MethodGet,
			Pattern:     "/:distributorID/checkPermission",
			HandlerFunc: handler.CheckDistributorPermission,
		},
		{
			Name:        "add sub distributor",
			Method:      http.MethodPost,
			Pattern:     "/:distributorID/add-distributor",
			HandlerFunc: handler.AddSubDistributor,
		},
	}
	return routes
}

func AttachRoutes(server *gin.Engine, routes Routes) {
	for _, route := range routes {
		server.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}
}
