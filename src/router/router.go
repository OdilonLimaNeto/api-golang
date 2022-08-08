package router

import "github.com/gorilla/mux"

// Generate routers
func Generate() *mux.Router {
	return mux.NewRouter()
}
