package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"lfernandez.com/todo/models"
	"lfernandez.com/todo/services"
)

func Get(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	var err error

	tasks, err = services.GetTasks(10)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var task models.Task = services.GetTaskById(idInt)
	jsonTask, err := json.Marshal(task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonTask)
}

func Post(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	json.NewDecoder(r.Body).Decode(&tasks)
	createTasks, err := services.CreateTask(tasks)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createTasks)
}

func Put(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func Patch(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func Options(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
