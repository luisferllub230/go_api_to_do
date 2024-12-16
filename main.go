package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"lfernandez.com/todo/db"
	"lfernandez.com/todo/routes"
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

	// Create a new router
	router := mux.NewRouter()
	router.HandleFunc("/", routes.Get).Methods("GET")

	// Start the server
	http.ListenAndServe(":3000", router)
}
