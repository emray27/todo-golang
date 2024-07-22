package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoStructure struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	IsCompleted bool       `json:"isCompleted"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"` // *time.Time This allows the field to be nil
}

type TodoMongoStructure struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	IsCompleted bool               `json:"isCompleted" bson:"isCompleted"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   *time.Time         `json:"updatedAt" bson:"updatedAt,omitempty"`
}

type TodoBodyToUpdate struct {
	Name        *string `json:"name"`
	IsCompleted *bool   `json:"isCompleted"`
}

func ValidateTodo(todo TodoStructure) bool {
	return len(todo.Name) >= 2
}
