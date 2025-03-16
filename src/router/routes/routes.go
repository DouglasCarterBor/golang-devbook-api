package routes

import "net/http"

type Route struct {
	Uri string
	Method string
	Func func(http.ResponseWriter, *http.Request)
	RequiresAutentication bool
}