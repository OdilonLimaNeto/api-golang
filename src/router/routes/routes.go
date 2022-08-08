package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Struct representing all routes on API
type Routes struct {
	URI                    string
	Method                 string
	HandlerFunc            func(w http.ResponseWriter, r *http.Request)
	requiredAuthentication bool
}

func RoutesConfiguration(r *mux.Router) *mux.Router {
	routes := usersRoutes
	for _, route := range routes {
		r.HandleFunc(route.URI, route.HandlerFunc).Methods(route.Method)
	}
	return r
}
