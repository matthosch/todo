package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/matthosch/todo/internal/todo"
)

type jsonStore struct {
	filePath string
}

func NewJSONStore(filePath string) *jsonStore {
	return &jsonStore{
		filePath: filePath,
	}
}

func (ds *jsonStore) Add(todo *todo.Todo) error {
	todos, err := ds.Load()
	if err != nil {
		return fmt.Errorf("error loading todos: %v", err)
	}

	todos = append(todos, *todo)
	if err := ds.Save(todos); err != nil {
		return fmt.Errorf("error saving todos: %v", err)
	}
	return nil
}

func (ds *jsonStore) Load() ([]todo.Todo, error) {
	data, err := os.ReadFile(ds.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Data file does not exist, creating %s.\n", ds.filePath)
			err := os.WriteFile(ds.filePath, []byte("[]"), 0644)
			if err != nil {
				return nil, err
			}
			return []todo.Todo{}, nil
		}
	}

	var todos []todo.Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, fmt.Errorf("error unmarshaling todos: %w", err)
	}

	return todos, nil
}

func (ds *jsonStore) Save(todos []todo.Todo) error {
	data, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return fmt.Errorf("error marshaling todos: %w", err)
	}

	err = os.WriteFile(ds.filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing todos to file: %w", err)
	}
	return nil
}
