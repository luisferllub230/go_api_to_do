package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"lfernandez.com/todo/models"
	"lfernandez.com/todo/services"
)

func TaskGet(w http.ResponseWriter, r *http.Request) {
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

func TaskGetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	idInt, errAtoi := strconv.Atoi(id)

	if errAtoi != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errAtoi.Error()))
	}

	var task = services.GetTaskById(idInt)

	if task.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errAtoi.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func TaskPost(w http.ResponseWriter, r *http.Request) {
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

func TaskPut(w http.ResponseWriter, r *http.Request) {

	var tasks []models.Task
	json.NewDecoder(r.Body).Decode(&tasks)
	updateTasks, err := services.UpdateTask(tasks)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateTasks)
}

func TaskDelete(w http.ResponseWriter, r *http.Request) {

	var tasks []models.Task
	json.NewDecoder(r.Body).Decode(&tasks)
	deleteTasks, err := services.DeleteTask(tasks)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deleteTasks)
}
