package routes

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Uri string
	Method string
	Func func(http.ResponseWriter, *http.Request)
	RequiresAutentication bool
}

func Config(router *mux.Router) *mux.Router {
	routes := routesUsers
	
	for _, route := range routes {
		router.HandleFunc(route.Uri, route.Func).Methods(route.Method)
	}

	return router
}
