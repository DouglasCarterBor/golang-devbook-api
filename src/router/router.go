package router

import "github.com/gorilla/mux"

// ToGenerate will return a router with the configured routes
func ToGenerate() *mux.Router {
	return mux.NewRouter()
}
