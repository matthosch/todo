package storage

import (
	"database/sql"

	"github.com/matthosch/todo/internal/todo"
	_ "github.com/mattn/go-sqlite3"
)

type sqliteStore struct {
	db *sql.DB
}

func NewSQLiteStore(dbPath string) *sqliteStore {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	createTableSQL := `CREATE TABLE IF NOT EXISTS todos (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"completed" BOOLEAN,
		"created_at" DATETIME
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
	return &sqliteStore{
		db: db,
	}
}

func (ds *sqliteStore) Add(todo *todo.Todo) error {
	// Implementation for adding a todo to SQLite database
	return nil
}

func (ds *sqliteStore) Load() ([]todo.Todo, error) {
	// Implementation for loading todos from SQLite database
	return []todo.Todo{}, nil
}

func (ds *sqliteStore) Save(todos []todo.Todo) error {
	// Implementation for saving todos to SQLite database
	return nil
}

// Close closes the database connection.
func (ds *sqliteStore) Close() error {
	return ds.db.Close()
}
