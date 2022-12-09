package db

import (
	"database/sql"
)

func setTables(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id		    SERIAL PRIMARY KEY ,
			username		TEXT NOT NULL UNIQUE,
			password	TEXT
		);
	`)

	if err != nil {
		panic(err)
	}
}
