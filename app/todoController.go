package app

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"net/http"

	"github.com/emray27/todo-golang/helper"
	"github.com/emray27/todo-golang/types"
	"github.com/gorilla/mux"
)

// GetAllTodos handles the GET request to retrieve all todos
func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := helper.ReadTodos()
	if err != nil {
		http.Error(w, "Failed to read todos", http.StatusInternalServerError)
		fmt.Println("Error reading todos:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// CreateTodo handles the POST request to create a new todo
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	todos, err := helper.ReadTodos()
	if err != nil {
		http.Error(w, "Error while reading todos", http.StatusInternalServerError)
		fmt.Println("Error while reading todos:", err)
		return
	}

	var newTodo types.TodoStructure
	decoder := json.NewDecoder(r.Body)
	errInDecoding := decoder.Decode(&newTodo)
	if errInDecoding != nil {
		http.Error(w, "Error while decoding newTodo", http.StatusBadRequest)
		fmt.Println("Error while decoding newTodo:", errInDecoding)
		return
	}

	if !types.ValidateTodo(newTodo) {
		http.Error(w, "Validation error: The length of your todo name must be > 2", http.StatusBadRequest)
		fmt.Println("Validation error: The length of your todo name must be > 2")
		return
	}

	newTodo.Id = len(todos) + 1
	newTodo.CreatedAt = time.Now()

	todos = append(todos, newTodo)

	errInWriteTodos := helper.WriteTodos(todos)
	if errInWriteTodos != nil {
		http.Error(w, "Error while writing the todos in the DB", http.StatusInternalServerError)
		fmt.Println("Error while writing the todos in the DB:", errInWriteTodos)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTodo)
}

// UpdateTodo handles the PUT request to update an existing todo
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todos, errInReadingTodos := helper.ReadTodos()
	if errInReadingTodos != nil {
		http.Error(w, "Error in reading todos", http.StatusInternalServerError)
		fmt.Println("Error in reading todos:", errInReadingTodos)
		return
	}

	params := mux.Vars(r)
	idToUpdate := params["id"]
	intIdToUpdate, errInConvertingStringToInteger := strconv.Atoi(idToUpdate)
	if errInConvertingStringToInteger != nil {
		http.Error(w, "Error in converting id", http.StatusBadRequest)
		fmt.Println("Error in converting id:", errInConvertingStringToInteger)
		return
	}

	var todoBodyToUpdate types.TodoBodyToUpdate
	decoder := json.NewDecoder(r.Body)
	errInDecoding := decoder.Decode(&todoBodyToUpdate)
	if errInDecoding != nil {
		http.Error(w, "Error while decoding the body", http.StatusBadRequest)
		fmt.Println("Error while decoding the body:", errInDecoding)
		return
	}

	now := time.Now()
	for index, todo := range todos {
		if intIdToUpdate == todo.Id {
			if todoBodyToUpdate.Name != nil {
				todos[index].Name = *todoBodyToUpdate.Name
			}
			if todoBodyToUpdate.IsCompleted != nil {
				todos[index].IsCompleted = *todoBodyToUpdate.IsCompleted
			}
			todos[index].UpdatedAt = &now
			break
		}
	}

	errInWriteTodos := helper.WriteTodos(todos)
	if errInWriteTodos != nil {
		http.Error(w, "Error while writing the updated todos", http.StatusInternalServerError)
		fmt.Println("Error while writing the updated todos:", errInWriteTodos)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Updated!")
}

// DeleteTodo handles the DELETE request to delete an existing todo
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todos, err := helper.ReadTodos()
	if err != nil {
		http.Error(w, "Error while reading todos", http.StatusInternalServerError)
		fmt.Println("Error while reading todos:", err)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	idInInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error in converting id", http.StatusBadRequest)
		fmt.Println("Error in converting id:", err)
		return
	}

	for index, todo := range todos {
		if idInInt == todo.Id {
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}

	errInWriteTodos := helper.WriteTodos(todos)
	if errInWriteTodos != nil {
		http.Error(w, "Error while writing the updated todos", http.StatusInternalServerError)
		fmt.Println("Error while writing the updated todos:", errInWriteTodos)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Deleted!")
}
