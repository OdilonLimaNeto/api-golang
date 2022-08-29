package repositories

import (
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
func (repository Users) Create(user models.User) (id uint64, err error) {
	var (
		name     = user.Name
		nick     = user.Nick
		email    = user.Email
		password = user.Password
	)
	SQL := `
	INSERT INTO public.user (name, nick, email, password) 
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	err = repository.db.QueryRow(SQL, name, nick, email, password).Scan(&id)

	return id, err
}

func (repository Users) Get(id int64) (user models.User, err error) {
	SQL := `
	SELECT 
		id, 
		name, 
		nick, 
		email, 
		password, 
		created_at, 
		updated_at
	FROM 
		public.user
	WHERE 
		id = $1 
	`
	err = repository.db.QueryRow(SQL, id).Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repository Users) GetAll() (users []models.User, err error) {
	SQL := `
	SELECT id,
		name,
		nick,
		email,
		password,
		created_at,
		updated_at
	FROM 
		public.user
	`
	rows, err := repository.db.Query(SQL)
	if err != nil {
		log.Printf("Error on get all users: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Printf("Error to get rows: %v", err)
			continue
		}
		users = append(users, user)
	}
	return users, nil
}

func (repository Users) Update(id int64, user models.User) (int64, error) {
	SQL := `
	UPDATE 
		public.user
	SET 
		name = $1, 
		nick = $2, 
		email = $3, 
		password = $4,
		created_at = $5,
		updated_at = $6
	WHERE 
		id = $7
	`
	response, err := repository.db.Exec(SQL, user.Name, user.Nick, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, id)
	if err != nil {
		return 0, err
	}
	return response.RowsAffected()
}

func (repository Users) Delete(id int64) (int64, error) {
	SQL := `DELETE FROM public.user WHERE id = $1`
	response, err := repository.db.Exec(SQL, id)
	if err != nil {
		log.Printf("Error on delete user: %v", err)
		return 0, err
	}
	return response.RowsAffected()
}
