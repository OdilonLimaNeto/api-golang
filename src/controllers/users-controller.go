package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/lib/pq"
)

// Create is the handler for POST /users
func Create(w http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	var user models.User
	if err = json.Unmarshal(request, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	db, err := database.ConnectDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	repository := repositories.NewUsersRepository(db)

	id, err := repository.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Println("id --> ", id)

	w.Write([]byte(fmt.Sprintf("Inserted ID %d", id)))
}

// List is the handler for /users
func List(w http.ResponseWriter, r *http.Request) {
}

// Get is the handler for /users/{id}
func Get(w http.ResponseWriter, r *http.Request) {
}

// Update is the handler for PUT /users/{id}
func Update(w http.ResponseWriter, r *http.Request) {
}

// Delete is the handler for DELETE /users/{id}
func Delete(w http.ResponseWriter, r *http.Request) {
}
