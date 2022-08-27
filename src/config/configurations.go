package config

import (
	"fmt"

	_ "github.com/lib/pq"
)

var (
	API_PORT                   = 5000
	STRING_CONNECTION_DATABASE = ""
	POSTGRES_HOST              = "localhost"
	POSTGRES_PORT              = 5432
	POSTGRES_USER              = "api"
	POSTGRES_PASSWORD          = "api"
	POSTGRES_DB                = "api"
)

func Init() {
	// Load the configuration environment variables
	var err error

	// Set the configuration values
	if err != nil {
		API_PORT = 5000
	}

	// Set the configuration values
	STRING_CONNECTION_DATABASE = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", POSTGRES_HOST, POSTGRES_PORT, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB)

	fmt.Println("Connect database status Ok:", STRING_CONNECTION_DATABASE)
}
