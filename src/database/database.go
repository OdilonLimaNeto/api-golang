package database

import (
	"api/src/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func OpenConnectionDATABASE() (*sql.DB, error) {
	configurations := config.GetDB()

	var (
		HOST     = configurations.HOST
		PORT     = configurations.PORT
		USER     = configurations.USER
		PASSWORD = configurations.PASSWORD
		NAME     = configurations.NAME
	)

	STRING_CONNECTION_DATABASE := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, NAME)

	database, err := sql.Open("postgres", STRING_CONNECTION_DATABASE)
	if err != nil {
		log.Printf("Error open connection to database: %s", err)
		return nil, err
	}

	err = database.Ping()
	if err != nil {
		log.Printf("Error ping to database: %s", err)
	}
	return database, nil
}
