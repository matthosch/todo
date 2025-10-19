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

func (ds *DataStore) Add(todo todo) error {
	todos, err := ds.Load()
	if err != nil {
		return fmt.Errorf("error loading todos: %v", err)
	}

	todos = append(todos, todo)
	if err := ds.Save(todos); err != nil {
		return fmt.Errorf("error saving todos: %v", err)
	}
	return nil
}

func (ds *DataStore) Load() ([]todo, error) {
	data, err := os.ReadFile(ds.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			os.Create(ds.FilePath)
			return []todo{}, nil
		}
		return nil, err
	}
	var todos []todo
	json.Unmarshal(data, &todos)

	return todos, nil
}

func (ds *DataStore) Save(todos []todo) error {
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
