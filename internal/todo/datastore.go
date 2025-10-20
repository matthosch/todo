package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type DataStore struct {
	FilePath string
}

func NewDataStore(filePath string) *DataStore {
	return &DataStore{
		FilePath: filePath,
	}
}

func (ds *DataStore) Add(todo *Todo) error {
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

func (ds *DataStore) Load() ([]Todo, error) {
	data, err := os.ReadFile(ds.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Data file does not exist, creating %s.\n", ds.FilePath)
			err := os.WriteFile(ds.FilePath, []byte{}, 0644)
			if err != nil {
				return nil, err
			}
			return []Todo{}, nil
		}
	}

	var todos []Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, fmt.Errorf("error unmarshaling todos: %w", err)
	}

	return todos, nil
}

func (ds *DataStore) Save(todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return fmt.Errorf("error marshaling todos: %w", err)
	}

	err = os.WriteFile(ds.FilePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing todos to file: %w", err)
	}
	return nil
}
