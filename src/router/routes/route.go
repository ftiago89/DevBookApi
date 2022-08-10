package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI        string
	Method     string
	Function   func(http.ResponseWriter, *http.Request)
	ShouldAuth bool
}

func Configure(r *mux.Router) *mux.Router {
	for _, route := range userRoutes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
