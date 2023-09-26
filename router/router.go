package router

import (
	"devbook-api/router/routes"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.RoutesConfiguration(r)
}
