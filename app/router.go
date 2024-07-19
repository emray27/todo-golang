package app

import (
	"github.com/gorilla/mux"
)

// SetupRouter initializes the router and sets up the API routes
func SetupRouter() *mux.Router {
	// Create a new router
	r := mux.NewRouter()

	// Define route for getting all todos
	r.HandleFunc("/", GetAllTodos).Methods("GET")

	// Define route for creating a new todo
	r.HandleFunc("/createTodo", CreateTodo).Methods("POST")

	// Define route for updating an existing todo
	r.HandleFunc("/updateTodo/{id}", UpdateTodo).Methods("PUT")

	// Define route for deleting an existing todo
	r.HandleFunc("/deleteTodo/{id}", DeleteTodo).Methods("DELETE")

	// Return the configured router
	return r
}
