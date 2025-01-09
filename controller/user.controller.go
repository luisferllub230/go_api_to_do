package controller

import (
	"encoding/json"
	"net/http"

	"lfernandez.com/todo/models"
	"lfernandez.com/todo/services"
)

func UserGet(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	var err error

	users, err = services.GetUsers(10)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func UserPost(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	json.NewDecoder(r.Body).Decode(&users)
	createUsers, err := services.CreateUsers(users)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createUsers)
}

func UserPut(w http.ResponseWriter, r *http.Request) {

	var users []models.User
	json.NewDecoder(r.Body).Decode(&users)
	UpdateUsers, err := services.UpdateUsers(users)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UpdateUsers)
}

func UserDelete(w http.ResponseWriter, r *http.Request) {

	var users []models.User
	json.NewDecoder(r.Body).Decode(&users)
	deletedUsers, err := services.DeleteUsers(users)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deletedUsers)
}
