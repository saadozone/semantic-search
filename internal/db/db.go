package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	// Create table if not exists
	createTable := `
	CREATE TABLE IF NOT EXISTS documents (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		content TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

