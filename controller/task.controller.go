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
	var tasks []models.Task = services.GetTasks(10)
	jsonTask, err := json.Marshal(tasks)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonTask)
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
	// body, err := iotil.ReadAll(r.Body)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// var tasks []models.Task
	// var addedTasks []models.Task
	// json.Unmarshal(body, &tasks)
	// addedTasks = services.CreateTask(tasks)

	// jsonTask, err := json.Marshal(addedTasks)

	// fmt.Println(jsonTask)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.Write(jsonTask)
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
