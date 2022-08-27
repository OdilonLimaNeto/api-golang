package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// represent the users repository
type Users struct {
	db *sql.DB
}

// create a new users repository
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

// insert a new user into the database
func (repository Users) Create(user models.User) (uint64, error) {
	stmt, err := repository.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES ($1, $2, $3, $4) RETURNING id")
	fmt.Println("ERRO AQUI --> ", err)
	if err != nil {
		fmt.Println("ERRO AQUI --> ", err)
		return 0, err
	}
	defer stmt.Close()

	var id uint64
	err = stmt.QueryRow(user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
