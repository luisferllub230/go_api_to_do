package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"lfernandez.com/todo/controller"
	"lfernandez.com/todo/db"
)

func main() {
	// Initialize the database
	// TODO: Move this to a config file
	db.InitConnection(db.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "lfernandez",
		Password: "strong.password",
		Name:     "todotasksgo",
		SSLMode:  "disable",
		DBType:   "postgres",
	})

	router := mux.NewRouter()

	// tasks routers
	router.HandleFunc("/get/tasks", controller.Get).Methods("GET")
	router.HandleFunc("/get/tasks/{id}", controller.GetById).Methods("GET")
	router.HandleFunc("/create/tasks", controller.Post).Methods("POST")
	router.HandleFunc("/update/tasks", controller.Put).Methods("PUT")
	router.HandleFunc("/delete/tasks", controller.Delete).Methods("DELETE")

	// options
	router.HandleFunc("/options", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Allow", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusNoContent)
	}).Methods("OPTIONS")

	// Start the server
	http.ListenAndServe(":3000", router)
}
