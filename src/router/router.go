package router

import (
	"devbookapi/src/router/routes"

	"github.com/gorilla/mux"
)

func GenerateRouter() *mux.Router {
	router := mux.NewRouter()

	return routes.Configure(router)
}
