package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	API_PORT          = 0
	HOST              = ""
	DATABASE_PORT     = 0
	DATABASE_USER     = ""
	DATABASE_PASSWORD = ""
	DATABASE_NAME     = ""
	DATABASE_HOST     = ""

	STRING_CONNECTION_DATABASE = ""
)

func Init() {
	// Load the configuration environment variables
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// Set the configuration values
	API_PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		API_PORT = 3001
	}

	// Set the configuration values
	STRING_CONNECTION_DATABASE = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_NAME"))

}
