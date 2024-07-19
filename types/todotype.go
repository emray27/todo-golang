package types

import "time"

type TodoStructure struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	IsCompleted bool       `json:"isCompleted"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"` // *time.Time This allows the field to be nil
}

type TodoBodyToUpdate struct {
	Name        *string `json:"name"`
	IsCompleted *bool   `json:"isCompleted"`
}

func ValidateTodo(todo TodoStructure) bool {
	return len(todo.Name) >= 2

}
