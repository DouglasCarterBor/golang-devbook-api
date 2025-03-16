package router

import (
	"api/src/router/routes"
	"github.com/gorilla/mux"
)

func ToGenerate() *mux.Router {
	router := mux.NewRouter()
	return routes.Config(router)
}
