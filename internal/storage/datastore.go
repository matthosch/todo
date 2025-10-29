package storage

import "github.com/matthosch/todo/internal/todo"

type DataStore interface {
	Add(todo *todo.Todo) error
	Load() ([]todo.Todo, error)
	Save(todos []todo.Todo) error
	Complete(id int) error
}
