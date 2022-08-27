package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate routers
func Generate() *mux.Router {
	router := mux.NewRouter()
	return routes.RoutesConfiguration(router)
}
