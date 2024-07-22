package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emray27/todo-golang/app"
	"github.com/joho/godotenv"
)

func main() {
	// Setup the router with routes
	r := app.SetupRouter()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not able to load godotenv")
	}

	_, mongoError := app.GetMongoClient()

	if mongoError != nil {
		fmt.Println(mongoError)
		return
	}

	// Define the server address
	address := ":8000"
	log.Printf("Starting server on %s", address)

	// Start the server and log any errors if the server fails to start
	if err := http.ListenAndServe(address, r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
