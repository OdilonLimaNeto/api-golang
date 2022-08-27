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
	// Load configurations
	config.Load()

	// Create a new router
	routes := router.Generate()

	fmt.Printf("Started server on PORT %s", config.GetServerPORT())

	// Start the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.GetServerPORT()), routes))
}
