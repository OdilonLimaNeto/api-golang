package repositories

import (
	"api/src/database"
	"api/src/models"
	"database/sql"
	"log"
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
func Create(user models.User) (id uint64, err error) {
	var (
		name     = user.Name
		nick     = user.Nick
		email    = user.Email
		password = user.Password
	)
	connection, err := database.OpenConnectionDATABASE()
	if err != nil {
		return 0, err
	}
	defer connection.Close()

	SQL := `
	INSERT INTO users (name, nick, email, password) 
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	err = connection.QueryRow(SQL, name, nick, email, password).Scan(&id)

	return id, err
}

func Get(id int64) (user models.User, err error) {
	var (
		name      = user.Name
		nick      = user.Nick
		email     = user.Email
		password  = user.Password
		createdAt = user.CreatedAt
		updatedAt = user.UpdatedAt
		deletedAt = user.DeletedAt
	)
	connection, err := database.OpenConnectionDATABASE()
	if err != nil {
		return user, err
	}
	defer connection.Close()

	SQL := `
	SELECT 
		id, 
		name, 
		nick, 
		email, 
		password, 
		created_at, 
		updated_at, 
		deleted_at 
	FROM 
		users 
	WHERE 
		id = $1
	`
	err = connection.QueryRow(SQL, id).Scan(&user.ID, &name, &nick, &email, &password, &createdAt, &updatedAt, &deletedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetAll() (users []models.User, err error) {
	connection, err := database.OpenConnectionDATABASE()
	if err != nil {
		log.Printf("Error opening connection: %v", err)
		return nil, err
	}
	defer connection.Close()

	SQL := `
	SELECT 
		id, 
		name, 
		nick, 
		email, 
		password, 
		created_at, 
		updated_at, 
		deleted_at 
	FROM 
		users
	`
	rows, err := connection.Query(SQL)
	if err != nil {
		log.Printf("Error on get all users: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			log.Printf("Error to get rows: %v", err)
			continue
		}
		users = append(users, user)
	}
	return users, nil
}

func Update(id int64, user models.User) (int64, error) {
	var (
		name      = user.Name
		nick      = user.Nick
		email     = user.Email
		password  = user.Password
		createdAt = user.CreatedAt
		updatedAt = user.UpdatedAt
		deletedAt = user.DeletedAt
	)
	connection, err := database.OpenConnectionDATABASE()
	if err != nil {
		return 0, err
	}
	defer connection.Close()

	SQL := `
	UPDATE 
		users 
	SET 
		name = $1, 
		nick = $2, 
		email = $3, 
		password = $4, 
		created_at = $5,
		updated_at = $6,
		deleted_at = $7 
	WHERE 
		id = $8
	`
	response, err := connection.Exec(SQL, name, nick, email, password, createdAt, updatedAt, deletedAt, id)
	if err != nil {
		return 0, err
	}
	return response.RowsAffected()
}

func Delete(id int64) (int64, error) {
	connection, err := database.OpenConnectionDATABASE()
	if err != nil {
		log.Printf("Error opening connection: %v", err)
		return 0, err
	}
	defer connection.Close()

	SQL := `
	DELETE FROM users 
	WHERE 
		id = $1
	`
	response, err := connection.Exec(SQL, id)
	if err != nil {
		log.Printf("Error on delete user: %v", err)
		return 0, err
	}
	return response.RowsAffected()
}
