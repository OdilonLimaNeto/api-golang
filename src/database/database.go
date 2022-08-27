package database

import (
	"api/src/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.STRING_CONNECTION_DATABASE)
	if err != nil {
		fmt.Println("ERRO AQUI ", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("ERRO AQUI ", err)
		panic(err)
	}

	return db, nil
}
