package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Name     string `json:"name"`
		Nick     string `json:"nick"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Parse the request body into the struct
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Validate the request body
	if request.Name == "" || request.Nick == "" || request.Email == "" || request.Password == "" {
		log.Printf("Error validating request body: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	database, err := database.OpenConnectionDATABASE()
	if err != nil {
		log.Printf("Error opening connection to database: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer database.Close()

	repository := repositories.NewUsersRepository(database)

	id, err := repository.Create(models.User{
		Name:     request.Name,
		Nick:     request.Nick,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		log.Printf("Error creating user: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("{\"New user ID\": %d}", id)))
}

func List(w http.ResponseWriter, r *http.Request) {

	database, err := database.OpenConnectionDATABASE()
	if err != nil {
		log.Printf("Error opening connection to database: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer database.Close()

	repository := repositories.NewUsersRepository(database)

	users, err := repository.GetAll()
	if err != nil {
		log.Printf("Error getting users: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func Get(w http.ResponseWriter, r *http.Request) {
	var err error
	id := mux.Vars(r)["id"]
	if id == "" {
		log.Printf("Error getting ID: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	convertedID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("Error converting ID: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	database, err := database.OpenConnectionDATABASE()
	if err != nil {
		log.Printf("Error opening connection to database: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer database.Close()

	repository := repositories.NewUsersRepository(database)

	user, err := repository.Get(convertedID)
	if err != nil {
		log.Printf("Error getting user: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Name     string `json:"name"`
		Nick     string `json:"nick"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	params := mux.Vars(r)
	convertedID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		log.Printf("Error converting id: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if request.Name == "" || request.Nick == "" || request.Email == "" || request.Password == "" {
		log.Printf("Error validating request body: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	database, err := database.OpenConnectionDATABASE()
	if err != nil {
		log.Printf("Error opening connection to database: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer database.Close()

	repository := repositories.NewUsersRepository(database)

	_, err = repository.Update(convertedID, models.User{
		Name:     request.Name,
		Nick:     request.Nick,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		log.Printf("Error updating user: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"Updated user ID\": %d}", convertedID)))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	convertedID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		log.Printf("Error converting id: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	database, err := database.OpenConnectionDATABASE()
	if err != nil {
		log.Printf("Error opening connection to database: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer database.Close()

	repository := repositories.NewUsersRepository(database)

	_, err = repository.Delete(convertedID)
	if err != nil {
		log.Printf("Error deleting user: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"Deleted user ID\": %d}", convertedID)))
}
