package database

import (
	"api/src/config"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB(*sql.DB, error) {
	db, err := sql.Open("postgres", config.STRING_CONNECTION_DATABASE)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
	}

	return
}
