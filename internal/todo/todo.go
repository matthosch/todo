package todo

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

func NewTodo(description string) *Todo {
	return &Todo{
		ID:        uuid.New().String(),
		Title:     description,
		Completed: false,
		CreatedAt: time.Now(),
	}
}
