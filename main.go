package main

import (
	"log"
	"net/http"

	"github.com/emray27/todo-golang/app"
)

func main() {
	// Setup the router with routes
	r := app.SetupRouter()

	// Define the server address
	address := ":8000"
	log.Printf("Starting server on %s", address)

	// Start the server and log any errors if the server fails to start
	if err := http.ListenAndServe(address, r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
