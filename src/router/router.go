package router

import "github.com/gorilla/mux"

func ToGenerate() *mux.Router {
	return mux.NewRouter()
}
