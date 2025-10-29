package todo

import (
	"time"
)

type Todo struct {
	Task      string    `json:"task"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

func NewTodo(task string) *Todo {
	return &Todo{
		Task:      task,
		Completed: false,
		CreatedAt: time.Now(),
	}
}
