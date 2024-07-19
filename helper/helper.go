package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/emray27/todo-golang/types"
)

func ReadTodos() ([]types.TodoStructure, error) {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error getting working directory: %w", err)
	}

	// Construct the path to the todos.json file
	absPath := filepath.Join(currentDir, "helper", "todos.json")

	// Open the file
	jsonFile, err := os.Open(absPath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer jsonFile.Close()

	// Read the file
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// Unmarshal the JSON
	var todos []types.TodoStructure
	err = json.Unmarshal(byteValue, &todos)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %w", err)
	}

	return todos, nil
}

func WriteTodos(todos []types.TodoStructure) error {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting working directory: %w", err)
	}

	// Construct the path to the todos.json file
	absPath := filepath.Join(currentDir, "helper", "todos.json")

	// Create or open the file
	file, err := os.Create(absPath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	// Marshal the todos to JSON
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling todos to json: %w", err)
	}

	// Write JSON data to the file
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error writing data to file: %w", err)
	}

	return nil
}
