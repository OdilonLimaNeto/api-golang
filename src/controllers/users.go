package controllers

import (
	"net/http"
)

// Create is the handler for POST /users
func Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user"))
}

// List is the handler for /users
func List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List users"))
}

// Get is the handler for /users/{id}
func Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user"))
}

// Update is the handler for PUT /users/{id}
func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

// Delete is the handler for DELETE /users/{id}
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
