package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	// Load the configuration environment.
	config.Init()

	// Create a new router
	routes := router.Generate()

	fmt.Println("Started server on port", config.API_PORT)

	// Start the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.API_PORT), routes))

}
